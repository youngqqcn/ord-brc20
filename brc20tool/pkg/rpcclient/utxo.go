package rpcclient

import (
	"log"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/minchenzz/brc20tool/pkg/btcapi"
)

func ListUnspent(address btcutil.Address) ([]*btcapi.UnspentOutput, error) {

	// net := &chaincfg.RegressionNetParams
	connCfg := &rpcclient.ConnConfig{
		// Host:         "localhost:8336",
		Host:         "127.0.0.1:18443",
		User:         "qiyihuo",
		Pass:         "qiyihuo1808",
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// net := &chaincfg.SigNetParams
	utxos, err := client.ListUnspent()
	if err != nil {
		log.Fatal(err)
	}

	unspentOutputs := make([]*btcapi.UnspentOutput, 0)
	for _, utxo := range utxos {
		txHash, err := chainhash.NewHashFromStr(utxo.TxID)
		if err != nil {
			return nil, err
		}

		unspentOutputs = append(unspentOutputs, &btcapi.UnspentOutput{
			Outpoint: wire.NewOutPoint(txHash, uint32(utxo.Vout)),
			Output:   wire.NewTxOut(int64(utxo.Amount * 10000_0000), address.ScriptAddress()),
		})
	}

	return unspentOutputs, nil
}
