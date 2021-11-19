package tokenPlugin

import "github.com/ALiwoto/mdparser/mdparser"

func (a *AssignValue) ParseToMd() mdparser.WMarkDown {
	md := mdparser.GetNormal("\u200D#Assignment request")
	return md
}
