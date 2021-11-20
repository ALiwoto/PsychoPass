package tokenPlugin

import (
	"strconv"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"
)

func (a *AssignValue) ParseToMd(info mdparser.WMarkDown) mdparser.WMarkDown {
	by := strconv.FormatInt(a.agent.UserId, 10)
	md := mdparser.GetNormal("\u200D#Assignment request\n")
	if a.agentProfile != nil {
		name := utils.GetNameFromUser(a.agentProfile, a.agent.GetStringPermission())
		md.AppendBoldThis(" • By: ").AppendMentionThis(name, a.agentProfile.Id)
	} else {
		md.AppendBoldThis(" • By: ").AppendMonoThis(by)
	}
	md.ElThis().AppendThis(info)
	return md
}
