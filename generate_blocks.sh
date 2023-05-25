# Script to generate a new block every minute
# Put this script at the root of your unpacked folder
#!/bin/bash

echo "Generating a block every minute. Press [CTRL+C] to stop.."

# address=`bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 getnewaddress`

while :
do
        echo "Generate a new block `date '+%d/%m/%Y %H:%M:%S'`"
        bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808  generatetoaddress 110 bcrt1p3m68zq7vg7j9hey63uasw0ejr3vsgg862luy7af64c8e4tar75gqxxcyvj
        sleep 100
done