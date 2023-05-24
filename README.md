# ord-brc20
ordinal and brc20

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