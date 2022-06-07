/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package tgCore

import (
	"fmt"
	"net/http"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func StartTelegramBot() {
	token := sibylConfig.GetBotToken()
	if len(token) == 0 {
		logging.Info("Helper bot token is not set")
		return
	}

	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: 6 * gotgbot.DefaultTimeout,
			APIURL:  sibylConfig.GetAPIUrl(),
		},
	})
	if err != nil {
		logging.Info("Unable to login to the helper bot due to:", err)
		return
	}

	mdparser.AddSecret(b.GetToken(), "$TOKEN")

	uOptions := &ext.UpdaterOpts{
		DispatcherOpts: ext.DispatcherOpts{
			MaxRoutines: -1,
		},
	}
	utmp := ext.NewUpdater(uOptions)
	updater := &utmp
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: sibylConfig.DropUpdates(),
	})
	if err != nil {
		logging.Warn("Failed to start polling for new updates due to:", err)
		return
	}

	logging.Info(fmt.Sprintf("%s has started | ID: %d", b.Username, b.Id))

	sibylValues.HelperBot = b
	sibylValues.BotUpdater = updater
	LoadAllHandlers(updater.Dispatcher, sibylConfig.GetCmdPrefixes())
}
