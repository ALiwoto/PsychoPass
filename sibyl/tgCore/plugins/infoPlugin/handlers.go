package infoPlugin

import (
	"strconv"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	statsCmd := handlers.NewCommand(StatsCmd, StatsHandler)
	statsCmd.Triggers = t
	d.AddHandler(statsCmd)
}

func StatsHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	t, err := database.GetTokenFromId(user.Id)
	if err != nil || t == nil || !t.CanChangePermission() {
		return ext.EndGroups
	}

	/*
	   Coefficients and Flags

	   ==== Flags     -  ========
	   Range 0-100 (No bans) (Dominator Locked)
	   • Civilian     - 0-80
	   • Past Banned  - 81-100
	   ==============
	   Range 100-300 (Auto-mute) (Non-lethal Paralyzer)
	   • TROLLING     - 101-125 - Trolling
	   • SPAM         - 126-200 - Spam/Unwanted Promotion
	   • EVADE        - 201-250 - Ban Evade using alts
	   x-------x
	   Manual Revert
	   • CUSTOM       - 251-300 - Any Custom reason
	   x-------x
	   ==============
	   Range 300+ (Ban on Sight) (Lethal Eliminator)
	   • PSYCHOHAZARD - 301-350 - Bulk banned due to some bad users
	   • MALIMP       - 351-400 - Malicious Impersonation
	   • NSFW         - 401-450 - Sending NSFW Content in SFW
	   • RAID         - 451-500 - Bulk join raid to vandalize
	   • MASSADD      - 501-600 - Mass adding to group/channel
	   ==============
	*/

	tbanned := strconv.FormatInt(database.GetBannedUsersCount(), 10)
	md := mdparser.GetBold("📊 Current stats of ")
	nme := func() mdparser.WMarkDown {
		return md.AppendNormalThis("\n• ")
	}
	reasonAppend := func(c int64, r string) mdparser.WMarkDown {
		nme().AppendMonoThis(strconv.FormatInt(c, 10))
		return md.AppendNormalThis(" banned due to ").AppendMonoThis(r)
	}

	md.AppendHyperLinkThis("Sibyl System:", "http://t.me/SibylSystem")
	nme().AppendMonoThis(tbanned).AppendNormalThis(" banned users")
	reasonAppend(10, "TROLLING")
	reasonAppend(20, "SPAM")
	reasonAppend(30, "EVADE")
	reasonAppend(40, "CUSTOM")
	reasonAppend(50, "PSYCHOHAZARD")
	reasonAppend(60, "MALIMP")
	reasonAppend(70, "NSFW")
	reasonAppend(80, "RAID")
	reasonAppend(90, "MASSADD")
	nme().AppendMonoThis("300").AppendNormalThis(" users with Cloudy Psychopass")
	nme().AppendMonoThis("90").AppendNormalThis(" tokens generated")
	nme().AppendMonoThis("14").AppendNormalThis(" registered Inspectors")
	nme().AppendMonoThis("250").AppendNormalThis(" registered Enforcers")

	_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode:             sibylValues.MarkDownV2,
		DisableWebPagePreview: true,
	})
	return nil
}
