regtest:
	bitcoind -conf=$(PWD)/bitcoin-regtest.conf

ord-index:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie index

balances:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 getbalance

ord-wallet:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet create

ord-receive:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet receive

ord-transactions:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet transactions

ord-server:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie server --http-port 8888


ord-inscribe:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet inscribe --fee-rate 10 ./inscription_content.txt

ord-inscriptions:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet inscriptions

newaddress:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=descriptors  getnewaddress test1 bech32m


loadwallet:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=descriptors loadwallet descriptors


getbalances:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=descriptors getbalance

getaddressesbylabel:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=descriptors getaddressesbylabel test1

#bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=descriptors -named sendtoaddress address="bcrt1px4ffhmxsmcdzjqkcmd3e3nec0n8ha9z3cqzfnxwhske5w3pmx0gs6ygf3m" amount=0.5 fee_rate=25
send:
	 bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=test1 -named sendtoaddress address="bcrt1p3m68zq7vg7j9hey63uasw0ejr3vsgg862luy7af64c8e4tar75gqxxcyvj" amount=100 fee_rate=25

generate_new_block:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808  generatetoaddress 1 bcrt1qqv33jfmpethnrcr6xrdxtleszv7lhufvx7p6yn

gettransaction:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 getrawtransaction abb5f352fa76cfd1c7160a64ee439fab4d1f13d1d1154310ca167933a64c8702 true