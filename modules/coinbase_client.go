package modules

import (
	"os"
	"time"

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

func (client *CoinbaseClient) Refresh() {
	accounts, err := client.Client.Accounts()
	if err != nil {
		return
	}

	client.Update(accounts)
}

func (client *CoinbaseClient) Update(accounts coinbase.APIAccounts) {
	for key := range client.Container.Currencies {
		client.Container.Currencies[key] = 0.0
	}

	client.Container.TotalValue = 0.0
	client.Container.Updated = time.Now().Format("2006-01-02 15:04:05")

	for _, account := range accounts.Data {
		client.Container.TotalValue += account.Native_balance.Amount

		client.Container.Currencies[account.Balance.Currency] += account.Native_balance.Amount
	}
}
