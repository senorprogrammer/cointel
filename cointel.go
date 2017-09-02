package main

import (
	"github.com/senorprogrammer/cointel/modules"

	"time"
)

const OneMinute = 60 * time.Second
const FiveMinutes = 300 * time.Second
const FifteenMinutes = (15 * 60) * time.Second

func main() {
	modules.MakeTermination()
	modules.MakeClearTerminal()

	// Each exchange needs it's own client that handles how to load up their own
	// CurrencyContainer with applicable data
	coinbaseClient := modules.NewCoinbaseClient()

	for {
		coinbaseClient.Refresh()

		// Output the results
		modules.Table(&coinbaseClient.Container)
		// fmt.Println(modules.Json(&coinbaseClient.Container))

		time.Sleep(OneMinute)
	}
}
