/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
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

func makeFirstPageAppealButtons() [][]gotgbot.InlineKeyboardButton {
	if len(FirstPageButtonsRows) != 0 {
		return FirstPageButtonsRows
	}
	rows := make([][]gotgbot.InlineKeyboardButton, 2)
	rows[0] = append(rows[0], gotgbot.InlineKeyboardButton{
		Text:         "I will not do this again!",
		CallbackData: AutoAppealCbPrefix + firstAcceptCbData,
	})
	rows[1] = append(rows[1], gotgbot.InlineKeyboardButton{
		Text:         "Close this message",
		CallbackData: AutoAppealCbPrefix + CloseCbData,
	})
	FirstPageButtonsRows = rows
	return FirstPageButtonsRows
}

func makeDetailsPageAppealButtons(canAppeal bool) gotgbot.InlineKeyboardMarkup {
	var markup gotgbot.InlineKeyboardMarkup
	var rows [][]gotgbot.InlineKeyboardButton
	if canAppeal {
		rows = make([][]gotgbot.InlineKeyboardButton, 2)
		rows[0] = append(rows[0], gotgbot.InlineKeyboardButton{
			Text:         "I read and understand, unban me!",
			CallbackData: AutoAppealCbPrefix + detailsAcceptCbData,
		})
		rows[1] = append(rows[1], gotgbot.InlineKeyboardButton{
			Text: "Take me to support group",
			Url:  "https://t.me/PublicSafetyBureau",
		})
	} else {
		rows = make([][]gotgbot.InlineKeyboardButton, 1)
		rows[0] = append(rows[0], gotgbot.InlineKeyboardButton{
			Text: "Take me to support group",
			Url:  "https://t.me/PublicSafetyBureau",
		})
	}

	markup.InlineKeyboard = rows
	return markup
}

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	startCmd := handlers.NewCommand(StartCmd, startHandler)
	createCmd := handlers.NewCommand(CreateCmd, startHandler)
	newCmd := handlers.NewCommand(NewCmd, startHandler)
	autoAppealCb := handlers.NewCallback(appealCallBackQuery, appealCallBackResponse)
	startCmd.Triggers = t
	createCmd.Triggers = t
	newCmd.Triggers = t
	d.AddHandler(startCmd)
	d.AddHandler(createCmd)
	d.AddHandler(newCmd)
	d.AddHandler(autoAppealCb)
}
