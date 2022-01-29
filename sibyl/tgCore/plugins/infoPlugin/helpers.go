package infoPlugin

import (
	"strconv"

	"github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/shellUtils"
)

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

	md.Normal("ℹ️ ").Link("Git ", gitBaseUrl)
	md.Bold("Status:")
	md.Normal("\n• Currently on: ").Link(shortGit, commitUrl)
	md.Normal("\n• Running behind by: ").Mono(strconv.Itoa(vsInt))
	md.Normal(" commits\n\n")
}
