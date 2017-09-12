package modules

import (
	"time"
)

type CurrencyHistory struct {
	CryptoCurrencies []CryptoCurrency
	Symbol           string
	Updated          string
}

func NewCurrencyHistory(symbol string) CurrencyHistory {
	hist := CurrencyHistory{}
	hist.Symbol = symbol
	hist.Updated = ""

	return hist
}

func (hist *CurrencyHistory) Last() *CryptoCurrency {
	curr := hist.CryptoCurrencies[len(hist.CryptoCurrencies)-1]
	return &curr
}

func (hist *CurrencyHistory) Append(curr CryptoCurrency) *CurrencyHistory {
	hist.CryptoCurrencies = append(hist.CryptoCurrencies, curr)
	hist.MarkAsUpdated()
	return hist
}

func (hist *CurrencyHistory) MarkAsUpdated() *CurrencyHistory {
	hist.Updated = time.Now().Format("2006-01-02 15:04:05")
	return hist
}

func (hist *CurrencyHistory) Depth() int {
	return len(hist.CryptoCurrencies)
}
