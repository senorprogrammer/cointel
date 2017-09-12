package modules

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
)

var clear map[string]func() //create a map for storing clear funcs
func MakeClearTerminal() {
	clear = make(map[string]func())
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func callClear() {
	value, ok := clear[runtime.GOOS] // runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          // if we defined a clear func for that platform:
		value() // we execute it
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

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
