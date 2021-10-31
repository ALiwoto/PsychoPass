package tgCore

import (
	"fmt"
	"net/http"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylConfig"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
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
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		logging.Info("Unable to login to the helper bot due to:", err)
		return
	}

	url := sibylConfig.GetAPIUrl()
	if len(url) != 0 {
		b.APIURL = url
	}

	utmp := ext.NewUpdater(nil)
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
