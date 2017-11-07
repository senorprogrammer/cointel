package modules

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
)

func buildTableData(container *CurrencyContainer, accountant *accounting.Accounting) [][]string {
	tableData := [][]string{}

	symbols := sortSymbols(container)

	for _, symbol := range symbols {
		hist := container.Histories[symbol]

		quantStr := strconv.FormatFloat(hist.Last().Quantity, 'f', 4, 64)
		cashStr := accountant.FormatMoney(hist.Last().CashValue)

		arr := []string{symbol, quantStr, cashStr}
		tableData = append(tableData, arr)
	}

	return tableData
}

func Table(client *CoinbaseClient, persistFlag bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Currency", "Quantity", "Value"})

	accountant := accounting.Accounting{Symbol: "$", Precision: 2}

	for {
		if persistFlag {
			CallClear()
		}

		client.Refresh()

		tableData := buildTableData(&client.Container, &accountant)
		for _, v := range tableData {
			table.Append(v)
		}

		table.SetFooter([]string{"", "", accountant.FormatMoney(client.Container.TotalCashValue())})

		table.Render()

		fmt.Println(time.Now().Format("15:04:15\n"))

		// fmt.Println()

		// GET RID OF THIS
		time.Sleep(300 * time.Second)
	}
}
