package main

import (
	"fmt"
	"os"
	"os/signal"
	// "strings"
	"syscall"
	"time"

	"github.com/Zauberstuhl/go-coinbase"
)

const OneMinute = 60 * time.Second
const FiveMinutes = 300 * time.Second

func cleanup() {
	fmt.Println("terminating")
}

func main() {
	// Creates a channel to watch for Ctl-C to terminate the program
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	// Coinbase client
	client := coinbase.APIClient{
		Key:    os.Getenv("COINBASE_KEY"),
		Secret: os.Getenv("COINBASE_SECRET"),
	}

	// Loop the app, periodically calling Coinbase for account balances
	for {
		accounts, err := client.Accounts()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%v", accounts)
		fmt.Println("Checking...")

		for _, acc := range accounts.Data {
			fmt.Printf("ID: %s\nName: %s\nType: %s\nAmount: %f\nCurrency: %s\n",
				acc.Id, acc.Name, acc.Type,
				acc.Balance.Amount, acc.Balance.Currency)
		}

		time.Sleep(OneMinute)
	}
}
