package tokenPlugin

import (
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/hashing"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/logging"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/database"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	startCmd := handlers.NewCommand(StartCmd, startHandler)
	createCmd := handlers.NewCommand(CreateCmd, startHandler)
	newCmd := handlers.NewCommand(NewCmd, startHandler)
	revokeCmd := handlers.NewCommand(RevokeCmd, revokeHandler)
	startCmd.Triggers = t
	createCmd.Triggers = t
	newCmd.Triggers = t
	revokeCmd.Triggers = t
	d.AddHandler(startCmd)
	d.AddHandler(createCmd)
	d.AddHandler(newCmd)
	d.AddHandler(revokeCmd)
}

func startHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	t, err := database.GetTokenFromId(user.Id)
	if err != nil {
		logging.UnexpectedError(err)
	}

	md := mdparser.GetNormal("Hi ").AppendMentionThis(user.FirstName, user.Id)
	md.AppendNormalThis(" !\nHere is your token:\n")
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

func revokeHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	t, err := database.GetTokenFromId(user.Id)
	if err != nil {
		logging.UnexpectedError(err)
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
