package sibylValues

import (
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var (
	HelperBot         *gotgbot.Bot
	BotUpdater        *ext.Updater
	SendReportHandler ReportHandler
)

var (
	reportMutex *sync.Mutex
	// key: report unique id; value: reporter.
	reportUniqueMap map[int64]*Report
	// the current advance of unique id in reporting.
	// please notice that default value of this variable
	// SHOULD be 0x1. please don't change it.
	currentUniqueId int64 = 0x1
)
