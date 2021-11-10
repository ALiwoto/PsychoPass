package startPlugin

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func makeNormalButtons() [][]gotgbot.InlineKeyboardButton {
	if len(normalButtonsRows) != 0 {
		return normalButtonsRows
	}
	rows := make([][]gotgbot.InlineKeyboardButton, 3)
	rows[0] = append(rows[0], gotgbot.InlineKeyboardButton{
		Text: "What is PsychoPass?",
		Url:  "https://t.me/PsychoPass",
	})

	rows[1] = append(rows[1], gotgbot.InlineKeyboardButton{
		Text: "Support group",
		Url:  "https://t.me/PublicSafetyBureau",
	})
	rows[1] = append(rows[1], gotgbot.InlineKeyboardButton{
		Text: "Report Spam",
		Url:  "https://t.me/PublicSafetyBureau",
	})

	rows[2] = append(rows[2], gotgbot.InlineKeyboardButton{
		Text:         "Get API token",
		CallbackData: "get_token",
	})

	normalButtonsRows = rows
	return normalButtonsRows
}

func makeSingleAppealButtons() [][]gotgbot.InlineKeyboardButton {
	if len(singleButtonsRows) != 0 {
		return singleButtonsRows
	}
	rows := make([][]gotgbot.InlineKeyboardButton, 1)
	rows[0] = append(rows[0], gotgbot.InlineKeyboardButton{
		Text: "Appeal ban",
		Url:  "https://t.me/PublicSafetyBureau",
	})
	singleButtonsRows = rows
	return singleButtonsRows
}

func makeFirstPageAppealButtons() [][]gotgbot.InlineKeyboardButton {
	if len(FirstPageButtonsRows) != 0 {
		return FirstPageButtonsRows
	}
	rows := make([][]gotgbot.InlineKeyboardButton, 2)
	rows[0] = append(rows[0], gotgbot.InlineKeyboardButton{
		Text:         "I will not do this again!",
		CallbackData: AutoAppealCbPrefix + AcceptCbData,
	})
	rows[1] = append(rows[1], gotgbot.InlineKeyboardButton{
		Text:         "Close this message",
		CallbackData: AutoAppealCbPrefix + CloseCbData,
	})
	FirstPageButtonsRows = rows
	return FirstPageButtonsRows
}

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	startCmd := handlers.NewCommand(StartCmd, startHandler)
	createCmd := handlers.NewCommand(CreateCmd, startHandler)
	newCmd := handlers.NewCommand(NewCmd, startHandler)
	startCmd.Triggers = t
	createCmd.Triggers = t
	newCmd.Triggers = t
	d.AddHandler(startCmd)
	d.AddHandler(createCmd)
	d.AddHandler(newCmd)
}
