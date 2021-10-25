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
