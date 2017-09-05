package modules

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

func Table(container *CurrencyContainer) {
	callClear()

	// Used to format dollar values in the table output
	accountant := accounting.Accounting{Symbol: "$", Precision: 2}

	tableData := [][]string{}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Currency", "Quantity", "Value"})

	for symbol, currency := range container.Currencies {
		quantStr := strconv.FormatFloat(currency.Quantity, 'f', 4, 64)
		cashStr := accountant.FormatMoney(currency.CashValue)

		arr := []string{symbol, quantStr, cashStr}
		tableData = append(tableData, arr)
	}

	for _, v := range tableData {
		table.Append(v)
	}

	table.SetFooter([]string{"", "", accountant.FormatMoney(container.TotalCashValue())})

	table.Render()

	t := time.Now()
	fmt.Println(t.Format("15:04:15"))

	fmt.Println()
}
