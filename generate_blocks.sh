# Script to generate a new block every minute
# Put this script at the root of your unpacked folder
#!/bin/bash

echo "Generating a block every minute. Press [CTRL+C] to stop.."

# address=`bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808 getnewaddress`

while :
do
        echo "Generate a new block `date '+%d/%m/%Y %H:%M:%S'`"
        bitcoin-cli -chain=regtest -rpcuser=qiyihuo -rpcpassword=qiyihuo1808  generatetoaddress 200 bcrt1p27gduney7a3pxl3wqc3p9rzy9q4ew5a43du8eznat7scpqnuaf0s0pxnq0
        sleep 100
done