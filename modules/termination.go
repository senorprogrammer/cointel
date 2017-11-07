package modules

import (
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
}

/*
* If the persistFlag is false (ie don't keep executing), then clean up the app and exit
 */
func Terminate(persistFlag bool) {
	if !persistFlag {
		Cleanup()
		os.Exit(0)
	}
}
