regtest:
	bitcoind -conf=$(PWD)/bitcoin-regtest.conf

ord-index:
	ord --regtest --bitcoin-data-dir=/home/yqq/mine/ord-brc20/data/  --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie index

balances:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 getbalance

ord-wallet:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet create

ord-receive:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet receive

ord-transactions:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet transactions

ord-server:
	ord --regtest --bitcoin-data-dir=/home/yqq/mine/ord-brc20/data/ --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie server --http-port 8888


ord-inscribe:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet inscribe --fee-rate 10 ./inscription_content.txt

ord-inscriptions:
	ord --regtest --cookie-file /home/yqq/mine/ord-brc20/bitcoin.cookie wallet inscriptions

createwallet:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=yqq  -named createwallet wallet_name=yqq avoid_reuse=true descriptors=true load_on_startup=true

newaddress:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=yqq  getnewaddress test1 bech32m


loadwallet:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=yqq loadwallet yqq


getbalances:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=yqq getbalance

getaddressesbylabel:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=yqq getaddressesbylabel test1

#bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=yqq -named sendtoaddress address="bcrt1px4ffhmxsmcdzjqkcmd3e3nec0n8ha9z3cqzfnxwhske5w3pmx0gs6ygf3m" amount=0.5 fee_rate=25
#  bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=yqq -named sendtoaddress address="bcrt1p3m68zq7vg7j9hey63uasw0ejr3vsgg862luy7af64c8e4tar75gqxxcyvj" amount=100 fee_rate=25
send:
	 bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=yqq -named sendtoaddress address="bcrt1p78vllj6tchpe0tsf3pg3t33eyha5fv04qangma8njwdv2lewftpq3purje" amount=10 fee_rate=25

generate_new_block:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808  generatetoaddress 1 bcrt1qqv33jfmpethnrcr6xrdxtleszv7lhufvx7p6yn

gettransaction:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 getrawtransaction abb5f352fa76cfd1c7160a64ee439fab4d1f13d1d1154310ca167933a64c8702 true


getblockchaininfo:
	curl -s --user qiyihuo:qiyihuo1808  --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockchaininfo", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:18443/



getblock:
	curl -s --user qiyihuo:qiyihuo1808  --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblock", "params": ["2b24119c63a7f4fd218b944f49552fbdbae095201e66b0c10d35e166387c3a66"]}' -H 'content-type: text/plain;' http://127.0.0.1:18443/


#bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808  importdescriptors '[{ "desc": "rawtr(cS4bEaUoFkWM5qRaPXzGTmUje73b5zDkbamXDv5SuMWCM3fHJnyy)#aj0gxagn",  "active":false, "timestamp":"now", "internal": true }]'
importdescriptors:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808  importdescriptors '[{ "desc": "rawtr(cVHTRk2g4YFiWXufCLJ8ZV2KVqLaqHqksKg3Ay8wRRztJFSEJHto)#4las3lll",  "active":false, "timestamp":"now", "internal": true }]'

getdescriptorinfo:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808  getdescriptorinfo "rawtr(SmTDjW72t5BQ4UuNXJzCQwZSeTNcTXL8RSeC4ybqhDyGiizZz5n)"

listunspent:
	 bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=yqq listunspent 0 1999999 [\"bcrt1p27gduney7a3pxl3wqc3p9rzy9q4ew5a43du8eznat7scpqnuaf0s0pxnq0\"]


cleanindex:
	rm -rf ~/.local/share/ord/regtest/

listdescriptors:
	bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 -rpcwallet=yqq  listdescriptors