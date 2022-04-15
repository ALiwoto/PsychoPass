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
	if sv.SendMultiReportHandler == nil {
		// normally, impossible to reach here
		// this condition is added just in case
		return
	}

	// check and remove repeated user-ids sent by the client.
	// this checker is there to solve the problem of the client sending
	// multiple-repeated user-ids.
	// See also: https://github.com/MinistryOfWelfare/PsychoPass/issues/2
	data.Users = removeRepeatedUsers(data.Users)

	data.GenerateID()
	database.AddMultiScan(data)
	sv.SendMultiReportHandler(data)
}

// removeRepeatedUsers will remove repeated user-ids from the given slice.
func removeRepeatedUsers(users []sv.MultiScanUserInfo) []sv.MultiScanUserInfo {
	var result []sv.MultiScanUserInfo
	var shouldIgnore bool

	for _, currentUser := range users {
		if len(result) == 0 {
			result = append(result, currentUser)
			continue
		}

		for _, rUser := range result {
			if rUser.UserId == currentUser.UserId {
				shouldIgnore = true
				break
			}
		}

		if shouldIgnore {
			shouldIgnore = false
			continue
		}

		result = append(result, currentUser)
	}

	return result
}
