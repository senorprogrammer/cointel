package main

import (
	"github.com/senorprogrammer/cointel/modules"

	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	// "github.com/Zauberstuhl/go-coinbase"
)

const OneMinute = 60 * time.Second
const FiveMinutes = 300 * time.Second
const FifteenMinutes = (15 * 60) * time.Second

// Creates a channel to watch for Ctl-C to terminate the program
func makeTermination() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
}

func cleanup() {}

func main() {
	makeTermination()
	modules.MakeClear()

	// A struct to aggregate the financial values of each coin type
	cryptoCurrencies := modules.NewCryptoCurrencies()

	// Coinbase client
	cbClient := modules.CoinbaseClient()

	// Loop the app, periodically calling Coinbase for account balances
	for {
		cbAccounts, err := cbClient.Accounts()
		if err != nil {
			return
		}

		cryptoCurrencies.CoinbaseUpdate(&cbAccounts)

		// Output the results
		modules.Table(&cryptoCurrencies)
		fmt.Println(modules.Json(&cryptoCurrencies))

		fmt.Println("\n")

		// And wait awhile until we hit the API for another update
		time.Sleep(FifteenMinutes)
	}
}
