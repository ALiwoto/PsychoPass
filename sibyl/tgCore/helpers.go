package tgCore

import (
	"fmt"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylConfig"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/logging"
)

func StartTelegramBot() {
	token := sibylConfig.GetBotToken()
	if len(token) == 0 {
		logging.Warn("token of the helper bot is not set")
		logging.Warn("Sibyl System won't be able to interact with telegram")
		return
	}

	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})

	if err != nil {
		logging.Warn("Unable to login to the helper bot due to:", err)
		return
	}

	utmp := ext.NewUpdater(nil)
	updater := &utmp
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: sibylConfig.DropUpdates(),
	})
	if err != nil {
		logging.Warn("failed to start polling for new updates due to:", err)
		return
	}

	logging.Info(fmt.Sprintf("%s has started | ID: %d", b.Username, b.Id))

	sibylValues.HelperBot = b
	sibylValues.BotUpdater = updater
	LoadAllHandlers(updater.Dispatcher, sibylConfig.GetCmdPrefixes())
}
