package modules

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"time"

	// "github.com/Zauberstuhl/go-coinbase"
	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
)

var clear map[string]func() //create a map for storing clear funcs
func MakeClear() {
	clear = make(map[string]func())
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func callClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// func DisplayTable(accounts *coinbase.APIAccounts) {
func Table(cc *CryptoCurrencies) {
	callClear()

	// Used to format dollar values in the table output
	accountant := accounting.Accounting{Symbol: "$", Precision: 2}

	tableData := [][]string{}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Currency", "Value"})

	for currency, value := range cc.Currencies {
		valStr := accountant.FormatMoney(value)

		arr := []string{currency, valStr}
		tableData = append(tableData, arr)
	}

	for _, v := range tableData {
		table.Append(v)
	}

	table.SetFooter([]string{"", accountant.FormatMoney(cc.TotalValue)})

	table.Render()

	t := time.Now()
	fmt.Println(t.Format("15:04:15"))

	fmt.Println()
}