package modules

import (
	"fmt"
	"os"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
)

func buildHistoryData(symbols []string, container *CurrencyContainer, accountant *accounting.Accounting) [][]string {
	tableData := [][]string{}

	return tableData
}

func History(container *CurrencyContainer) {
	callClear()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"IDX", "", "", ""})

	accountant := accounting.Accounting{Symbol: "$", Precision: 2}
	symbols := sortSymbols(container)
	tableData := buildHistoryData(symbols, container, &accountant)

	for _, v := range tableData {
		table.Append(v)
	}

	table.Render()

	fmt.Println(time.Now().Format("15:04:15"))

	fmt.Println()
}
