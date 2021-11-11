package startPlugin

import (
	"strconv"
	"strings"
	"time"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylConfig"
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

	sv.RateLimiter.AddCustomIgnore(user.Id, 5*time.Minute, true)
	// user is already in the database
	if theUser.Banned {
		go startForBanned(b, ctx, theUser, t)
		return ext.EndGroups
	} else {
		go startForNotBanned(b, ctx, theUser, t)
		return ext.EndGroups
	}
}

func startForBanned(b *gotgbot.Bot, ctx *ext.Context, u *sv.User, t *sv.Token) {
	message := ctx.EffectiveMessage
	user := ctx.EffectiveUser
	welcomeMd := mdparser.GetNormal("Welcome to Sibyl System!")
	md := welcomeMd.AppendNormal("\nPlease wait while we finish your cymatic scan...")
	msg, err := message.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode: sv.MarkDownV2,
	})
	if err != nil {
		// most probably the user has deleted their message.
		// we don't need to do anything.
		return
	}

	time.Sleep(3 * time.Second)

	md = welcomeMd.AppendNormalThis("Cymatic Scan results:")
	md.AppendBoldThis("\n • User").AppendNormalThis(": ")
	md.AppendMentionThis(user.FirstName, user.Id)
	md.AppendBoldThis("\n • ID").AppendNormalThis(": ")
	md.AppendMonoThis(strconv.FormatInt(user.Id, 10))
	md.AppendBoldThis("\n • Is banned").AppendNormalThis(": ")
	md.AppendMonoThis(ws.YesOrNo(u.Banned))
	md.AppendBoldThis("\n • Status").AppendNormalThis(": ")
	md.AppendMonoThis(t.GetTitleStringPermission())
	md.AppendBoldThis("\n • Crime Coefficient").AppendNormalThis(": ")
	md.AppendMonoThis(u.GetStringCrimeCoefficient())
	md.AppendBoldThis("\n • Ban Reason(s)").AppendNormalThis(": ")
	md.AppendThis(u.FormatFlags())
	md.AppendBoldThis("\n • Description").AppendNormalThis(": ")
	md.AppendMonoThis(u.Reason)

	var markup gotgbot.InlineKeyboardMarkup

	if !u.CanTryAppealing() {
		md.AppendNormalThis("\n\nYour ban is not appealable.")
		markup.InlineKeyboard = makeSingleAppealButtons()
		_, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
			ParseMode:   sv.MarkDownV2,
			ReplyMarkup: markup,
		})
		sv.RateLimiter.RemoveCustomIgnore(user.Id)
		return
	}

	md.AppendNormalThis("\n\nSince this is your first time we can allow you")
	md.AppendNormalThis(" a one time exception provided that you will not")
	md.AppendNormalThis(" repeat this ever again.")
	markup.InlineKeyboard = makeFirstPageAppealButtons()
	_, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:   sv.MarkDownV2,
		ReplyMarkup: markup,
	})
	sv.RateLimiter.RemoveCustomIgnore(user.Id)
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
	md = welcomeMd.AppendNormalThis("Cymatic Scan results:")
	md.AppendBoldThis("\n • User").AppendNormalThis(": ")
	md.AppendMentionThis(user.FirstName, user.Id)
	md.AppendBoldThis("\n • ID").AppendNormalThis(": ")
	md.AppendMonoThis(strconv.FormatInt(user.Id, 10))
	md.AppendBoldThis("\n • Is banned").AppendNormalThis(": ")
	md.AppendMonoThis(ws.YesOrNo(u.Banned))
	md.AppendBoldThis("\n • Status").AppendNormalThis(": ")
	md.AppendMonoThis(t.GetTitleStringPermission())
	md.AppendBoldThis("\n • Crime Coefficient").AppendNormalThis(": ")
	md.AppendMonoThis(u.EstimateCrimeCoefficient()).ElThis()
	_, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:   sv.MarkDownV2,
		ReplyMarkup: *markup,
	})
	sv.RateLimiter.RemoveCustomIgnore(user.Id)
}

func showAppealDetails(b *gotgbot.Bot, ctx *ext.Context, u *sv.User) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	md := mdparser.GetUserMention(user.FirstName, user.Id)
	md.AppendNormalThis("! You were blacklisted on ")
	md.AppendHyperLinkThis("Sibyl System", "https://t.me/SibylSystem")
	md.AppendNormalThis(" for ")
	md.AppendThis(u.FormatCuteFlags())
	md.AppendThis(u.FormatDetailStrings(true))
	md.AppendNormalThis("Such type of actions are often unwanted and unwelcome around Sibyl.")
	md.AppendNormalThis(" Please do note that should this ever happen again your ban will be")
	md.AppendNormalThis(" swift and its damage, measurable on the richter scale!\n")
	md.AppendNormalThis("Click the button below to confirm that you understand this")
	md.AppendNormalThis(" and if you have questions please click the Support button")
	md.AppendNormalThis(" to take your query to the bureau.")
	_, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:             sv.MarkDownV2,
		DisableWebPagePreview: true,
		ReplyMarkup:           makeDetailsPageAppealButtons(true),
	})
	return ext.EndGroups
}

