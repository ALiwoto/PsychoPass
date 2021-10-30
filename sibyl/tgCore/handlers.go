package tgCore

import (
	"github.com/AnimeKaizoku/PsychoPass/sibyl/tgCore/plugins/infoPlugin"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/tgCore/plugins/reportPlugin"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/tgCore/plugins/tokenPlugin"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	infoPlugin.LoadAllHandlers(d, triggers)
	reportPlugin.LoadAllHandlers(d, triggers)
	tokenPlugin.LoadAllHandlers(d, triggers)
}
