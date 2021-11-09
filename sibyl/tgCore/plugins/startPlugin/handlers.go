package startPlugin

import (
	"strconv"
	"time"

	"github.com/ALiwoto/mdparser/mdparser"
	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// startHandler is the handler for the /start command.
// It will send a message to the user with their token.
func startHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveChat.Type != "private" {
		return ext.EndGroups
	}

	var t *sv.Token
	var theUser *sv.User
	var err error
	user := ctx.EffectiveUser

	t, err = database.GetTokenFromId(user.Id)
	if err != nil {
		logging.UnexpectedError(err)
		return ext.EndGroups
	}

	if t == nil {
		// should create a new token
		t, err = utils.CreateToken(user.Id, sv.NormalUser)
		if err != nil {
			logging.UnexpectedError(err)
			return ext.EndGroups
		}
	}

	theUser, err = database.GetUserFromId(user.Id)
	if theUser == nil && err == nil {
		// err is nil and user is nil as well: user not found.
		// save it to db and send the cymatic scan result.
		theUser = database.ForceInsert(user.Id, t.Permission)
	} else if err != nil {
		// internal database error?
		logging.UnexpectedError(err)
		return err
	}

	// user is already in the database
	if theUser.Banned {
		go startForBanned(b, ctx, theUser, t)
		return ext.EndGroups
	} else {
		go startForNotBanned(b, ctx, theUser, t)
		return ext.EndGroups
	}
	/*


		md := mdparser.GetNormal("Hi ").AppendMentionThis(user.FirstName, user.Id)
		md.AppendNormalThis(" !\nHere is your token:\n")
		md.AppendMonoThis(t.Hash).AppendNormalThis("\n\n")
		md.AppendBoldThis("Please don't share this token with anyone!")
		if t.HasRole() {
			md.AppendItalicThis("\nYou are a valid").AppendNormalThis(" ")
			md.AppendItalicThis(t.GetStringPermission()).AppendNormal(".")
		}

		b.SendMessage(user.Id, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: sv.MarkDownV2,
		})

		return ext.EndGroups
	*/
}

func startForBanned(b *gotgbot.Bot, ctx *ext.Context, u *sv.User, t *sv.Token) {
	message := ctx.EffectiveMessage
	user := ctx.EffectiveUser
	welcomeMd := mdparser.GetNormal("Welcome to Sibyl System!\n")
	md := welcomeMd.AppendNormal("Please wait while we finish your cymatic scan...")
	msg, err := message.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode: sv.MarkDownV2,
	})
	if err != nil {
		// most probably the user has deleted their message.
		// we don't need to do anything.
		return
	}
	time.Sleep(3 * time.Second)
	md = welcomeMd.AppendNormalThis("Cymatic Scan results:\n")
	md.AppendBoldThis(" • User").AppendNormalThis(": ")
	md.AppendMentionThis(user.FirstName+"\n", user.Id)
	md.AppendBoldThis(" • ID").AppendNormalThis(": ")
	md.AppendMonoThis(strconv.FormatInt(user.Id, 10)).ElThis()
	md.AppendBoldThis(" • Is banned").AppendNormalThis(": ")
	md.AppendMonoThis(strconv.FormatInt(user.Id, 10)).ElThis()
	md.AppendBoldThis(" • Status").AppendNormalThis(": ")
	md.AppendBoldThis(" • Status").AppendNormalThis(": ")
	md.AppendMonoThis(t.GetTitleStringPermission()).ElThis()
	md.AppendBoldThis(" • Crime Coefficient").AppendNormalThis(": ")
	md.AppendMonoThis(u.EstimateCrimeCoefficient()).ElThis()
	msg, err = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode: sv.MarkDownV2,
	})
	if err != nil || msg == nil {
		// most probably the user has deleted our message.
		// we don't need to do anything.
		return
	}

}

func startForNotBanned(b *gotgbot.Bot, ctx *ext.Context, u *sv.User, t *sv.Token) {
	message := ctx.EffectiveMessage
	user := ctx.EffectiveUser
	welcomeMd := mdparser.GetNormal("Welcome to Sibyl System!\n")
	md := welcomeMd.AppendNormal("Please wait while we finish your cymatic scan...")
	msg, err := message.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode: sv.MarkDownV2,
	})
	if err != nil || msg == nil {
		// most probably the user has deleted their message.
		// we don't need to do anything.
		return
	}

	time.Sleep(3 * time.Second)

	markup := &gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: makeNormalButtons(),
	}
	md = welcomeMd.AppendNormalThis("Cymatic Scan results:\n")
	md.AppendBoldThis(" • User").AppendNormalThis(": ")
	md.AppendMentionThis(user.FirstName+"\n", user.Id)
	md.AppendBoldThis(" • ID").AppendNormalThis(": ")
	md.AppendMonoThis(strconv.FormatInt(user.Id, 10)).ElThis()
	md.AppendBoldThis(" • Is banned").AppendNormalThis(": ")
	md.AppendMonoThis(strconv.FormatInt(user.Id, 10)).ElThis()
	md.AppendBoldThis(" • Status").AppendNormalThis(": ")
	md.AppendMonoThis(t.GetTitleStringPermission()).ElThis()
	md.AppendBoldThis(" • Crime Coefficient").AppendNormalThis(": ")
	md.AppendMonoThis(u.EstimateCrimeCoefficient()).ElThis()
	msg, err = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:   sv.MarkDownV2,
		ReplyMarkup: *markup,
	})
	if err != nil || msg == nil {
		// most probably the user has deleted our message.
		// we don't need to do anything.
		return
	}
}

func makeNormalButtons() [][]gotgbot.InlineKeyboardButton {
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
	return rows
}
