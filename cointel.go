package main

import (
	"github.com/senorprogrammer/cointel/modules"

	"flag"
	_ "fmt"
	_ "os"
	"time"
)

const OneMinute = 60 * time.Second
const FiveMinutes = 300 * time.Second
const FifteenMinutes = (15 * 60) * time.Second

func main() {
	formatFlag := flag.String("format", "json", "Either 'history', 'table' or 'json'")
	persistFlag := flag.Bool("persist", false, "Either true or false")
	flag.Parse()

	modules.MakeTermination()
	modules.MakeClearTerminal()

	coinbaseClient := modules.NewCoinbaseClient()

	switch *formatFlag {
	case "json":
		modules.Json(&coinbaseClient, *persistFlag)
	case "history":
		modules.History(&coinbaseClient, *persistFlag)
	case "table":
		modules.Table(&coinbaseClient, *persistFlag)
	}
}
