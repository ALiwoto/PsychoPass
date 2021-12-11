package reportPlugin

import (
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
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
		ParseMode:             sv.MarkDownV2,
		DisableWebPagePreview: true,
		ReplyMarkup:           getReportButtons(r.UniqueId),
	}

	for _, chat := range bases {
		sendReportMessage(chat, text, opts)
	}
}

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	sv.SendReportHandler = SendReportHandler
}
