package modules

import (
	"time"

	"github.com/Zauberstuhl/go-coinbase"
)

type CryptoCurrencies struct {
	Currencies map[string]float64
	TotalValue float64
	Updated    string
}

func NewCryptoCurrencies() CryptoCurrencies {
	cc := CryptoCurrencies{}
	cc.Currencies = make(map[string]float64)
	cc.TotalValue = 0.0
	cc.Updated = ""

	return cc
}

func (cc *CryptoCurrencies) CoinbaseUpdate(accounts *coinbase.APIAccounts) {
	for key := range cc.Currencies {
		cc.Currencies[key] = 0.0
	}

	cc.TotalValue = 0.0
	cc.Updated = time.Now().Format("2006-01-02 15:04:05")

	for _, account := range accounts.Data {
		cc.TotalValue += account.Native_balance.Amount

		cc.Currencies[account.Balance.Currency] += account.Native_balance.Amount
	}
}
