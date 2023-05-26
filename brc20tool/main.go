package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	btcrpcclient "github.com/btcsuite/btcd/rpcclient"
	"github.com/minchenzz/brc20tool/internal/ord"
	"github.com/minchenzz/brc20tool/pkg/rpcclient"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

var (
	gwif    string // WIF private key
	gop     string // brc20 mint or transfer
	gtick   string // brc20 tick
	gamount string // brc20 amount
	grepeat string //  count
	gsats   string // feeRate    n sat/vbytes
)

func main() {
	//  bcrt1p4anml5s767csvrhwm2lehx9h2wyeqnj9gazdrxhygag89fruz8eqyjetzt
	// gwif = "cS4bEaUoFkWM5qRaPXzGTmUje73b5zDkbamXDv5SuMWCM3fHJnyy"

	// bcrt1p78vllj6tchpe0tsf3pg3t33eyha5fv04qangma8njwdv2lewftpq3purje
	gwif = "cVHTRk2g4YFiWXufCLJ8ZV2KVqLaqHqksKg3Ay8wRRztJFSEJHto"

	gop = "mint"
	gtick = "EAGLE"
	gamount = "1000"
	grepeat = "1"
	gsats = "25"

	run(false)

	// gen_address()
}

func gen_address() {

	// hexPrivateKey := "32000c4bbe088e517efe41d1c4e1da1cf05dbc9268ff53c8b1360a8d1455426c"

	if true {
		netParams := &chaincfg.RegressionNetParams
		privateKey, _ := btcec.NewPrivateKey()
		wifPrivKey, _ := btcutil.NewWIF(privateKey, netParams, true)
		fmt.Printf("wif compressed private key:%v\n", wifPrivKey.String())

		fmt.Printf("wif compressed public key:%v\n", wifPrivKey.SerializePubKey())

		utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(wifPrivKey.PrivKey.PubKey())), netParams)
		if err != nil {
			return
		}
		fmt.Printf(" address: %v\n ", utxoTaprootAddress.String())
	} else {

		netParams := &chaincfg.RegressionNetParams

		wifKey, err := btcutil.DecodeWIF(gwif)
		if err != nil {
			return
		}
		utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(wifKey.PrivKey.PubKey())), netParams)
		if err != nil {
			return
		}
		fmt.Printf(" address: %v\n ", utxoTaprootAddress.String())
	}

}

func run(forEstimate bool) (txid string, txids []string, fee int64, err error) {
	// netParams := &chaincfg.MainNetParams

	netParams := &chaincfg.RegressionNetParams
	// btcApiClient := mempool.NewClient(netParams)
	wifKey, err := btcutil.DecodeWIF(gwif)
	if err != nil {
		return
	}
	utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(wifKey.PrivKey.PubKey())), netParams)
	if err != nil {
		return
	}
	// unspentList, err := btcApiClient.ListUnspent(utxoTaprootAddress)
	unspentList, err := rpcclient.ListUnspent(utxoTaprootAddress)
	if err != nil {
		return
	}

	fmt.Printf("utxo size is %v\n", len(unspentList))

	// return

	// if len(unspentList) == 0 {
	// 	err = fmt.Errorf("no utxo for %s", utxoTaprootAddress)
	// 	return
	// }

	vinAmount := 0
	commitTxOutPointList := make([]*wire.OutPoint, 0)
	commitTxPrivateKeyList := make([]*btcec.PrivateKey, 0)
	for i := range unspentList {
		if unspentList[i].Output.Value < 10000 {
			continue
		}
		commitTxOutPointList = append(commitTxOutPointList, unspentList[i].Outpoint)
		commitTxPrivateKeyList = append(commitTxPrivateKeyList, wifKey.PrivKey)
		vinAmount += int(unspentList[i].Output.Value)
	}
	fmt.Printf("len(commitTxOutPointList) is %v\n", len(commitTxOutPointList))
	fmt.Printf("len(commitTxPrivateKeyList) is %v\n", len(commitTxPrivateKeyList))

	dataList := make([]ord.InscriptionData, 0)

	// read image from filename
	imgBs, err := ioutil.ReadFile("../eagle-1.png")
	if err != nil {
		fmt.Printf("error:%v\n", err.Error())
		return
	}
	mint := ord.InscriptionData{
		// ContentType: "image/jpeg",
		// ContentType: "image/gif",
		ContentType: "image/png",
		Body:        imgBs,
		Destination: utxoTaprootAddress.EncodeAddress(),
	}

	// mint := ord.InscriptionData{
	// 	ContentType: "text/plain;charset=utf-8",
	// 	Body:        []byte(fmt.Sprintf(`{"p":"brc-20","op":"%s","tick":"%s","amt":"%s"}`, gop, gtick, gamount)),
	// 	Destination: utxoTaprootAddress.EncodeAddress(),
	// }

	count, err := strconv.Atoi(grepeat)
	if err != nil {
		return
	}

	for i := 0; i < count; i++ {
		dataList = append(dataList, mint)
	}

	txFee, err := strconv.Atoi(gsats)
	if err != nil {
		return
	}

	request := ord.InscriptionRequest{
		CommitTxOutPointList:   commitTxOutPointList,
		CommitTxPrivateKeyList: commitTxPrivateKeyList,
		CommitFeeRate:          int64(txFee),
		FeeRate:                int64(txFee),
		DataList:               dataList,
		SingleRevealTxOnly:     false,
	}

	connCfg := &btcrpcclient.ConnConfig{
		// Host:         "localhost:8336",
		Host:         "127.0.0.1:18443",
		User:         "qiyihuo",
		Pass:         "qiyihuo1808",
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	client, err := btcrpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// tool, err := ord.NewInscriptionToolWithBtcApiClient(netParams, btcApiClient, &request)
	tool, err := ord.NewInscriptionTool(netParams, client, &request)
	if err != nil {
		return
	}

	baseFee := tool.CalculateFee()

	if forEstimate {
		fee = baseFee
		return
	}

	commitTxHash, revealTxHashList, _, _, err := tool.Inscribe()
	if err != nil {
		err = fmt.Errorf("send tx errr, %v", err)
		return
	}

	txid = commitTxHash.String()
	fmt.Println(txid)
	for i := range revealTxHashList {
		txids = append(txids, revealTxHashList[i].String())
		fmt.Println(revealTxHashList[i].String())
	}
	fmt.Printf("fee: %v\n", fee)

	return
}

