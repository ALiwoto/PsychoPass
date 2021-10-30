package sibylValues

import (
	"errors"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var (
	HelperBot         *gotgbot.Bot
	BotUpdater        *ext.Updater
	SendReportHandler ReportHandler
)

var (
	ErrInvalidPerm = errors.New("invalid permission provided")
)
