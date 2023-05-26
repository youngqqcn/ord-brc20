
# inscribe


铭刻命令

先分析inscribe命令的主流程：

```rust
pub(crate) fn run(self, options: Options) -> Result {

    // 读取文件获取获取铭刻内容
    let inscription = Inscription::from_file(options.chain(), &self.file)?;

    // 更新index索引
    let index = Index::open(&options)?;
    index.update()?;

    // 加载rpc和钱包
    let client = options.bitcoin_rpc_client_for_wallet_command(false)?;

    // 获取钱包utxos集合
    let mut utxos = index.get_unspent_outputs(Wallet::load(&options)?)?;

    // 获取已有的铭刻
    let inscriptions = index.get_inscriptions(None)?;

    // commit交易找零金额
    let commit_tx_change = [get_change_address(&client)?, get_change_address(&client)?];

    // 铭文接受者地址
    let reveal_tx_destination = self
      .destination
      .map(Ok)
      .unwrap_or_else(|| get_change_address(&client))?;

    // 构造
    //     未签名的commit_tx
    //     已签名的reveal_tx（taproot）交易
    //     已经恢复密钥对（因为commit_tx的taproot输出，
    //        是一个临时创建中间密钥对(地址)，因此，reveal_tx可以直接用这个“临时”密钥对的私钥进行签名，
    //        恢复密钥对用于对交易的恢复，不必细究
    let (unsigned_commit_tx, reveal_tx, recovery_key_pair) =
      Inscribe::create_inscription_transactions(
        self.satpoint,
        inscription,
        inscriptions,
        options.chain().network(),
        utxos.clone(),
        commit_tx_change,
        reveal_tx_destination,
        self.commit_fee_rate.unwrap_or(self.fee_rate),
        self.fee_rate,
        self.no_limit,
      )?;

    // 将 commit_tx的输出，亦即 reveal_tx的输入，插入index保存，
    utxos.insert(
      reveal_tx.input[0].previous_output,
      Amount::from_sat(
        unsigned_commit_tx.output[reveal_tx.input[0].previous_output.vout as usize].value,
      ),
    );

    // commit_tx 和 reveal_tx 总共的交易矿工费
    let fees = Self::calculate_fee(&unsigned_commit_tx, &utxos) + Self::calculate_fee(&reveal_tx, &utxos);

    if self.dry_run {
        // ======== 虚晃一枪， 不上链 ==============
        print_json(Output {
            commit: unsigned_commit_tx.txid(),
            reveal: reveal_tx.txid(),
            inscription: reveal_tx.txid().into(),
            fees,
        })?;
    } else {
        // ========== 动真格的 ， 上链 ============

        // 是否备份上面的“临时”密钥对的recovery_key
        if !self.no_backup {
            Inscribe::backup_recovery_key(&client, recovery_key_pair, options.chain().network())?;
        }

        // 对未签名的commit_tx进行签名
        let signed_raw_commit_tx = client
            .sign_raw_transaction_with_wallet(&unsigned_commit_tx, None, None)?
            .hex;

        // 广播已签名的commit_tx交易
        let commit = client
            .send_raw_transaction(&signed_raw_commit_tx)
            .context("Failed to send commit transaction")?;

        // 广播已签名的reveal_tx交易
        let reveal = client
            .send_raw_transaction(&reveal_tx)
            .context("Failed to send reveal transaction")?;

        // 打印结果
        print_json(Output {
            commit,
            reveal,
            inscription: reveal.into(),
            fees,
        })?;
    };

    Ok(())
  }
```



接着，我们来分析构造  `commit_tx` 以及 `reveal_tx` 的细节

