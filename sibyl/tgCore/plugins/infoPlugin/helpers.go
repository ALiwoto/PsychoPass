/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package infoPlugin

import (
	"strconv"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/shellUtils"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	statsCmd := handlers.NewCommand(StatsCmd, StatsHandler)
	statsCmd.Triggers = t
	d.AddHandler(statsCmd)
}

func fetchGitStats(md mdparser.WMarkDown) {
	rawGit := shellUtils.GetGitStats()
	if len(rawGit) == 0 {
		// try again; in some situations, when we recently have
		// pushed to HEAD, the git command may not be able to
		// find the HEAD commit.
		rawGit = shellUtils.GetGitStats()
		if len(rawGit) == 0 {
			// give up and return :(
			return
		}
	}
	allRaws := ssg.Split(rawGit, "\n")
	if len(allRaws) < 3 {
		return
	}
	shortGit := allRaws[0]
	longGit := allRaws[1]
	gitVs := ssg.Split(allRaws[2], " ", "\t")
	if len(gitVs) != 2 {
		return
	}
	upstream, err := strconv.Atoi(gitVs[0])
	if err != nil {
		return
	}
	local, err := strconv.Atoi(gitVs[1])
	if err != nil {
		return
	}
	vsInt := upstream - local
	commitUrl := gitBaseUrl + "/commit/" + longGit

	md.Normal("ℹ️ ").Link("Git ", gitBaseUrl)
	md.Bold("Status:")
	md.Normal("\n• Current commit: ").Link(shortGit, commitUrl)
	md.Normal("\n• Running behind by ").Mono(strconv.Itoa(vsInt))
	md.Normal(" commits\n\n")
}
