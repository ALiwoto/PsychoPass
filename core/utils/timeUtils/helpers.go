package timeUtils

import (
	"time"

	"gitlab.com/Dank-del/SibylAPI-Go/core/utils/stringUtils"
)

// format of the date time will be dd/MM/yyyy HH:mm:ss
func GenerateCurrentDateTime() string {
	t := time.Now()

	str := stringUtils.MakeSureNum(t.Day(), 2) + "/"
	str += stringUtils.MakeSureNum(int(t.Month()), 2) + "/"
	str += stringUtils.MakeSureNum(t.Year(), 4) + " "
	str += stringUtils.MakeSureNum(t.Hour(), 2) + ":"
	str += stringUtils.MakeSureNum(t.Minute(), 2) + ":"
	str += stringUtils.MakeSureNum(t.Second(), 2)

	return str
}
