package modules

import (
	"os"

	"github.com/Zauberstuhl/go-coinbase"
)

type CoinbaseClient struct {
	Client    coinbase.APIClient
	Container CurrencyContainer
}

func NewCoinbaseClient() CoinbaseClient {
	cc := CoinbaseClient{}
	cc.Client = coinbase.APIClient{
		Key:    os.Getenv("COINBASE_KEY"),
		Secret: os.Getenv("COINBASE_SECRET"),
	}
	cc.Container = NewCurrencyContainer()

	return cc
}

func (cc *CoinbaseClient) Refresh() {
	accounts, err := cc.Client.Accounts()
	if err != nil {
		return
	}

	cc.Container.CoinbaseUpdate(&accounts)
}
