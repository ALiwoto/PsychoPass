package tokenPlugin

import (
	"strconv"
	"time"

	"github.com/ALiwoto/mdparser/mdparser"
	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func showUserBanned(b *gotgbot.Bot, ctx *ext.Context, targetUser *sv.User, p string) {
	var err error
	msg := ctx.EffectiveMessage
	uMd := mdparser.GetUserMention(strconv.FormatInt(targetUser.UserID, 10), targetUser.UserID)
	md := mdparser.GetNormal("Scanning ").AppendThis(uMd).AppendNormalThis("...")
	msg, err = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode:                sv.MarkDownV2,
		AllowSendingWithoutReply: true,
		DisableWebPagePreview:    true,
	})
	if err != nil {
		logging.UnexpectedError(err)
		return
	}

	time.Sleep(3 * time.Second)

	md = mdparser.GetBold("User: ").Append(uMd).ElThis()
	md.AppendBoldThis("Is banned: ").AppendMonoThis(strconv.FormatBool(targetUser.Banned)).ElThis()
	md.AppendBoldThis("Crime Coefficient: ")
	md.AppendMonoThis(strconv.Itoa(targetUser.CrimeCoefficient)).ElThis()

	msg, err = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode: sv.MarkDownV2,
	})
	if err != nil {
		logging.UnexpectedError(err)
		return
	}

	time.Sleep(3 * time.Second)

	md.ElThis().AppendBoldThis("Verdict: ").AppendThis(uMd)
	md.AppendNormalThis("cannot be assigned as " + p + " because their crime coefficient is over ")
	md.AppendMonoThis(targetUser.EstimateCrimeCoefficient()).ElThis()
	md.AppendBoldThis("Attached reason: ").AppendMonoThis(targetUser.Reason)
	_, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode: sv.MarkDownV2,
	})
}

/*
md.AppendHyperLinkThis("Sibyl System.", "http://t.me/SibylSystem").ElThis()
	md.AppendBoldThis("Reason: ").AppendNormalThis(targetUser.Reason)
	if len(targetUser.BanFlags) > 0 {
		md.AppendNormalThis(" [ ")
		for i, flag := range targetUser.BanFlags {
			if i != 0 {
				md.AppendNormalThis(", ")
			}
			md.AppendMonoThis(string(flag))
		}
		md.AppendNormalThis(" ]")
	}
	md.ElThis()
*/
