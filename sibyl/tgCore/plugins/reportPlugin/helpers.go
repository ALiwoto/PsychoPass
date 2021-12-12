package reportPlugin

import (
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func getReportButtons(uniqueId string) *gotgbot.InlineKeyboardMarkup {

	kb := &gotgbot.InlineKeyboardMarkup{}

	kb.InlineKeyboard = make([][]gotgbot.InlineKeyboardButton, 2)

	kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
		Text:         "✅ Approve",
		CallbackData: ReportPrefix + sepChar + AcceptData + sepChar + uniqueId,
	})
	kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
		Text:         "❌ Reject",
		CallbackData: ReportPrefix + sepChar + RejectData + sepChar + uniqueId,
	})

	kb.InlineKeyboard[1] = append(kb.InlineKeyboard[1], gotgbot.InlineKeyboardButton{
		Text:         "Close",
		CallbackData: "ap_close", // let start package handle close button xD
	})

	return kb
}

func sendReportMessage(chat int64, text string, opts *gotgbot.SendMessageOpts) {
	_, err := sv.HelperBot.SendMessage(chat, text, opts)
	if err != nil {
		logging.Debug("Tried to send message to ", chat, err)
	}
}

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	sv.SendReportHandler = SendReportHandler
}
