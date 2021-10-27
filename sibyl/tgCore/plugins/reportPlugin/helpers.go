package reportPlugin

import (
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/logging"
)

func sendReportMessage(chat int64, text string, opts *gotgbot.SendMessageOpts) {
	_, err := sv.HelperBot.SendMessage(chat, text, opts)
	if err != nil {
		logging.Debug("tried to send message to ", chat, err)
	}
}

// getReportButtons will give you buttons with this data:
// "report prefix" + "_" + "close/accept/delete" + "_" + "unique id".
func getReportButtons(rId int64) gotgbot.ReplyMarkup {
	uniqueId := strconv.FormatInt(rId, reportIdBase)
	kb := &gotgbot.InlineKeyboardMarkup{}
	kb.InlineKeyboard = append(kb.InlineKeyboard, []gotgbot.InlineKeyboardButton{})
	kb.InlineKeyboard = append(kb.InlineKeyboard, []gotgbot.InlineKeyboardButton{})
	kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
		Text: "Accept",
		CallbackData: reportCbPrefix + cbDataSeparator +
			acceptCbValue + cbDataSeparator + uniqueId,
	})
	kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
		Text: "Close",
		CallbackData: reportCbPrefix + cbDataSeparator +
			closeCbValue + cbDataSeparator + uniqueId,
	})
	kb.InlineKeyboard[1] = append(kb.InlineKeyboard[1], gotgbot.InlineKeyboardButton{
		Text: "Delete message",
		CallbackData: reportCbPrefix + cbDataSeparator +
			deleteCbValue + cbDataSeparator + uniqueId,
	})
	return kb
}

func editAcceptedMessage(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	u := ctx.EffectiveUser
	msg.EditText(b, msg.Text+"\n"+"Accepted by "+u.FirstName+".",
		&gotgbot.EditMessageTextOpts{
			Entities:              msg.Entities,
			DisableWebPagePreview: true,
		})
	return nil
}

func editClosedMessage(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	u := ctx.EffectiveUser
	msg.EditText(b, msg.Text+"\n"+"Closed by "+u.FirstName+".",
		&gotgbot.EditMessageTextOpts{
			Entities:              msg.Entities,
			DisableWebPagePreview: true,
		})
	return nil
}
