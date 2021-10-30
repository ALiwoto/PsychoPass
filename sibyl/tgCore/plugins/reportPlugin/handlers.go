package reportPlugin

import (
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylConfig"
	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func SendReportHandler(r *sv.Report) {
	// prevent from panic xD
	if sv.HelperBot == nil {
		return
	}

	bases := sibylConfig.GetBaseChatIds()
	if len(bases) == 0 {
		// there is no chat to send the report to...
		// ignore the report...
		return
	}

	var text string
	var opts *gotgbot.SendMessageOpts

	md := r.ParseAsMd()

	text = md.ToString()
	opts = &gotgbot.SendMessageOpts{
		ParseMode: sv.MarkDownV2,
		//ReplyMarkup: getReportButtons(r.GetUniqueId()),
	}

	for _, chat := range bases {
		sendReportMessage(chat, text, opts)
	}
}
