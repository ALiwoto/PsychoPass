package tokenPlugin

import (
	"strings"

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
		md.AppendMonoThis(t.GetStringPermission()).AppendNormal(".")
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
		md.AppendMonoThis(t.GetStringPermission()).AppendNormal(".")
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
	t, err := database.GetTokenFromId(user.Id)
	if err != nil || t == nil || t.CanChangePermission() {
		return ext.EndGroups
	}

	args := strings.Split(msg.Text, " ")
	if len(args) < 2 {
		// show help.
		md := mdparser.GetNormal("Dear ").AppendMentionThis(user.FirstName, user.Id)
		md.AppendNormalThis(" ").AppendNormalThis(", this command lets you assign users to ")
		md.AppendHyperLinkThis("Sibyl", "http://t.me/SibylSystem")
		md.AppendNormalThis("\nPlease provide a type with the command.")
		md.AppendBoldThis("Your options are:")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign inspector ID")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign enforcer ID")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign civilian ID")
		msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
		})
		return ext.EndGroups
	}

	return ext.EndGroups
}
