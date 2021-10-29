package tests_test

import (
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/sibylConfig"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/utils"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/utils/logging"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/database"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/server"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/tgCore"
)

//const baseUrl = "http://localhost:8080/"
const (
	userId01 = "1341091260"
	userId02 = userId01
	userId03 = "895373440"
	userId04 = "792109647"
	userId05 = "701937965"
)

var (
	user01TokenTmp = ""
	User02TokenTmp = ""
	user03TokenTmp = ""
	user04TokenTmp = ""
	user05TokenTmp = ""
)

var (
	baseUrl = ""
)

func getBaseUrl() string {
	if len(baseUrl) == 0 {
		b, _ := ioutil.ReadFile("baseUrl.ini")
		baseUrl = string(b)
	}
	return baseUrl
}

func decideToRun() {
	b := getBaseUrl()
	if strings.Contains(b, "localhost") ||
		strings.Contains(b, "0.0.0.0") {
		// run the app in anoter goroutine
		go runApp()

		time.Sleep(time.Millisecond * 600)
	}
}

func closeServer() {
	if server.ServerEngine == nil {
		return
	}
	srv := &http.Server{
		Addr:    sibylConfig.GetPort(),
		Handler: server.ServerEngine,
	}

	srv.Close()
}

func runApp() {
	//defer recoverFromPanic()
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

// getOwnerToken returns the owner's token from owner.token file
func getOwnerToken() string {
	return string(utils.ReadFile("owner.token"))
}