func makeForm(_w fyne.Window) fyne.CanvasObject {
	pk := widget.NewPasswordEntry()

	op := widget.NewEntry()
	op.SetPlaceHolder("op")

	tick := widget.NewEntry()
	tick.SetPlaceHolder("tick")

	amount := widget.NewEntry()
	amount.SetPlaceHolder("amount")

	fee := widget.NewEntry()
	fee.SetPlaceHolder("sats")
	fee.SetText("20")

	repeat := widget.NewEntry()
	repeat.SetPlaceHolder("repeat")
	repeat.SetText("1")

	fees := widget.NewEntry()
	fees.SetPlaceHolder("fees")

	txid := widget.NewEntry()
	txid.SetPlaceHolder("main txid")

	inscribeTxs := widget.NewEntry()
	inscribeTxs.SetPlaceHolder("inscribe txs")
	inscribeTxs.MultiLine = true

	estimate := widget.NewButton("estimate", func() {
		gwif = pk.Text
		gop = op.Text
		gtick = tick.Text
		gamount = amount.Text
		grepeat = repeat.Text
		gsats = fee.Text

		_, _, fee, err := run(true)
		if err != nil {
			dialog.ShowError(err, _w)
			return
		}
		fees.SetText(strconv.FormatInt(fee, 10))
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "private key", Widget: pk, HintText: "Your wif private key"},
			{Text: "op", Widget: op, HintText: "eg: mint, transfer"},
			{Text: "tick", Widget: tick, HintText: "eg: ordi, SHIB"},
			{Text: "amount", Widget: amount, HintText: "eg: 1, 100000"},
			{Text: "sats", Widget: fee, HintText: "eg: 20, 30"},
			{Text: "repeat", Widget: repeat, HintText: "eg: 1, 5, 10"},
			{Text: "estimate fee", Widget: estimate, HintText: "estimate fee(sats)"},
			{Text: "fee", Widget: fees, HintText: "txs fee"},
			{Text: "txid", Widget: txid, HintText: "main txid"},
			{Text: "inscribe txids", Widget: inscribeTxs, HintText: "inscribe txids"},
		},
		OnSubmit: func() {
			gwif = pk.Text
			gop = op.Text
			gtick = tick.Text
			gamount = amount.Text
			grepeat = repeat.Text
			gsats = fee.Text

			_txid, txids, _, err := run(false)
			if err != nil {
				dialog.ShowError(err, _w)
				return
			}
			txid.SetText(_txid)
			txisstr := strings.Join(txids, "\n")
			inscribeTxs.SetText(txisstr)
		},
	}

	return form
}
