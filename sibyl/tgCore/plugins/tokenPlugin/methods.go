package tokenPlugin

import (
	"strconv"
	"strings"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func (a *AssignValue) ParseToMd(info mdparser.WMarkDown) mdparser.WMarkDown {
	by := strconv.FormatInt(a.agent.UserId, 10)
	md := mdparser.GetNormal(SpecialChar + "#Assignment request\n")
	if a.agentProfile != nil {
		name := utils.GetNameFromUser(a.agentProfile, a.agent.GetStringPermission())
		name = strings.ReplaceAll(name, SpecialChar, "")
		md.Bold(SpecialChar+" • By: ").Mention(name+SpecialChar, a.agentProfile.Id)
	} else {
		md.Bold(SpecialChar + " • By: ").Mono(by + SpecialChar)
	}
	md.ElThis().AppendThis(info)
	md.Bold(SpecialChar+" • Source: ").Link("here", a.src)
	return md
}

func (a *AssignValue) getAcceptCbData() string {
	return strconv.FormatInt(a.targetChat.Id, 10) + CbSep + a.permValue.ToString()
}

func (a *AssignValue) getAssignmentButton() *gotgbot.InlineKeyboardMarkup {
	kb := &gotgbot.InlineKeyboardMarkup{}
	kb.InlineKeyboard = make([][]gotgbot.InlineKeyboardButton, 2)

	kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
		Text:         "✅ Accept",
		CallbackData: AssignCbData + CbSep + a.getAcceptCbData(),
	})
	kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
		Text:         "❌ Reject",
		CallbackData: AssignCbData + CbSep + RejectCbData,
	})

	kb.InlineKeyboard[1] = append(kb.InlineKeyboard[1], gotgbot.InlineKeyboardButton{
		Text:         "Close",
		CallbackData: AssignCbData + CbSep + CloseCbData,
	})

	return kb
}
