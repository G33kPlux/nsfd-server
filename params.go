package hdwallet

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// init net params
var (
	BTCParams        = chaincfg.MainNetParams
	BTCTestnetParams = chaincfg.TestNet3Params
	LTCParams        = chaincfg.MainNetParams
	DOGEParams       = chaincfg.MainNetParams
	DASHParams       = chaincfg.MainNetParams
	BCHParams        = chaincfg.MainNetParams
	QTUMParams       = chaincfg.MainNetParams
	USDTParams       = chaincfg.MainNetParams
)

func init() {
	// ltc net params
	// https://g