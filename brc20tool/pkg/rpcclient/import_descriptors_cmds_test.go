package rpcclient

import (
	"fmt"
	"log"
	"testing"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/rpcclient"
)

func TestImportDescriptorsCmds(t *testing.T) {
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
	// net := &chaincfg.RegressionNetParams

	// privateKey, err := btcec.NewPrivateKey()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// privateKeyWIF, err := btcutil.NewWIF(privateKey, net, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Printf("=================\n")

	// btcApiClient := mempool.NewClient(netParams)
	// privateKeyWIF, err := btcutil.DecodeWIF("cS4bEaUoFkWM5qRaPXzGTmUje73b5zDkbamXDv5SuMWCM3fHJnyy")
	privateKeyWIF, err := btcutil.DecodeWIF("cVHTRk2g4YFiWXufCLJ8ZV2KVqLaqHqksKg3Ay8wRRztJFSEJHto")
	if err != nil {
		log.Printf("fffffffffffffffffff\n")
		log.Fatal(err)
		return
	}
	fmt.Printf("wif: %v", privateKeyWIF.String())

	descriptorInfo, err := client.GetDescriptorInfo(fmt.Sprintf("rawtr(%s)", privateKeyWIF))
	if err != nil {
		log.Fatal(err)
	}

	// nextIndex := 0
	// rangeN := []int{0, 100000}
	descriptors := []Descriptor{
		{
			Desc: *btcjson.String(fmt.Sprintf("rawtr(%s)#%s", privateKeyWIF, descriptorInfo.Checksum)),
			Timestamp: btcjson.TimestampOrNow{
				Value: "now",
			},
			Active:    btcjson.Bool(false),
			Range:     nil,
			NextIndex: nil,
			Internal:  btcjson.Bool(true),
			// Label:     btcjson.String("yqq"),
		},
	}

	log.Printf("desc: %v", *btcjson.String(fmt.Sprintf("rawtr(%s)#%s", privateKeyWIF, descriptorInfo.Checksum)))

	// cmd := &ImportDescriptorsCmd{
	// 	Descriptors: descriptors,
	// }

	// marshalledJSON, err := btcjson.MarshalCmd("jsonrpc1.0", 111, cmd)
	// log.Printf("%v\n", marshalledJSON)

	results, err := ImportDescriptors(client, descriptors)
	if err != nil {
		log.Fatal(err)
	}
	if results == nil {
		log.Fatalf("import failed, nil result")
	}
	for _, result := range *results {
		if !result.Success {
			log.Fatal(fmt.Errorf("import failed:%v", result.Error.Message))
		}
	}
	log.Printf("Import descriptors success.")
}
