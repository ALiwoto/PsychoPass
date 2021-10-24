package main

import (
	"fmt"
	"runtime/debug"

	"gitlab.com/Dank-del/SibylAPI-Go/core/sibylConfig"
	"gitlab.com/Dank-del/SibylAPI-Go/core/utils/logging"
	"gitlab.com/Dank-del/SibylAPI-Go/server"
)

func main() {
	defer recoverFromPanic()

	f := logging.LoadLogger()
	defer func() {
		if f != nil {
			f()
		}
	}()

	err := sibylConfig.LoadConfig()
	if err != nil {
		logging.Fatal(err)
	}

	err = server.RunSibylSystem()
	if err != nil {
		logging.Fatal(err)
	}
}

// recover from panic
// TODO: Start the sibyl system again with the
// appropriate configuration.
func recoverFromPanic() {
	if r := recover(); r != nil {
		details := debug.Stack()
		fmt.Println("Got panic: ", r)
		fmt.Println(string(details))
		logging.LogPanic(details)
	}
}
