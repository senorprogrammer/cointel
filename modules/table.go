package modules

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
)

func buildTableData(symbols []string, container *CurrencyContainer, accountant *accounting.Accounting) [][]string {
	tableData := [][]string{}

	for _, symbol := range symbols {
		hist := container.Histories[symbol]

		quantStr := strconv.FormatFloat(hist.Last().Quantity, 'f', 4, 64)
		cashStr := accountant.FormatMoney(hist.Last().CashValue)

		arr := []string{symbol, quantStr, cashStr}
		tableData = append(tableData, arr)
	}

	return tableData
}

func Table(container *CurrencyContainer) {
	callClear()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Currency", "Quantity", "Value"})

	accountant := accounting.Accounting{Symbol: "$", Precision: 2}
	symbols := sortSymbols(container)
	tableData := buildTableData(symbols, container, &accountant)

	for _, v := range tableData {
		table.Append(v)
	}

	table.SetFooter([]string{"", "", accountant.FormatMoney(container.TotalCashValue())})

	table.Render()

	fmt.Println(time.Now().Format("15:04:15"))

	fmt.Println()
}
