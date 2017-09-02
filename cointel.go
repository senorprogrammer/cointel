package main

import (
	"github.com/senorprogrammer/cointel/modules"

	"flag"
	"fmt"
	"time"
)

const OneMinute = 60 * time.Second
const FiveMinutes = 300 * time.Second
const FifteenMinutes = (15 * 60) * time.Second

func main() {
	outputPtr := flag.String("output", "table", "Either 'table' or 'json'")
	flag.Parse()

	modules.MakeTermination()
	modules.MakeClearTerminal()

	coinbaseClient := modules.NewCoinbaseClient()

	for {
		coinbaseClient.Refresh()

		// Output the results
		if *outputPtr == "json" {
			fmt.Println(modules.Json(&coinbaseClient.Container))
		} else {
			modules.Table(&coinbaseClient.Container)
		}

		time.Sleep(OneMinute)
	}
}
