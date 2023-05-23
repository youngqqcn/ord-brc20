# ord-brc20
ordinal and brc20

- Taproot分析
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

- BTC NFT生态分析
  - BTC NFT API: https://docs.nftscan.com/reference/btc/model/asset-model
- BRC20分析
- 工具平台
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