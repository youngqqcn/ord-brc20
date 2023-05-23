# Script to generate a new block every minute
# Put this script at the root of your unpacked folder
#!/bin/bash

echo "Generating a block every minute. Press [CTRL+C] to stop.."

# address=`bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 getnewaddress`

while :
do
        echo "Generate a new block `date '+%d/%m/%Y %H:%M:%S'`"
        bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808  generatetoaddress 100 bcrt1qqv33jfmpethnrcr6xrdxtleszv7lhufvx7p6yn
        sleep 60
done