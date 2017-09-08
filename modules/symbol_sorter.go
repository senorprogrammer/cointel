package modules

import (
	"sort"
)

func sortSymbols(container *CurrencyContainer) []string {
	symbolArr := make([]string, len(container.Currencies))
	i := 0

	for symbol, _ := range container.Currencies {
		symbolArr[i] = symbol
		i++
	}

	sort.Strings(symbolArr)

	return symbolArr
}
