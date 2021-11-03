package tokenPlugin

import (
	"strconv"
	"strings"
	"time"

	"github.com/ALiwoto/mdparser/mdparser"
	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

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

	md = suMd.AppendBoldThis("• Is banned: ")
	md.AppendMonoThis(strconv.FormatBool(targetUser.Banned)).ElThis()
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

func showUserAssigned(b *gotgbot.Bot, ctx *ext.Context,
	targetChat *gotgbot.Chat, perm string, msg *gotgbot.Message, targer *sv.User) {
	var err error
	var md, uMd mdparser.WMarkDown
	namae := targetChat.FirstName
	uMd = mdparser.GetUserMention(namae, targetChat.Id)
	strId := strconv.FormatInt(targetChat.Id, 10)
	md = mdparser.GetBold(" • User: ").AppendThis(uMd).ElThis()
	md.AppendBoldThis(" • ID: ").AppendMonoThis(strId).ElThis()
	md.AppendBoldThis(" • Is banned: ").AppendMonoThis("false").ElThis()
	md.AppendBoldThis(" • Crime Coefficient: ").AppendMonoThis(targer.EstimateCrimeCoefficient())
	md.ElThis()
	// let the goroutine sleep for 1 second
	time.Sleep(time.Second)
	msg, err = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:             sv.MarkDownV2,
		DisableWebPagePreview: true,
	})
	if err != nil {
		logging.UnexpectedError(err)
		return
	}

	time.Sleep(2 * time.Second)

	md = mdparser.GetBold("Assigned Successfully! ").ElThis().AppendThis(md).ElThis()
	md.AppendNormalThis("✳️ ").AppendThis(uMd).AppendNormalThis(" has now been assigned as ")
	md.AppendBoldThis(perm).ElThis()
	md.AppendNormalThis("Their dominator and token have been sent to their ")
	md.AppendHyperLinkThis("PM", "http://t.me/"+b.Username).AppendNormalThis(".")
	_, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:             sv.MarkDownV2,
		DisableWebPagePreview: true,
	})
}
