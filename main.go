package main

import (
	"fmt"
	"gitlab.com/Dank-del/SibylAPI-Go/core/utils"
	"gitlab.com/Dank-del/SibylAPI-Go/routes"
	"os"
	"runtime/debug"

	"gitlab.com/Dank-del/SibylAPI-Go/core/sibylConfig"
	"gitlab.com/Dank-del/SibylAPI-Go/core/utils/logging"
	"gitlab.com/Dank-del/SibylAPI-Go/database"
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
	}

	database.StartDatabase()
	if database.SESSION.Find(&database.Token{}).RowsAffected == 0 {
		d, err := utils.CreateToken(sibylConfig.SibylConfig.MasterId, server.AdminParam)
		if err != nil {
			logging.Fatal(err)
		}
		logging.Info("Creating Initial ADMIN token")
		logging.Info(d.Hash)
		logging.Info("Write it down, cause it won't appear again!")
	}
	serv := server.RunSibylSystem()
	if err != nil {
		logging.Fatal(err)
	}
	server.ServerEngine.GET("createToken", routes.CreateToken)
	err = serv.Run()
	if err != nil {
		logging.Error(err)
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
