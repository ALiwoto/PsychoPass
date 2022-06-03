/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package devPlugin

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	gitpullCmd := handlers.NewCommand(GitpullCmd, gitpullHandler)
	restartCmd := handlers.NewCommand(RestartCmd, gitpullHandler)
	gitpullCmd.Triggers = t
	restartCmd.Triggers = t
	d.AddHandler(gitpullCmd)
	d.AddHandler(restartCmd)
}
