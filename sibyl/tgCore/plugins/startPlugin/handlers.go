package startPlugin

import (
	"strconv"
	"strings"
	"time"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
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
		return ext.EndGroups
	}

	sv.RateLimiter.AddCustomIgnore(user.Id, 5*time.Minute, true)
	// user is already in the database
	if theUser.Banned {
		go startForBanned(b, ctx, theUser, t)
	} else {
		go startForNotBanned(b, ctx, theUser, t)
	}
	return ext.EndGroups
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

	md = welcomeMd.Normal("\nCymatic Scan results:")
	md.Bold("\n • User").Normal(": ")
	md.Mention(user.FirstName, user.Id)
	md.Bold("\n • ID").Normal(": ")
	md.Mono(strconv.FormatInt(user.Id, 10))
	if user.Username != "" {
		md.Bold("\n • Username").Normal(": ")
		md.Mono(user.Username)
	}
	md.Bold("\n • Is banned").Normal(": ")
	md.Mono(ws.YesOrNo(u.Banned))
	md.Bold("\n • Status").Normal(": ")
	md.Mono(t.GetTitleStringPermission())
	md.Bold("\n • Crime Coefficient").Normal(": ")
	md.Mono(u.GetStringCrimeCoefficient())
	md.Bold("\n • Ban Reason(s)").Normal(": ")
	md.AppendThis(u.FormatFlags())
	md.Bold("\n • Description").Normal(": ")
	md.Mono(u.Reason)
	if u.BanSourceUrl != "" {
		md.Bold("\n • Scan source").Normal(": ")
		md.Normal(u.BanSourceUrl)
	}

	var markup gotgbot.InlineKeyboardMarkup

	if !u.CanTryAppealing() {
		md.Normal("\n\nYour ban is not appealable.")
		_, _, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
			ParseMode: sv.MarkDownV2,
		})
		sv.RateLimiter.RemoveCustomIgnore(user.Id)
		return
	}

	md.Normal("\n\nSince this is your first time we can allow you")
	md.Normal(" a one time exception provided that you will not")
	md.Normal(" repeat this ever again.")
	markup.InlineKeyboard = makeFirstPageAppealButtons()
	_, _, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:             sv.MarkDownV2,
		DisableWebPagePreview: true,
		ReplyMarkup:           markup,
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
	md = welcomeMd.Normal("Cymatic Scan results:")
	md.Bold("\n • User").Normal(": ")
	md.Mention(user.FirstName, user.Id)
	md.Bold("\n • ID").Normal(": ")
	md.Mono(strconv.FormatInt(user.Id, 10))
	md.Bold("\n • Is banned").Normal(": ")
	md.Mono(ws.YesOrNo(u.Banned))
	md.Bold("\n • Status").Normal(": ")
	md.Mono(t.GetTitleStringPermission())
	md.Bold("\n • Crime Coefficient").Normal(": ")
	md.Mono(u.EstimateCrimeCoefficient()).ElThis()
	_, _, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:   sv.MarkDownV2,
		ReplyMarkup: *markup,
	})
	sv.RateLimiter.RemoveCustomIgnore(user.Id)
}

func showAppealDetails(b *gotgbot.Bot, ctx *ext.Context, u *sv.User) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	md := mdparser.GetUserMention(user.FirstName, user.Id)
	md.Normal("! You were blacklisted on ")
	md.Link("Sibyl System", "https://t.me/SibylSystem/13")
	md.Normal(" for ")
	md.AppendThis(u.FormatCuteFlags())
	md.AppendThis(u.FormatDetailStrings(true))
	md.Normal("Such type of actions are often unwanted and unwelcome around Sibyl.")
	md.Normal(" Please do note that should this ever happen again your ban will be")
	md.Normal(" swift and its damage, measurable on the richter scale!")
	md.Normal("\nClick the button below to confirm that you understand this")
	md.Normal(" and if you have questions please click the Support button")
	md.Normal(" to take your query to the bureau.")
	_, _, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
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
	md.Normal("! You have been unbanned!")
	md.Bold("\nNote: ").Normal("You will ")
	md.Bold("not ").Normal("be able to appeal this ban again.")
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
		logMd.Bold("\n • User").Normal(": ")
		logMd.Mention(user.FirstName, user.Id)
		logMd.Normal("[").Mono(ws.ToBase10(user.Id)).Normal("]")
		if user.Username != "" {
			logMd.Bold("\n • Username").Normal(": ")
			logMd.Mono(user.Username)
		}
		logMd.Bold("\n • Crime Coefficient").Normal(": ")
		logMd.Mono(uPre.GetStringCrimeCoefficient())
		logMd.Bold("\n • Reason(s)").Normal(": ")
		logMd.AppendThis(uPre.FormatFlags())
		logMd.Bold("\n • Description").Normal(": ")
		logMd.Normal(uPre.Reason)
		logMd.Bold("\n • Scan Date").Normal(": ")
		logMd.Mono(uPre.GetDateAsShort())
		logMd.Bold("\n • Appeal Date").Normal(": ")
		logMd.Mono(time.Now().Format(sv.AppealLogDateFormat))
		if uPre.BanSourceUrl != "" {
			logMd.Bold("\n • Scan source").Normal(": ")
			logMd.Normal(uPre.BanSourceUrl)
		}

		go utils.SendMultipleMessages(chats, logMd.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:             sv.MarkDownV2,
			DisableWebPagePreview: true,
		})
	}

	_, _, _ = msg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
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
			md.Normal(", you are no longer able to use auto appeal system.")
			md.Normal("\nPlease take your questions to @PublicSafetyBureau")
			md.Normal(" if you want an unban.")
			_, _, _ = ctx.EffectiveMessage.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
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
			md.Normal(", you are no longer able to use auto appeal system.\n")
			md.Normal("Please take your questions to @PublicSafetyBureau if you want an unban.")
			_, _, _ = ctx.EffectiveMessage.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
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
