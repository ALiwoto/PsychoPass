/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package infoPlugin

import (
	"runtime"
	"strconv"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func StatsHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	t, err := database.GetTokenFromId(user.Id)
	if err != nil || t == nil || !t.CanGetStats() {
		return ext.EndGroups
	}

	md := mdparser.GetEmpty()

	fetchGitStats(md)

	md.Bold("📊 Current census of ")
	nme := func() mdparser.WMarkDown {
		return md.Normal("\n• ")
	}
	stat, err := database.FetchStat()
	if err != nil {
		md = mdparser.GetItalic("There was a problem when fetching stats from database.")
		logging.UnexpectedError(err)
		return ext.EndGroups
	}

	reasonAppend := func(c int64, r string) mdparser.WMarkDown {
		nme().Mono(strconv.FormatInt(c, 10))
		return md.Normal(" banned for ").Mono(r)
	}

	md.Link("Sibyl System:", "https://t.me/SibylSystem/13")
	nme().Normal("Total bans: ")
	md.Mono(stat.GetBannedCountString())

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

	nme().Mono(stat.GetCloudyCountString())
	md.Normal(" with Cloudy Psychopass")

	nme().Mono(stat.GetTokenCountString())
	md.Normal(" tokens generated")

	nme().Mono(stat.GetInspectorsCountString())
	md.Normal(" registered Inspectors")

	nme().Mono(stat.GetEnforcesCountString())
	md.Normal(" registered Enforcers")

	md.Normal("\n\n• Server uptime: ")
	md.Mono(sibylValues.GetPrettyUptime())
	md.Normal("\n• Version: ")
	md.Mono(runtime.Version())
	md.Normal("\n• Cgo calls: ")
	md.Mono(ssg.ToBase10(runtime.NumCgoCall()))
	md.Normal("\n• Goroutines: ")
	md.Mono(ssg.ToBase10(int64(runtime.NumGoroutine())))

	_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode:             sibylValues.MarkDownV2,
		DisableWebPagePreview: true,
	})
	return nil
}
