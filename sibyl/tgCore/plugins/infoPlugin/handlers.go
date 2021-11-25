package infoPlugin

import (
	"strconv"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
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
	if err != nil || t == nil || !t.CanGetStats() {
		return ext.EndGroups
	}

	md := mdparser.GetEmpty()

	fetchGitStats(md)

	md.AppendBoldThis("📊 Current census of ")
	nme := func() mdparser.WMarkDown {
		return md.AppendNormalThis("\n• ")
	}
	stat, err := database.FetchStat()
	if err != nil {
		md = mdparser.GetItalic("There was a problem when fetching stats from database.")
		logging.UnexpectedError(err)
		return ext.EndGroups
	}

	reasonAppend := func(c int64, r string) mdparser.WMarkDown {
		nme().AppendMonoThis(strconv.FormatInt(c, 10))
		return md.AppendNormalThis(" banned due to ").AppendMonoThis(r)
	}

	md.AppendHyperLinkThis("Sibyl System:", "https://t.me/SibylSystem/13")
	nme().AppendNormalThis("Total ban count: ")
	md.AppendMonoThis(stat.GetBannedCountString())

	reasonAppend(stat.TrollingBanCount, "TROLLING")
	reasonAppend(stat.SpamBanCount, "SPAM")
	reasonAppend(stat.EvadeBanCount, "EVADE")
	reasonAppend(stat.CustomBanCount, "CUSTOM")
	reasonAppend(stat.PsychoHazardBanCount, "PSYCHOHAZARD")
	reasonAppend(stat.MalImpBanCount, "MALIMP")
	reasonAppend(stat.NSFWBanCount, "NSFW")
	reasonAppend(stat.RaidBanCount, "RAID")
	reasonAppend(stat.SpamBotBanCount, "SPAMBOT")
	reasonAppend(stat.MassAddBanCount, "MASSADD")

	nme().AppendMonoThis(stat.GetCloudyCountString())
	md.AppendNormalThis(" with Cloudy Psychopass")

	nme().AppendMonoThis(stat.GetTokenCountString())
	md.AppendNormalThis(" tokens generated")

	nme().AppendMonoThis(stat.GetInspectorsCountString())
	md.AppendNormalThis(" registered Inspectors")

	nme().AppendMonoThis(stat.GetEnforcesCountString())
	md.AppendNormalThis(" registered Enforcers")

	md.AppendNormalThis("\n\n- Server uptime: ")
	md.AppendMonoThis(sibylValues.GetPrettyUptime())
	_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode:             sibylValues.MarkDownV2,
		DisableWebPagePreview: true,
	})
	return nil
}
