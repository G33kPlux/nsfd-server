package hdwallet

func init() {
	coins[USDT] = newUSDT
}

type usdt struct {
	*btc
}

func newUSDT(key *Key) Wallet {
	token := newBTC(key).(*btc)
	token.name = "Tether"
	token.symbol = "