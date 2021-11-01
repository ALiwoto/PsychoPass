package tokenPlugin

import (
	"strconv"

	"github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/ALiwoto/mdparser/mdparser"
	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/hashing"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	startCmd := handlers.NewCommand(StartCmd, startHandler)
	createCmd := handlers.NewCommand(CreateCmd, startHandler)
	newCmd := handlers.NewCommand(NewCmd, startHandler)
	revokeCmd := handlers.NewCommand(RevokeCmd, revokeHandler)
	assignCmd := handlers.NewCommand(AssignCmd, assignHandler)
	startCmd.Triggers = t
	createCmd.Triggers = t
	newCmd.Triggers = t
	revokeCmd.Triggers = t
	assignCmd.Triggers = t
	d.AddHandler(startCmd)
	d.AddHandler(createCmd)
	d.AddHandler(newCmd)
	d.AddHandler(revokeCmd)
	d.AddHandler(assignCmd)
}

// startHandler is the handler for the /start command.
// It will send a message to the user with their token.
func startHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveChat.Type != "private" {
		return ext.EndGroups
	}

	user := ctx.EffectiveUser
	t, err := database.GetTokenFromId(user.Id)
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
}

// revokeHandler is the handler for the /revoke command.
// It will revoke the token of the user.
func revokeHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	t, err := database.GetTokenFromId(user.Id)
	if err != nil || t == nil {
		return ext.EndGroups
	}

	database.RevokeTokenHash(t, hashing.GetUserToken(user.Id))

	md := mdparser.GetNormal("Hi ").AppendMentionThis(user.FirstName, user.Id)
	md.AppendNormalThis(" !\nHere is your new token:\n")
	md.AppendMonoThis(t.Hash).AppendNormalThis("\n\n")
	md.AppendBoldThis("Please don't share this token with anyone!")
	if t.HasRole() {
		md.AppendItalicThis("You are a valid").AppendNormalThis(" ")
		md.AppendItalicThis(t.GetStringPermission()).AppendNormal(".")
	}

	b.SendMessage(user.Id, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode: sv.MarkDownV2,
	})

	return ext.EndGroups
}

