package tokenPlugin

import (
	"strconv"

	"github.com/ALiwoto/mdparser/mdparser"
)

func (a *AssignValue) ParseToMd(info mdparser.WMarkDown) mdparser.WMarkDown {
	by := strconv.FormatInt(a.agent.UserId, 10)
	md := mdparser.GetNormal("\u200D#Assignment request\n")
	md.AppendBoldThis(" ・ By: ").AppendMonoThis(by).ElThis()
	md.AppendThis(info)
	return md
}
