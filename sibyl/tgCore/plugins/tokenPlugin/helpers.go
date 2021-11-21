package tokenPlugin

import (
	"strconv"
	"strings"
	"time"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylConfig"
	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	loadButtons(d)
	revokeCmd := handlers.NewCommand(RevokeCmd, revokeHandler)
	assignCmd := handlers.NewCommand(AssignCmd, assignHandler)
	getTokenCb := handlers.NewCallback(getTokenCallBackQuery, getTokenCallBackResponse)
	revokeTokenCb := handlers.NewCallback(revokeTokenCallBackQuery, revokeTokenCallBackResponse)
	revokeCmd.Triggers = t
	assignCmd.Triggers = t
	d.AddHandler(revokeCmd)
	d.AddHandler(assignCmd)
	d.AddHandler(getTokenCb)
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
		suMd.AppendMentionThis(strName, targetUser.UserID).ElThis()
		suMd.AppendBoldThis("• ID: ").AppendMonoThis(strNameId).ElThis()
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
		suMd.AppendMentionThis(strName, targetUser.UserID).ElThis()
	}
	uMd = mdparser.GetUserMention(strName, targetUser.UserID)
	md = mdparser.GetNormal("Scanning ").AppendThis(uMd).AppendNormalThis("...")
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

	md = md.AppendBoldThis("• Is banned: ")
	md.AppendMonoThis(ws.YesOrNo(targetUser.Banned)).ElThis()
	md.AppendBoldThis("• Crime Coefficient: ")
	md.AppendMonoThis(strconv.Itoa(targetUser.CrimeCoefficient)).ElThis()

	msg, err = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode: sv.MarkDownV2,
	})
	if err != nil {
		logging.UnexpectedError(err)
		return
	}

	time.Sleep(2 * time.Second)

	md.ElThis().AppendBoldThis("Verdict: ").AppendThis(uMd)
	md.AppendNormalThis(" cannot be assigned as " + p + " because their crime coefficient is ")
	se, cc := targetUser.EstimateCrimeCoefficientSep()
	md.AppendNormalThis(se).AppendMonoThis(cc).ElThis()
	md.AppendBoldThis("Attached reason: ").AppendMonoThis(targetUser.Reason)
	_, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode: sv.MarkDownV2,
	})
}

func showUserAssigned(b *gotgbot.Bot, ctx *ext.Context, aValue *AssignValue) {
	var err error
	var md, uMd mdparser.WMarkDown
	namae := aValue.targetChat.FirstName
	uMd = mdparser.GetUserMention(namae, aValue.targetChat.Id)
	strId := strconv.FormatInt(aValue.targetChat.Id, 10)
	md = mdparser.GetBold("\u200D • User: ").AppendThis(uMd).ElThis()
	md.AppendBoldThis(" • ID: ").AppendMonoThis(strId).ElThis()
	md.AppendBoldThis(" • Is banned: ").AppendMonoThis("No").ElThis()
	md.AppendBoldThis(" • Crime Coefficient: ")
	md.AppendMonoThis(aValue.targer.EstimateCrimeCoefficient())
	md.ElThis()
	// let the goroutine sleep for 1 second
	time.Sleep(2 * time.Second)
	aValue.msg, err = aValue.msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
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
		md.AppendNormalThis("✳️ ").AppendThis(uMd).AppendNormalThis(" has now been assigned as ")
		md.AppendBoldThis(aValue.perm)
		md.AppendNormalThis("!\nTheir dominator and token have been sent to their ")
		md.AppendHyperLinkThis("PM", "http://t.me/"+b.Username).AppendNormalThis(".")
		_, _ = aValue.msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
			ParseMode:             sv.MarkDownV2,
			DisableWebPagePreview: true,
		})
	} else {
		md = mdparser.GetBold("Assignment request has been sent to Sibyl System! \n")
		md.AppendThis(mdBack)
		md.AppendNormalThis("✳️ ").AppendThis(uMd).AppendNormalThis(" will be assigned as ")
		md.AppendBoldThis(aValue.perm).AppendNormalThis(" after verification.")
		_, _ = aValue.msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
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
			ParseMode:   sv.MarkDownV2,
			ReplyMarkup: aValue.getAssignmentButton(),
		}

		for _, chat := range bases {
			sendRequestMessage(chat, text, opts)
		}
	}
}

func sendRequestMessage(chat int64, text string, opts *gotgbot.SendMessageOpts) {
	_, err := sv.HelperBot.SendMessage(chat, text, opts)
	if err != nil {
		logging.Debug("Tried to send message to ", chat, err)
	}
}
