package main

import (
	"fmt"
	"io/fs"
	"os"
	"runtime/debug"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylConfig"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/server"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/tgCore"
)

func main() {
	f := logging.LoadLogger(true)
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
	if database.IsFirstTime() {
		d, err := utils.CreateToken(sibylConfig.GetMasterId(), sibylValues.Owner)
		if err != nil {
			logging.Fatal(err)
		}

		logging.Info("Creating initial owner token...")
		logging.Info(d.Hash)
		os.WriteFile("owner.token", []byte(d.Hash), fs.ModePerm)
	}

	tgCore.StartTelegramBot()
	server.RunSibylSystem()
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
