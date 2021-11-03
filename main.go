package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
	"time"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylConfig"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/server"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/tgCore"
)

func main() {
	sibylValues.ServerStartTime = time.Now()
	f := logging.LoadLogger(true)
	if f != nil {
		defer f()
	}

	runApp()
}

func runApp() {
	defer recoverFromPanic()
	err := sibylConfig.LoadConfig()
	if err != nil {
		logging.Fatal(err)
	}
	err = sibylConfig.LoadTriggers()
	if err != nil {
		logging.Fatal(err)
	}
	database.StartDatabase()
	prepareOwnerTokens()
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
		detailsStr := fmt.Sprint(r) + "\n" + string(details)
		logging.LogPanic([]byte(detailsStr))
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

func prepareOwnerTokens() {
	if database.IsFirstTime() {
		ids := sibylConfig.GetOwnersID()
		if len(ids) == 0 {
			logging.Fatal("There should be at least one owner")
		}

		err := ioutil.WriteFile(sibylValues.OwnersTokenFileName, []byte(""), 0644)
		if err != nil {
			logging.Fatal(err)
		}

		file, err := os.OpenFile(sibylValues.OwnersTokenFileName,
			os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			logging.Fatal(err)
		}

		defer file.Close()

		for _, id := range ids {
			generateOwnerToken(id, file)
		}
	}
}

func generateOwnerToken(id int64, file *os.File) {
	d, err := utils.CreateToken(id, sibylValues.Owner)
	if err != nil {
		logging.Fatal(err)
	}

	_, err = file.WriteString(d.Hash + "\n")
	if err != nil {
		logging.Fatal(err)
	}
}
