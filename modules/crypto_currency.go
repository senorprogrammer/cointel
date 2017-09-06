package modules

import (
	"time"
)

/*
* Currency represents a cryptocurrency, like Bitcoin or Ethereum
* It stores the number of units of that currency held, and their
* cash value in whatever the native fiat currency is
 */

type CryptoCurrency struct {
	Quantity  float64
	CashValue float64
	Symbol    string
	Updated   string
}

func NewCryptoCurrency(symbol string) CryptoCurrency {
	curr := CryptoCurrency{}
	curr.Quantity = 0.0
	curr.CashValue = 0.0
	curr.Symbol = symbol
	curr.Updated = ""

	return curr
}

func (curr *CryptoCurrency) ZeroOut() *CryptoCurrency {
	curr.Quantity = 0.0
	curr.CashValue = 0.0

	return curr
}

func (curr *CryptoCurrency) Add(quantity float64, cashValue float64) *CryptoCurrency {
	if quantity >= 0 {
		curr.Quantity += quantity
	}

	if cashValue >= 0 {
		curr.CashValue += cashValue
	}

	return curr
}

func (curr *CryptoCurrency) MarkAsUpdated() *CryptoCurrency {
	curr.Updated = time.Now().Format("2006-01-02 15:04:05")

	return curr
}
