package modules

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
)

func buildHistoryData(container *CurrencyContainer, accountant *accounting.Accounting) [][]string {
	// historyData is an array of arrays of strings (which are converted currency values, aka floats)
	historyData := [][]string{}

	symbols := sortSymbols(container)

	for _, symbol := range symbols {
		arr := []string{}

		history := container.Histories[symbol]

		for _, curr := range history.CryptoCurrencies {
			arr = append(arr, strconv.FormatFloat(curr.CashValue, 'f', 2, 64))
		}

		// Don't display more than n entries in the list, else it wraps awkwardly
		if len(arr) > 10 {
			arr = arr[len(arr)-10:]
		}
		arr = append([]string{symbol}, arr...)

		historyData = append(historyData, arr)
	}

	return historyData
}

func History(container *CurrencyContainer) {
	table := tablewriter.NewWriter(os.Stdout)

	accountant := accounting.Accounting{Symbol: "$", Precision: 2}
	historyData := buildHistoryData(container, &accountant)

	for _, v := range historyData {
		table.Append(v)
	}

	// The header is a bunch of empty cells, equal to the number of history entries (minus the first label)
	headers := append([]string{"Currency"}, make([]string, len(historyData[0])-1)...)
	table.SetHeader(headers)

	table.Render()

	fmt.Println(time.Now().Format("15:04:15"))

	fmt.Println()
}
