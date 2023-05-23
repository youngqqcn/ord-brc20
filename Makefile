regtest:
	bitcoind -conf=$(PWD)/bitcoin-regtest.conf

index:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie index

balances:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 getbalance