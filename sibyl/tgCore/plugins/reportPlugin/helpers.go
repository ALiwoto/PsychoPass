package reportPlugin

import (
	sv "github.com/AnimeKaizoku/sibylapi-go/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/utils/logging"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func sendReportMessage(chat int64, text string, opts *gotgbot.SendMessageOpts) {
	_, err := sv.HelperBot.SendMessage(chat, text, opts)
	if err != nil {
		logging.Debug("tried to send message to ", chat, err)
	}
}