// assignHandler is the handler for the /assign command.
// It will change the permission of the token of the user.
func assignHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	if sv.IsInvalidID(user.Id) {
		return ext.EndGroups
	}

	t, err := database.GetTokenFromId(user.Id)
	if err != nil || t == nil || !t.CanChangePermission() {
		return ext.EndGroups
	}

	args := strongStringGo.Split(msg.Text, " ", "\n", "\t", ",", "-", "~")
	if len(args) < 2 {
		// show help.
		md := mdparser.GetNormal("Dear ").AppendMentionThis(user.FirstName, user.Id)
		md.AppendNormalThis(", this command lets you authorise dominator access for ")
		md.AppendHyperLinkThis("Sibyl", "http://t.me/SibylSystem")
		md.AppendNormalThis("\nRun command again in the following format")
		md.AppendBoldThis("\nYour options are:")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign inspector ID")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign enforcer ID")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign civilian ID")
		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	}

	perm, err := sv.ConvertToPermission(args[1])
	if err != nil {
		md := mdparser.GetNormal("Invalid permission provided: ")
		md.AppendMonoThis(args[1])
		md.AppendNormalThis("!\nHere is a list of possible permissions:")
		md.AppendNormalThis("\n- ").AppendMonoThis("inspector")
		md.AppendNormalThis("\n- ").AppendMonoThis("enforcer")
		md.AppendNormalThis("\n- ").AppendMonoThis("civilian")

		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	} else if perm.IsOwner() {
		md := mdparser.GetNormal("This decision is of Sibyl to make!")
		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	}

	var replied = msg.ReplyToMessage != nil && msg.ReplyToMessage.From != nil
	var preReplied = replied // reserved value
	var isBot bool
	var targetId int64

	if len(args) < 3 && !replied {
		// show help.
		md := mdparser.GetNormal("You need to provide a user ID for this command.")
		md.AppendBoldThis("\nYour options are:")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign inspector ID")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign enforcer ID")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign civilian ID")

		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	} else if len(args) > 2 && replied {
		// ignore this if the message itself contains the user ID.
		replied = false
	}

	if replied {
		targetId = msg.ReplyToMessage.From.Id
		isBot = msg.ReplyToMessage.From.IsBot
	} else {
		targetId, err = strconv.ParseInt(args[2], 10, 64)
	}

	if err != nil || targetId == 0 {
		if !replied && !preReplied {
			md := mdparser.GetNormal("Invalid ID provided: ")
			md.AppendMonoThis(args[2])
			md.AppendNormalThis("!\nPlease make sure the target's ID is valid.")

			_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
				ParseMode:                sv.MarkDownV2,
				AllowSendingWithoutReply: true,
				DisableWebPagePreview:    true,
			})
			if err != nil {
				logging.UnexpectedError(err)
			}

			return ext.EndGroups
		} else if preReplied {
			targetId = msg.ReplyToMessage.From.Id
			isBot = msg.ReplyToMessage.From.IsBot
		}
	}

	if targetId == user.Id {
		md := mdparser.GetNormal("You can't change your own permissions.")
		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	} else if targetId == b.Id {
		md := mdparser.GetNormal("You can't change my permissions.")
		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	} else if isBot || sv.IsInvalidID(targetId) {
		md := mdparser.GetNormal("Dominator authorisation is only for humans.")
		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	}

	var md mdparser.WMarkDown
	u, _ := database.GetTokenFromId(targetId)
	if u != nil {
		if u.IsOwner() {
			md = mdparser.GetNormal("This decision is of Sibyl to make.\n")
			md.AppendNormalThis("Please try another user ID.\n")
		} else if u.Permission == perm {
			md = mdparser.GetNormal("The user ")
			md.AppendMentionThis(strconv.FormatInt(targetId, 10), u.UserId)
			md.AppendNormalThis(" is already assigned ")
			md.AppendMonoThis(perm.GetStringPermission()).AppendNormal(".")
		} else {
			pre := u.Permission
			database.UpdateTokenPermission(u, perm)
			md = mdparser.GetNormal("The user ")
			md.AppendMentionThis(strconv.FormatInt(targetId, 10), u.UserId)
			md.AppendNormalThis(" has been assigned the permission ")
			md.AppendMonoThis(perm.GetStringPermission()).AppendNormal(".")
			md.AppendItalicThis("\n\nThe previous permission was ")
			md.AppendMonoThis(pre.GetStringPermission()).AppendNormal(".")
		}
	} else {
		md = mdparser.GetUserMention(strconv.FormatInt(targetId, 10), targetId) //\u200D
		//md = mdparser.GetUserMention(strconv.FormatInt(targetId, 10), targetId) //\u200D
		md.AppendNormalThis(" needs to start me in PM to connect to Sibyl.")
		_, err = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	}

	mdback := mdparser.GetNormal("Your permission has been changed to ")
	mdback.AppendMonoThis(u.GetStringPermission())
	mdback.AppendNormalThis("!\n\nHere is your token:\n")
	mdback.AppendMonoThis(u.Hash).AppendNormalThis("\n\n")
	mdback.AppendBoldThis("Please don't share this token with anyone!")
	_, err = b.SendMessage(targetId, mdback.ToString(), &gotgbot.SendMessageOpts{
		ParseMode: sv.MarkDownV2,
	})
	if err != nil {
		md = mdparser.GetUserMention(strconv.FormatInt(targetId, 10), targetId)
		md.AppendNormalThis(" needs to start me in PM to connect to Sibyl.")
	}

	_, err = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode:                sv.MarkDownV2,
		AllowSendingWithoutReply: true,
		DisableWebPagePreview:    true,
	})
	if err != nil {
		logging.UnexpectedError(err)
	}

	return ext.EndGroups
}