[src/subcommand/wallet/inscribe.rs](https://github.com/youngqqcn/ord/blob/master/src/subcommand/wallet/inscribe.rs)


```rust
fn create_inscription_transactions(
    satpoint: Option<SatPoint>,                         // 可指定使用某个 UTXO来进行 inscribe
    inscription: Inscription,                           // 铭刻内容
    inscriptions: BTreeMap<SatPoint, InscriptionId>,    // 已铭刻的集合
    network: Network,                                   // 比特币网络类型
    utxos: BTreeMap<OutPoint, Amount>,                  // utxo集合
    change: [Address; 2],                               // commit_tx交易找零地址
    destination: Address,                               // 铭文接收地址
    commit_fee_rate: FeeRate,                           // commit_tx 费率
    reveal_fee_rate: FeeRate,                           // reveal_tx 费率
    no_limit: bool,                                     // 是否限制reveal交易weight权重
) -> Result<(Transaction, Transaction, TweakedKeyPair)> {

    // 1) 获取进行铭刻UTXO
    let satpoint = if let Some(satpoint) = satpoint {
        // 如果指定来UTXO, 则直接使用指定的utxo进行铭刻
        satpoint
    } else {
        // 如果没有指定utxo, 则在utxos集合中找一个没有铭刻过的utxo

        let inscribed_utxos = inscriptions
        .keys()
        .map(|satpoint| satpoint.outpoint)
        .collect::<BTreeSet<OutPoint>>();

        utxos
        .keys()
        .find(|outpoint| !inscribed_utxos.contains(outpoint))
        .map(|outpoint| SatPoint {
            outpoint: *outpoint,
            offset: 0,
        })
        .ok_or_else(|| anyhow!("wallet contains no cardinal utxos"))?
    };

    // 2) 判断上一步的UTXO是否没有被铭刻过
    for (inscribed_satpoint, inscription_id) in &inscriptions {
        if inscribed_satpoint == &satpoint {
        return Err(anyhow!("sat at {} already inscribed", satpoint));
        }

        if inscribed_satpoint.outpoint == satpoint.outpoint {
        return Err(anyhow!(
            "utxo {} already inscribed with inscription {inscription_id} on sat {inscribed_satpoint}",
            satpoint.outpoint,
        ));
        }
    }

    // 3） 搞一个“临时”密钥对，用来作为 commit_tx的接受者，并作为 reveal_tx的花费（揭示）者
    let secp256k1 = Secp256k1::new();
    let key_pair = UntweakedKeyPair::new(&secp256k1, &mut rand::thread_rng());
    let (public_key, _parity) = XOnlyPublicKey::from_keypair(&key_pair);


    // 4） 按照ordinals官方的脚本规范，创建reveal脚本， 将铭文内容附加到reveal脚本
    let reveal_script = inscription.append_reveal_script(
        script::Builder::new()
        .push_slice(&public_key.serialize())
        .push_opcode(opcodes::all::OP_CHECKSIG),
    );

    // 5） 构造 taproot utxo 的花费交易， 关于taproot细节，不必西九
    let taproot_spend_info = TaprootBuilder::new()
        .add_leaf(0, reveal_script.clone())
        .expect("adding leaf should work")
        .finalize(&secp256k1, public_key)
        .expect("finalizing taproot builder should work");

    let control_block = taproot_spend_info
        .control_block(&(reveal_script.clone(), LeafVersion::TapScript))
        .expect("should compute control block");

    // 6） 根据上一步的信息，生产一个临时地址，接收包含 reveal脚本  的 交易输出(TXO)
    let commit_tx_address = Address::p2tr_tweaked(taproot_spend_info.output_key(), network);

    // 7) 根据交易字节数计算 reveal_tx 所需要的 手续费
    let (_, reveal_fee) = Self::build_reveal_transaction(
        &control_block,
        reveal_fee_rate,
        OutPoint::null(), // 并非空，而是 有字节数的，这样才能计算手续费费用
        TxOut {
        script_pubkey: destination.script_pubkey(),
        value: 0,
        },
        &reveal_script,
    );

    // 8） 因为 需要输出一个包含reveal脚本的 TXO, 所以， 其中一个 TXO是用于后面的 reveal操作的
    let unsigned_commit_tx = TransactionBuilder::build_transaction_with_value(
        satpoint,
        inscriptions,
        utxos,
        commit_tx_address.clone(),
        change,
        commit_fee_rate,

        // reveal交易手续费  +  铭文UTXO占位金额TARGET_POSTAGE，一般是 10000sat, 即 0.00010000 BTC
        // 为什么是  0.00010000 BTC ?   不可以是 0.00000546?
        reveal_fee + TransactionBuilder::TARGET_POSTAGE,
    )?;

    // 9） 获取交易输出大小，以及 交易输出， 用作构造 reveal_tx
    let (vout, output) = unsigned_commit_tx
        .output
        .iter()
        .enumerate()
        .find(|(_vout, output)| output.script_pubkey == commit_tx_address.script_pubkey())
        .expect("should find sat commit/inscription output");

    // 10) 根据 上一步的 commit_tx 的交易输出， 构造 reveal_tx
    let (mut reveal_tx, fee) = Self::build_reveal_transaction(
        &control_block,
        reveal_fee_rate,
        OutPoint {
        txid: unsigned_commit_tx.txid(),
        vout: vout.try_into().unwrap(),
        },
        TxOut {
        script_pubkey: destination.script_pubkey(),
        value: output.value, // 暂时用这个，下一步会修改
        },
        &reveal_script,
    );

    // 11) 修改 reveal_tx 的交易输出金额 为  value - fee , 正常来说修改后的交易输出金额为 TransactionBuilder::TARGET_POSTAGE
    reveal_tx.output[0].value = reveal_tx.output[0]
        .value
        .checked_sub(fee.to_sat())
        .context("commit transaction output value insufficient to pay transaction fee")?;

    // 12) 判断是否为 dust（尘埃）交易
    if reveal_tx.output[0].value < reveal_tx.output[0].script_pubkey.dust_value().to_sat() {
        bail!("commit transaction output would be dust");
    }

    // 13） 生成用于签名的hash
    let mut sighash_cache = SighashCache::new(&mut reveal_tx);

    let signature_hash = sighash_cache
        .taproot_script_spend_signature_hash(
        0,
        &Prevouts::All(&[output]),
        TapLeafHash::from_script(&reveal_script, LeafVersion::TapScript),
        SchnorrSighashType::Default,
        )
        .expect("signature hash should compute");

    // 14） 使用 第 3 步中的 “临时”密钥，对上一步生成的hash进行  schnorr签名
    let signature = secp256k1.sign_schnorr(
        &secp256k1::Message::from_slice(signature_hash.as_inner())
        .expect("should be cryptographically secure hash"),
        &key_pair,
    );

    // 15) 将 上一步生成的签名放到 见证数据中
    let witness = sighash_cache
        .witness_mut(0)
        .expect("getting mutable witness reference should work");
    witness.push(signature.as_ref());
    witness.push(reveal_script);
    witness.push(&control_block.serialize());

    let recovery_key_pair = key_pair.tap_tweak(&secp256k1, taproot_spend_info.merkle_root());

    let (x_only_pub_key, _parity) = recovery_key_pair.to_inner().x_only_public_key();
    assert_eq!(
        Address::p2tr_tweaked(
        TweakedPublicKey::dangerous_assume_tweaked(x_only_pub_key),
        network,
        ),
        commit_tx_address
    );

    let reveal_weight = reveal_tx.weight();

    if !no_limit && reveal_weight > MAX_STANDARD_TX_WEIGHT.try_into().unwrap() {
        bail!(
        "reveal transaction weight greater than {MAX_STANDARD_TX_WEIGHT} (MAX_STANDARD_TX_WEIGHT): {reveal_weight}"
        );
    }

    Ok((unsigned_commit_tx, reveal_tx, recovery_key_pair))
}



//=================
pub(crate) fn append_reveal_script(&self, builder: script::Builder) -> Script {
    self.append_reveal_script_to_builder(builder).into_script()
}

fn append_reveal_script_to_builder(&self, mut builder: script::Builder) -> script::Builder {
    // 参考： https://docs.ordinals.com/inscriptions.html

    builder = builder
      .push_opcode(opcodes::OP_FALSE)
      .push_opcode(opcodes::all::OP_IF)
      .push_slice(PROTOCOL_ID);

    if let Some(content_type) = &self.content_type {
      builder = builder
        .push_slice(CONTENT_TYPE_TAG)
        .push_slice(content_type);
    }

    if let Some(body) = &self.body {
      builder = builder.push_slice(BODY_TAG);
      for chunk in body.chunks(520) {
        builder = builder.push_slice(chunk);
      }
    }

    builder.push_opcode(opcodes::all::OP_ENDIF)
  }
//=====================

```