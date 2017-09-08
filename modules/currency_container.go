package modules

import (
	"time"
)

/*
* CurrencyContainer is a struct that holds instances of CryptoCurrency.
* It has the ability to update currencies with new data, and
* provide a cash value total of their worth
 */
type CurrencyContainer struct {
	Currencies map[string]CryptoCurrency
	Updated    string
}

func NewCurrencyContainer() CurrencyContainer {
	cont := CurrencyContainer{}
	cont.Currencies = make(map[string]CryptoCurrency)
	cont.Updated = ""

	return cont
}

/*
* Creates a new set of CryptoCurrency entries with empty values ready to be populated with
* sweet, sweet data numbers
 */
func (cont *CurrencyContainer) NewVersion() *CurrencyContainer {
	for _, curr := range cont.Currencies {
		nextCurr := NewCryptoCurrency(curr.Symbol, &curr, curr.Index+1)
		curr.Next = &nextCurr

		cont.Currencies[curr.Symbol] = nextCurr
	}

	return cont
}

/*
* Adds a sum to the given currency in the container
* If the currency doesn't exist, create a new instance of CryptoCurrency and
* populate that. If it does, add it to the existing
 */
func (cont *CurrencyContainer) AddToCurrency(symbol string, quantity, cashValue float64) CryptoCurrency {
	curr, ok := cont.Currencies[symbol]

	if !ok {
		cont.Currencies[symbol] = NewCryptoCurrency(symbol, &curr, 0)
		curr = cont.Currencies[symbol]
	}

	curr.Add(quantity, cashValue)
	curr.MarkAsUpdated()

	cont.Currencies[symbol] = curr

	return curr
}

func (cont *CurrencyContainer) MarkAsUpdated() *CurrencyContainer {
	cont.Updated = time.Now().Format("2006-01-02 15:04:05")

	return cont
}

// The sum of the CashValues of all CryptoCurrencies
func (cont *CurrencyContainer) TotalCashValue() float64 {
	totalValue := 0.0

	for key := range cont.Currencies {
		curr := cont.Currencies[key]
		totalValue += curr.CashValue
	}

	return totalValue
}
