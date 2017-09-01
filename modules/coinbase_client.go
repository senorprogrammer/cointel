package modules

import (
	"os"

	"github.com/Zauberstuhl/go-coinbase"
)

type CoinbaseClient struct {
	Client           coinbase.APIClient
	CryptoCurrencies *CryptoCurrencies
}

func NewCoinbaseClient() CoinbaseClient {

}

func CoinbaseClient() coinbase.APIClient {
	client := coinbase.APIClient{
		Key:    os.Getenv("COINBASE_KEY"),
		Secret: os.Getenv("COINBASE_SECRET"),
	}

	return client
}

func Update() {

}
