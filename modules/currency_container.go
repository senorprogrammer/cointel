package modules

import ()

type CurrencyContainer struct {
	Currencies     map[string]float64
	TotalCashValue float64
	Updated        string
}

func NewCurrencyContainer() CurrencyContainer {
	cc := CurrencyContainer{}
	cc.Currencies = make(map[string]float64)
	cc.TotalCashValue = 0.0
	cc.Updated = ""

	return cc
}

func (cont *CurrencyContainer) ZeroOut() {
	for key := range cont.Currencies {
		cont.Currencies[key] = 0.0
	}

	cont.TotalCashValue = 0.0
}
