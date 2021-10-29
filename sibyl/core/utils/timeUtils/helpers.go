package timeUtils

import (
	"time"

	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/utils/stringUtils"
)

// GenerateCurrentDateTime format of the date time will be dd/MM/yyyy HH:mm:ss
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

// GenerateSuitableDateTime will format of the date time to dd-MM-yyyy HH-mm-ss
func GenerateSuitableDateTime() string {
	t := time.Now()

	str := stringUtils.MakeSureNum(t.Day(), 2) + "-"
	str += stringUtils.MakeSureNum(int(t.Month()), 2) + "-"
	str += stringUtils.MakeSureNum(t.Year(), 4) + "--"
	str += stringUtils.MakeSureNum(t.Hour(), 2) + "-"
	str += stringUtils.MakeSureNum(t.Minute(), 2) + "-"
	str += stringUtils.MakeSureNum(t.Second(), 2)

	return str
}
