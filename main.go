package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"gitlab.com/Dank-del/SibylAPI-Go/core/sibylConfig"
	"gitlab.com/Dank-del/SibylAPI-Go/core/utils/logging"
	"gitlab.com/Dank-del/SibylAPI-Go/server"
)

func main() {
	f := logging.LoadLogger()
	defer func() {
		if f != nil {
			f()
		}
	}()

	runApp()
}

func runApp() {
	defer recoverFromPanic()
	err := sibylConfig.LoadConfig()
	if err != nil {
		logging.Fatal(err)
	} else {
		panic("test")
	}

	err = server.RunSibylSystem()
	if err != nil {
		logging.Fatal(err)
	}
}

var totalPanics int

// recover from panic
// TODO: Start the sibyl system again with the
// appropriate configuration.
func recoverFromPanic() {
	if r := recover(); r != nil {
		details := debug.Stack()
		fmt.Println("Got panic:", r)
		fmt.Println(string(details))
		logging.LogPanic(details)
		max := sibylConfig.GetMaxPanics()
		if max != -1 && totalPanics >= int(max) {
			fmt.Println("Too many panics, exiting")
			os.Exit(0x1)
		} else {
			totalPanics++
			runApp()
		}
	}
}
