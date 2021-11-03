package infoPlugin

import (
	"strconv"

	"github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/shellUtils"
)

func fetchGitStats(md mdparser.WMarkDown) {
	rawGit := shellUtils.GetGitStats()
	if len(rawGit) == 0 {
		return
	}
	allRaws := strongStringGo.Split(rawGit, "\n")
	if len(allRaws) < 3 {
		return
	}
	shortGit := allRaws[0]
	longGit := allRaws[1]
	gitVs := strongStringGo.Split(allRaws[2], " ", "\t")
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

	md.AppendNormalThis("ℹ️ ").AppendHyperLinkThis("Git ", gitBaseUrl)
	md.AppendBoldThis("Status:")
	md.AppendNormalThis("\n• Currently on: ").AppendHyperLinkThis(shortGit, commitUrl)
	md.AppendNormalThis("\n• Running behind by: ").AppendMonoThis(strconv.Itoa(vsInt))
	md.AppendNormalThis(" commits\n\n")
}
