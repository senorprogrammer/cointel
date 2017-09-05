package modules

import ()

type CurrencyContainer struct {
	Currencies map[string]float64
	TotalValue float64
	Updated    string
}

func NewCurrencyContainer() CurrencyContainer {
	cont := CurrencyContainer{}
	cont.Currencies = make(map[string]float64)
	cont.Updated = ""

	cont.ZeroOut()

	return cont
}

func (cont *CurrencyContainer) ZeroOut() {
	for key := range cont.Currencies {
		cont.Currencies[key] = 0.0
	}

	cont.TotalValue = 0.0
}
