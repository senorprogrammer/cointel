package modules

import (
	"sort"
)

func sortSymbols(container *CurrencyContainer) []string {
	symbolArr := make([]string, len(container.Histories))
	i := 0

	for symbol, _ := range container.Histories {
		symbolArr[i] = symbol
		i++
	}

	sort.Strings(symbolArr)

	return symbolArr
}
