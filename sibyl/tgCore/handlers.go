package tgCore

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/tgCore/plugins/infoPlugin"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/tgCore/plugins/reportPlugin"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/tgCore/plugins/tokenPlugin"
)

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	infoPlugin.LoadAllHandlers(d, triggers)
	reportPlugin.LoadAllHandlers(d, triggers)
	tokenPlugin.LoadAllHandlers(d, triggers)
}
