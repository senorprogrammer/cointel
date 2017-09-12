package modules

import (
	_ "fmt"
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
	/*
	* First sum up currencies of like types to get a single value for them
	* For instance, if a batch of accounts has multiple BTC accounts, then this
	* batches their values into one CryptoCurrency instance that contains the
	* sum total of the parts (which makes the data suitable for being passed along
	* to CurrencyContainer)
	 */
	currencies := make(map[string]CryptoCurrency)

	for _, account := range accounts.Data {
		curr, ok := currencies[account.Balance.Currency]

		if !ok {
			currencies[account.Balance.Currency] = CryptoCurrency{}
			curr = currencies[account.Balance.Currency]
			curr.Symbol = account.Balance.Currency
		}

		curr.Quantity += account.Balance.Amount
		curr.CashValue += account.Native_balance.Amount

		currencies[account.Balance.Currency] = curr
	}

	for _, curr := range currencies {
		client.Container.Append(curr)
	}
}
