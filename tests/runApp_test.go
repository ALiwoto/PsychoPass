/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package tests_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/server"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/tgCore"
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
		// run the app in another goroutine
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
	err = sibylConfig.LoadTriggers()
	if err != nil {
		logging.Fatal(err)
	}
	database.StartDatabase()
	prepareOwnerToken()
	tgCore.StartTelegramBot()
	server.RunSibylSystem()
}

// getOwnerToken returns the owner's token from owner.token file
func getOwnerToken() string {
	t := utils.ReadOneFile("owner.token", "owners.token")
	if strings.Contains(t, "\n") {
		strs := strings.Split(t, "\n")
		return strs[0]
	}
	return t
}

func prepareOwnerToken() {
	if database.IsFirstTime() {
		err := ioutil.WriteFile(sibylValues.OwnersTokenFileName, []byte(""), 0644)
		if err != nil {
			logging.Fatal(err)
		}

		file, err := os.OpenFile(sibylValues.OwnersTokenFileName, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			logging.Fatal(err)
		}

		for _, id := range sibylConfig.GetOwnersID() {
			generateOwnerToken(id, file)
		}
	}
}

func generateOwnerToken(id int64, file *os.File) {
	d, err := utils.CreateToken(id, sibylValues.Owner)
	if err != nil {
		logging.Fatal(err)
	}

	file.WriteString(d.Hash + "\n")
}