func showAppealDoneDetails(b *gotgbot.Bot, ctx *ext.Context, u *sv.User) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	md := mdparser.GetUserMention(user.FirstName, user.Id)
	md.AppendNormalThis("! You have been unbanned!")
	md.AppendBoldThis("\nNote: ").AppendNormalThis("You will ")
	md.AppendBoldThis("not ").AppendNormalThis("be able to appeal this ban again.")
	markup := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: makeNormalButtons(),
	}

	// copy the banned user's info, so database package doesn't change it in future.
	pre := *u

	// lift the ban.
	database.RemoveUserBan(u, false)

	// send the log message to the log channel(s).
	chats := sibylConfig.GetAppealLogChatIds()
	if len(chats) > 0 {
		uPre := &pre
		logMd := mdparser.GetNormal("#AutoAppeal")
		logMd.AppendBoldThis("\n • User").AppendNormalThis(": ")
		logMd.AppendMentionThis(user.FirstName, user.Id)
		logMd.AppendBoldThis("\n • Crime Coefficient").AppendNormalThis(": ")
		logMd.AppendMonoThis(uPre.GetStringCrimeCoefficient())
		logMd.AppendBoldThis("\n • Reason(s)").AppendNormalThis(": ")
		logMd.AppendThis(uPre.FormatFlags())
		logMd.AppendBoldThis("\n • Description").AppendNormalThis(": ")
		logMd.AppendNormalThis(uPre.Reason)
		logMd.AppendBoldThis("\n • Scan Date").AppendNormalThis(": ")
		logMd.AppendMonoThis(uPre.GetDateAsShort())
		logMd.AppendBoldThis("\n • Appeal Date").AppendNormalThis(": ")
		logMd.AppendMonoThis(time.Now().Format(sv.AppealLogDateFormat))

		go func() {
			for _, chatId := range chats {
				_, _ = b.SendMessage(chatId, logMd.ToString(), &gotgbot.SendMessageOpts{
					ParseMode: sv.MarkDownV2,
				})
			}
		}()
	}

	_, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:             sv.MarkDownV2,
		DisableWebPagePreview: true,
		ReplyMarkup:           markup,
	})
	return ext.EndGroups
}

func appealCallBackQuery(cq *gotgbot.CallbackQuery) bool {
	return strings.HasPrefix(cq.Data, AutoAppealCbPrefix)
}

func appealCallBackResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	action := strings.TrimPrefix(ctx.CallbackQuery.Data, AutoAppealCbPrefix)
	switch action {
	case CloseCbData:
		_, _ = ctx.EffectiveMessage.Delete(b)
	case firstAcceptCbData:
		date := time.Unix(ctx.CallbackQuery.Message.Date, 0)
		if time.Since(date) > time.Minute*5 {
			_, _ = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text:      "You took too long to respond. Please try again.",
				ShowAlert: true,
			})
			_, _ = ctx.EffectiveMessage.Delete(b)
			return ext.EndGroups
		}
		u, err := database.GetUserFromId(ctx.CallbackQuery.From.Id)
		if u == nil || err != nil || !u.Banned {
			_, _ = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text:      "You are not banned.",
				ShowAlert: true,
			})
			_, _ = ctx.EffectiveMessage.Delete(b)
			return ext.EndGroups
		}
		if !u.CanTryAppealing() || !u.CanAppeal() {
			user := ctx.EffectiveUser
			md := mdparser.GetUserMention(user.FirstName, user.Id)
			md.AppendNormalThis(", you are no longer able to use auto appeal system.\n")
			md.AppendNormalThis("Please take your questions to @PublicSafetyBureau if you want an unban.")
			_, _ = ctx.EffectiveMessage.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
				ParseMode:             sv.MarkDownV2,
				DisableWebPagePreview: true,
				ReplyMarkup:           makeDetailsPageAppealButtons(false),
			})
			return ext.EndGroups
		}
		return showAppealDetails(b, ctx, u)
	case detailsAcceptCbData:
		date := time.Unix(ctx.CallbackQuery.Message.Date, 0)
		if time.Since(date) > time.Minute*5 {
			_, _ = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text:      "You took too long to respond. Please try again.",
				ShowAlert: true,
			})
			_, _ = ctx.EffectiveMessage.Delete(b)
			return ext.EndGroups
		}
		u, err := database.GetUserFromId(ctx.CallbackQuery.From.Id)
		if u == nil || err != nil || !u.Banned {
			_, _ = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text:      "You are not banned.",
				ShowAlert: true,
			})
			_, _ = ctx.EffectiveMessage.Delete(b)
			return ext.EndGroups
		}
		if !u.CanTryAppealing() || !u.CanAppeal() {
			user := ctx.EffectiveUser
			md := mdparser.GetUserMention(user.FirstName, user.Id)
			md.AppendNormalThis(", you are no longer able to use auto appeal system.\n")
			md.AppendNormalThis("Please take your questions to @PublicSafetyBureau if you want an unban.")
			_, _ = ctx.EffectiveMessage.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
				ParseMode:             sv.MarkDownV2,
				DisableWebPagePreview: true,
				ReplyMarkup:           makeDetailsPageAppealButtons(false),
			})
			return ext.EndGroups
		}
		return showAppealDoneDetails(b, ctx, u)
	}
	return ext.EndGroups
}
