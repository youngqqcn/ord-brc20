# ord-brc20
ordinal and brc20

- 铭刻教程：
  - https://www.panewslab.com/zh/articledetails/27ojhj8a.html

- Taproot升级
  - https://gamma.io/learn/blockchain/bitcoin/taproot-upgrade
- Tapscript
  - https://docs.ordinals.com/inscriptions.html
  - 深入理解： https://medium.com/@bun919tw/%E6%B7%B1%E5%85%A5%E7%90%86%E8%A7%A3-bitcoin-nft-ordinals-3811b0eb9fed
  - 深度解析：https://mirror.xyz/quentangle.eth/zRV-TCg62FGhee89fTgAYUqywcc6x9wpTG6gVdMrEp0
    - **由于taproot脚本的花费只能从现有的taproot输出中进行**
      - 第一步： 将铭文写进taproot的交易输出（TXO）作为UTXO，这一步叫做commit
      - 第二步： 花掉第一步产生的UTXO，铭文内容被揭示（如果不花费UTXO,那么铭文内容一直在UTXO中，不能算真正的上链）
  - ordinals的tapscript脚本
    ```
    <signature>
    OP_FALSE
    OP_IF
        OP_PUSH "ord"
        OP_1
        OP_PUSH "text/plain;charset=utf-8"
        OP_0
        OP_PUSH "Hello, world!"
    OP_ENDIF
    <public key>
    ```
  - `OP_FALSE` 會 push 一個 empty array 到 stack，注意這邊是有 push 東西的，只是它是空的。
  -  `OP_IF` 檢查 stack 頂部，如果為 true 才會做接下來的事情，因為前面 OP_FALSE 的動作，導致這個 if 不會成立。
    接下來 `OP_PUSH` … 等一系列操作都會被忽略，因為上一個 if 條件沒有達成。
  - `OP_ENDIF` 結束這個 if 區塊。
  - 可以看出來中間這些操作因為 `OP_IF` **一定不會成立**，所以等於什麼狀態都沒改變，於是就可以把圖片的完整資料都放在 OP_IF 裡面而**不影響本來 Bitcoin script 的 validation**，多虧了 taproot 升級，script 現在是沒有大小上限了，所以只要 transaction 的大小小於 block 的大小 (4 MB)，script 你要多大都可以，也就是說我們可以達到**類似 OP_RETURN 的效果**，把無關的資料放上去 Bitcoin 卻還沒有 80 bytes 的大小限制了。
  - 其中 `OP_0`后面跟随的是incribe的内容，每个块不能超过`520 bytes`

  - https://medium.com/@bun919tw/%E6%B7%B1%E5%85%A5%E7%90%86%E8%A7%A3-bitcoin-nft-ordinals-3811b0eb9fed
- BTC生态全景
- 主流平台和钱包分析
- ordinal协议分析
  - 源码解析： https://github.com/unisat-wallet/unisat-docs/blob/master/docs/guide/unisat-api.md
  - https://unchainedcrypto.com/how-to-create-a-bitcoin-ordinal/
  - https://github.com/casey/ord/
  - https://docs.ordinals.com/
    - 自己部署铭刻：https://docs.ordinals.com/guides/inscriptions.html
    - https://learnblockchain.cn/article/5376
  - ordinal主网浏览器： https://ordinals.com/
  - 搭建节点： https://zhuanlan.zhihu.com/p/612394795
  - 测试网： https://signet.ordinals.com/
  - regtest: https://github.com/casey/ord/issues/1638
    - regtest使用：  https://gist.github.com/System-Glitch/cb4e87bf1ae3fec9925725bb3ebe223a
- 代码实现
  - https://learnblockchain.cn/article/5782
- BTC NFT生态分析
  - BTC NFT API: https://docs.nftscan.com/reference/btc/model/asset-model
- BRC20分析
- 工具平台
  - https://gamma.io/collections
  - https://ordinalswallet.com/
  - https://unisat.io
  - https://idclub.io/index

- 开源铭刻：
  - https://looksordinal.com/
    - https://medium.com/@rarity.garden/inscriptions-with-looksordinal-cfa2d635f720

- 区块浏览器：
  - https://mempool.space/zh/

- 解决 ord  连接 bitcoin rpc 400错误：
  - truncate -s -1 ./bitcoin.cookie


```
bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 createwallet test1

bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 getnewaddress
```



mnemonic:
```
wish daring pottery stuff project laundry finish impact mind hover actress slogan
```

receive
{
  "address": "bcrt1px4ffhmxsmcdzjqkcmd3e3nec0n8ha9z3cqzfnxwhske5w3pmx0gs6ygf3m"
}

bcrt1qkd522aj7hyevsa7mv2893wqtm7ey0v6se83fzn


```
$ make ordinscribe
ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet inscribe --fee-rate 10 ./helloworld.txt
{
  "commit": "a62c8132698e80c96ef1db32230080790b678cd49c8b3f4ad1143687a8d2c3ab",
  "inscription": "88057a76b3b5b58c80162785472f99f52906495111c4774480838f56c6edcd5ai0",
  "reveal": "88057a76b3b5b58c80162785472f99f52906495111c4774480838f56c6edcd5a",
  "fees": 2940
}
```



----


bitcoin testnet  faucet

https://testnet-faucet.com/btc-testnet/




----

# BitcoinEagle技术方案设计

## 账户模型设计

基于BIP44为每个用户（订单）生成唯一的 P2TR格式的BTC充币地址


## 如何检测用户充值？

调用mempool的API接口扫描区块，匹配地址

或者直接监听内存池中的交易

> 如果时间允许，可以获取内存中的交易，提升用户体验


## 如果铭刻？

- 需要记录每笔订单的费率，以及铭刻内容，异步铭刻
- 当收到用户的充值后 直接使用用户充值交易的那个`UTXO`构造`commit_tx`,并用`commit_tx`的输出构造`reveal_tx`
- 充值交易在内存中，是否可以直接铭刻，按理说是可以的！


## 如何抽水？

构造`commit_tx`的时候，插入我们的抽水输出即可



# 模块划分

- 后台任务模块：
  - BTC地址生成模块，基于BIP44规范
  - BTC区块（内存池）扫描模块，监听用户充值交易
  - 异步铭刻模块

API服务：
  - 预约接口（已完成）
  - 链上实施费率查询（前端定时直接查mempool的接口）
    - https://mempool.space/api/v1/fees/recommended
  - 费用估算接口（服务费）
    - 返回比特币价格
  - 订单创建接口
    - 返回
      - 订单id,
      - 充值地址： 前端根据充值地址生成二维码
      - btc价格
      - 字节大小
      - 铭刻费用
      - 服务费用
      - 总费用
      - 总费用（美元）

  - 获取用户接收地址的详情：
    - 提示用户： 请不要使用存在未确认转出交易的账户进行付款
    - https://blockstream.info/api/address/bc1phjsyw73de6ap8nfjzg4erxmdw7lzlfgvm447v82fytn78nm0mwnsq654e7

  - 订单查询接口
    - 根据订单号查询
    - 或根据接收地址查询
    - 或根据付款地址查询

  - 盲盒查询接口
    - 盲盒id
    - 盲盒名称
    - 描述
    - 价格(BTC)
    - 价格(USD)
    - 支付币种
    - 总数量
    - 剩余数量
