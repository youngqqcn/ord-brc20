# ord-brc20
ordinal and brc20

- Taproot分析
- BTC生态全景
- 主流平台和钱包分析
- ordinal协议分析
  - https://github.com/casey/ord/
  - https://docs.ordinals.com/
    - 自己部署铭刻：https://docs.ordinals.com/guides/inscriptions.html
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