package startPlugin

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	startCmd := handlers.NewCommand(StartCmd, startHandler)
	createCmd := handlers.NewCommand(CreateCmd, startHandler)
	newCmd := handlers.NewCommand(NewCmd, startHandler)
	startCmd.Triggers = t
	createCmd.Triggers = t
	newCmd.Triggers = t
	d.AddHandler(startCmd)
	d.AddHandler(createCmd)
	d.AddHandler(newCmd)
}
