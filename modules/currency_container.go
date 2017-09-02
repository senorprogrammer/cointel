package modules

import ()

type CurrencyContainer struct {
	Currencies map[string]float64
	TotalValue float64
	Updated    string
}

func NewCurrencyContainer() CurrencyContainer {
	cc := CurrencyContainer{}
	cc.Currencies = make(map[string]float64)
	cc.TotalValue = 0.0
	cc.Updated = ""

	return cc
}
