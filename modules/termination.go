package modules

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Creates a channel to watch for Ctl-C to terminate the program
func MakeTermination() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Cleanup()
		os.Exit(0)
	}()
}

func Cleanup() {
	fmt.Println("goodbye")
}
