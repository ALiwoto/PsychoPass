package devPlugin

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	statsCmd := handlers.NewCommand(GitpullCmd, gitpullHandler)
	statsCmd.Triggers = t
	d.AddHandler(statsCmd)
}
