package modules

import (
	_ "fmt"
	"time"
)

type CurrencyContainer struct {
	Histories map[string]CurrencyHistory
	Updated   string
}

func NewCurrencyContainer() CurrencyContainer {
	cont := CurrencyContainer{}
	cont.Histories = make(map[string]CurrencyHistory)
	cont.Updated = ""

	return cont
}

func (cont *CurrencyContainer) Append(currency CryptoCurrency) CurrencyHistory {
	hist, ok := cont.Histories[currency.Symbol]

	if !ok {
		cont.Histories[currency.Symbol] = NewCurrencyHistory(currency.Symbol)
		hist = cont.Histories[currency.Symbol]
	}

	hist.Append(currency)
	cont.Histories[currency.Symbol] = hist

	cont.MarkAsUpdated()

	return hist
}

func (cont *CurrencyContainer) MarkAsUpdated() *CurrencyContainer {
	cont.Updated = time.Now().Format("2006-01-02 15:04:05")

	return cont
}

func (cont *CurrencyContainer) TotalCashValue() float64 {
	totalValue := 0.0

	for key := range cont.Histories {
		hist := cont.Histories[key]
		totalValue += hist.Last().CashValue
	}

	return totalValue
}

func (cont *CurrencyContainer) Depth() int {
	depth := 0

	for key := range cont.Histories {
		hist := cont.Histories[key]
		depth = hist.Depth()
	}

	return depth
}
