package modules

import (
	"os"

	"github.com/Zauberstuhl/go-coinbase"
)

type CoinbaseClient struct {
	Client    coinbase.APIClient
	Container CurrencyContainer
	Label     string
}

func NewCoinbaseClient() CoinbaseClient {
	client := CoinbaseClient{}
	client.Client = coinbase.APIClient{
		Key:    os.Getenv("COINBASE_KEY"),
		Secret: os.Getenv("COINBASE_SECRET"),
	}
	client.Container = NewCurrencyContainer()
	client.Label = "Coinbase"

	return client
}

func (client *CoinbaseClient) Refresh() {
	accounts, err := client.Client.Accounts()

	if err != nil {
		return
	}

	client.Update(accounts)
}

func (client *CoinbaseClient) Update(accounts coinbase.APIAccounts) {
	client.Container.ZeroOut()

	for _, account := range accounts.Data {
		client.Container.AddToCurrency(
			account.Balance.Currency,
			account.Balance.Amount,
			account.Native_balance.Amount,
		)
	}

	client.Container.MarkAsUpdated()
}
