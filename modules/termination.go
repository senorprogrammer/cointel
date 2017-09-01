package modules

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Creates a channel to watch for Ctl-C to terminate the program
func MakeTermination() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
}

func cleanup() {
	fmt.Println("goodbye")
}
