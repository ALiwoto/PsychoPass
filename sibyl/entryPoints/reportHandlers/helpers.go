package reportHandlers

import (
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
)

// applyMultiScan will apply multi-scan reqest. make sure that sv.SendMultiReportHandler is not nil
// before running this function.
// this function should be run in a different goroutine rather than http handler's
// goroutine.
func applyMultiScan(data *sv.MultiScanRawData) {
	data.GenerateID()
	database.AddMultiScan(data)
	sv.SendMultiReportHandler(data)
}
