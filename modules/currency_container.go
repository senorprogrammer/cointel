package modules

import (
	// "fmt"
	"time"
)

/*
* CurrencyContainer is a struct that holds instances of Currency.
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

func (cont *CurrencyContainer) ZeroOut() {
	for key := range cont.Currencies {
		curr := cont.Currencies[key]
		curr.ZeroOut()
	}
}

/*
* Adds a sum to the given currency in the container
* If the currency doesn't exist, create a new instance of CryptoCurrency and
* populate that. If it does, add it to the existing
 */
func (cont *CurrencyContainer) AddToCurrency(symbol string, quantity float64, cashValue float64) {
	curr := cont.Currencies[symbol]

	/* If the currency does not yet exist (compared to the struct's zero value),
	* create it, store it, and assign it
	 */
	if curr == (CryptoCurrency{}) {
		cont.Currencies[symbol] = NewCryptoCurrency(symbol)
		curr = cont.Currencies[symbol]
	}

	curr.Quantity += quantity
	curr.CashValue += cashValue
}

func (cont *CurrencyContainer) MarkAsUpdated() {
	cont.Updated = time.Now().Format("2006-01-02 15:04:05")
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
