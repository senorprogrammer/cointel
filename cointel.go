package main

import (
	"github.com/senorprogrammer/cointel/modules"

	"flag"
	"fmt"
	"os"
	"time"
)

const OneMinute = 60 * time.Second
const FiveMinutes = 300 * time.Second
const FifteenMinutes = (15 * 60) * time.Second

func main() {
	formatFlag := flag.String("format", "json", "Either 'table' or 'json'")
	persistFlag := flag.Bool("persist", false, "Either true or false")
	flag.Parse()

	modules.MakeTermination()
	modules.MakeClearTerminal()

	coinbaseClient := modules.NewCoinbaseClient()

	for {
		if *persistFlag {
			modules.CallClear()
		}

		coinbaseClient.Refresh()

		if *formatFlag == "json" {
			fmt.Println(modules.Json(&coinbaseClient.Container))
		} else if *formatFlag == "history" {
			modules.History(&coinbaseClient.Container)
		} else if *formatFlag == "table" {
			modules.Table(&coinbaseClient.Container)
		} else {
			fmt.Println(modules.Json(&coinbaseClient.Container))
		}

		// If this process should not persist, then kill itself
		// Otherwise it'll periodically check until Ctl-C is pressed
		if !*persistFlag {
			modules.Cleanup()
			os.Exit(0)
		}

		time.Sleep(FifteenMinutes)
	}
}
