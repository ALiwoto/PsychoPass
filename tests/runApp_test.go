package tests_test

import (
	"io/fs"
	"os"

	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylConfig"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/logging"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/database"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/server"
)

const baseUrl = "http://localhost:8080/"
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

func runApp() {
	//defer recoverFromPanic()
	err := sibylConfig.LoadConfig()
	if err != nil {
		logging.Fatal(err)
	}

	database.StartDatabase()
	if database.IsFirstTime() {
		d, err := utils.CreateToken(sibylConfig.SibylConfig.MasterId,
			sibylValues.Owner)
		if err != nil {
			logging.Fatal(err)
		}
		logging.Info("Creating Initial ADMIN token")
		logging.Info(d.Hash)
		os.WriteFile("owner.token", []byte(d.Hash), fs.ModePerm)
	}

	server.RunSibylSystem()
}

// getOwnerToken returns the owner's token from owner.token file
func getOwnerToken() string {
	return string(utils.ReadFile("owner.token"))
}
