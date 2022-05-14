/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package tokenPlugin

import (
	"strconv"
	"strings"
	"time"

	"github.com/ALiwoto/mdparser/mdparser"
	ws "github.com/AnimeKaizoku/ssg/ssg"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	loadButtons(d)
	revokeCmd := handlers.NewCommand(RevokeCmd, revokeHandler)
	assignCmd := handlers.NewCommand(AssignCmd, assignHandler)
	getTokenCb := handlers.NewCallback(getTokenCallBackQuery, getTokenCallBackResponse)
	assignCb := handlers.NewCallback(assignCallBackQuery, assignCallBackResponse)
	revokeTokenCb := handlers.NewCallback(revokeTokenCallBackQuery, revokeTokenCallBackResponse)
	revokeCmd.Triggers = t
	assignCmd.Triggers = t
	assignCb.AllowChannel = true
	d.AddHandler(revokeCmd)
	d.AddHandler(assignCmd)
	d.AddHandler(getTokenCb)
	d.AddHandler(assignCb)
	d.AddHandler(revokeTokenCb)
}

func loadButtons(d *ext.Dispatcher) {
	if startCymaticScanButton == nil {
		kb := &gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: make([][]gotgbot.InlineKeyboardButton, 1),
		}

		kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
			Text: "Start Cymatic Scan",
			Url:  "https://t.me/" + sv.HelperBot.Username + "?start",
		})

		startCymaticScanButton = kb
	}
}

func showUserIsBanned(b *gotgbot.Bot, ctx *ext.Context, targetUser *sv.User, p string, replied bool) {
	var err error
	var md, uMd, suMd mdparser.WMarkDown
	msg := ctx.EffectiveMessage
	var strName string
	strNameId := strconv.FormatInt(targetUser.UserID, 10) // reserved value
	if replied {
		strName = msg.ReplyToMessage.From.FirstName
		suMd = mdparser.GetBold("• User: ")
		suMd.Mention(strName, targetUser.UserID).ElThis()
		suMd.Bold("• ID: ").Mono(strNameId).ElThis()
	} else {
		ch, err := b.GetChat(targetUser.UserID)
		if err != nil {
			return
		}
		strName = strings.TrimSpace(ch.FirstName)
		if len(strName) == 0 {
			strName = strings.TrimSpace(ch.LastName)
		}
		if len(strName) == 0 {
			strName = strNameId
		}
		suMd = mdparser.GetBold("• User: ")
		suMd.Mention(strName, targetUser.UserID).ElThis()
	}
	uMd = mdparser.GetUserMention(strName, targetUser.UserID)
	md = mdparser.GetNormal("Scanning ").AppendThis(uMd).Normal("...")
	msg, err = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode:                sv.MarkDownV2,
		AllowSendingWithoutReply: true,
		DisableWebPagePreview:    true,
	})
	if err != nil {
		logging.UnexpectedError(err)
		return
	}

	time.Sleep(2 * time.Second)

	md = md.Bold("• Is banned: ")
	md.Mono(ws.YesOrNo(targetUser.Banned)).ElThis()
	md.Bold("• Crime Coefficient: ")
	md.Mono(strconv.Itoa(targetUser.CrimeCoefficient)).ElThis()

	msg, _, err = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode: sv.MarkDownV2,
	})
	if err != nil {
		logging.UnexpectedError(err)
		return
	}

	time.Sleep(2 * time.Second)

	md.ElThis().Bold("Verdict: ").AppendThis(uMd)
	md.Normal(" cannot be assigned as " + p + " because their crime coefficient is ")
	se, cc := targetUser.EstimateCrimeCoefficientSep()
	md.Normal(se).Mono(cc).ElThis()
	md.Bold("Attached reason: ").Mono(targetUser.Reason)
	_, _, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode: sv.MarkDownV2,
	})
}

func showUserAssigned(b *gotgbot.Bot, ctx *ext.Context, aValue *AssignValue) {
	var err error
	var md, uMd mdparser.WMarkDown
	namae := aValue.targetChat.FirstName
	uMd = mdparser.GetUserMention(namae+SpecialChar, aValue.targetChat.Id)
	strId := strconv.FormatInt(aValue.targetChat.Id, 10)
	md = mdparser.GetBold(SpecialChar + " • User: ").AppendThis(uMd).ElThis()
	md.Bold(SpecialChar + " • ID: ").Mono(strId).ElThis()
	md.Bold(SpecialChar + " • Is banned: ").Mono("No").ElThis()
	md.Bold(SpecialChar + " • Crime Coefficient: ")
	md.Mono(aValue.target.EstimateCrimeCoefficient())
	md.ElThis()
	// let the goroutine sleep for 2 seconds
	time.Sleep(2 * time.Second)
	aValue.msg, _, err = aValue.msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:             sv.MarkDownV2,
		DisableWebPagePreview: true,
	})
	if err != nil {
		logging.UnexpectedError(err)
		return
	}

	time.Sleep(3 * time.Second)
	mdBack := md.El()

	if aValue.agent.CanTryChangePermission(true) {
		md = mdparser.GetBold("Assigned Successfully! ").ElThis().AppendThis(md)
		md.Normal("\n✳️ ").AppendThis(uMd).Normal(" has now been assigned as ")
		md.Bold(aValue.perm)
		md.Normal("!\nTheir dominator and token have been sent to their ")
		md.Link("PM", "http://t.me/"+b.Username).Normal(".")
		_, _, _ = aValue.msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
			ParseMode:             sv.MarkDownV2,
			DisableWebPagePreview: true,
		})
	} else {
		md = mdparser.GetBold("Assignment request has been sent to Sibyl System! \n")
		md.AppendThis(mdBack)
		md.Normal("✳️ ").AppendThis(uMd).Normal(" will be assigned as ")
		md.Bold(aValue.perm).Normal(" after verification.")
		_, _, _ = aValue.msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
			ParseMode:             sv.MarkDownV2,
			DisableWebPagePreview: true,
		})

		bases := sibylConfig.GetBaseChatIds()
		if len(bases) == 0 {
			// there is no chat to send the assignment request to...
			// ignore the request...
			return
		}

		text := aValue.ParseToMd(mdBack).ToString()
		opts := &gotgbot.SendMessageOpts{
			ParseMode:             sv.MarkDownV2,
			ReplyMarkup:           aValue.getAssignmentButton(),
			DisableWebPagePreview: true,
		}

		utils.SendMultipleMessages(bases, text, opts)
	}
}

func toAssignValue(msg *gotgbot.Message, perm sv.UserPermission) *AssignValue {
	text := msg.Text
	myStrs := ws.Split(text, SpecialChar)
	a := &AssignValue{
		permValue: perm, // for now, since it's impossible for another values
		src:       utils.GetLinkFromMessage(msg),
	}

	/*
		How to parse?
		1- first text_mention is always the person who tried to assign.
		2 - for finding the target, we should look for "• ID: ".
	*/

	for _, current := range msg.Entities {
		if current.Type == "text_mention" {
			a.agentId = current.User.Id
			break
		}
	}

	for _, current := range myStrs {
		if strings.Contains(current, "• ID: ") {
			idStr := strings.TrimPrefix(current, " • ID: ")
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err == nil {
				a.targetId = id
			}
			break
		}
	}

	return a
}
