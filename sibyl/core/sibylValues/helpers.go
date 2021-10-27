package sibylValues

import (
	"sync"
	"time"

	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/timeUtils"
)

func SetReportHander(h ReportHandler) {
	if h == nil || SendReportHandler != nil {
		return
	}

	SendReportHandler = h
	reportMutex = &sync.Mutex{}
	reportUniqueMap = make(map[int64]*Report)

	go reportMapCleaner()
}

func GetReportFromUniqueId(uniqueId, reporterId int64) *Report {
	if uniqueId < currentUniqueId {
		return nil
	}

	reportMutex.Lock()
	r := reportUniqueMap[uniqueId]
	reportMutex.Unlock()
	if r != nil && r.ReporterId != reporterId {
		return nil
	}

	return r
}

func NewReport(reason, message string, target, reporter int64,
	reporterPerm UserPermission) *Report {

	return &Report{
		state:              reportStateWaiting,
		ReportReason:       reason,
		ReportMessage:      message,
		TargetUser:         target,
		ReporterId:         reporter,
		ReportDate:         timeUtils.GenerateCurrentDateTime(),
		ReporterPermission: reporterPerm.GetStringPermission(),
	}
}

func reportMapCleaner() {
	for {
		time.Sleep(MaxReportCacheTime)
		if len(reportUniqueMap) == 0 || reportMutex == nil {
			// something unexpected had happened?
			// or maybe this is intended?
			// in each case, this means we have nothing to clean,
			// simply return from the function.
			return
		}

		reportMutex.Lock()
		for key, value := range reportUniqueMap {
			// if this key contains a nil value, remove the key from
			// the map. (this condition is unlikely to happen, but
			// wrote it just in case).
			if value == nil {
				delete(reportUniqueMap, key)
				continue
			}

			// if this report is too old, remove it from the memory.
			// if someone click on a report and mark it as accepted,
			// the reporter won't recieve anything.
			if time.Since(value.date) > MaxReportCacheTime {
				delete(reportUniqueMap, key)
				continue
			}

			// if this report is not waiting (either accepted or closed),
			// remove it from the memory.
			if value.state != reportStateWaiting {
				delete(reportUniqueMap, key)
				continue
			}
		}
		reportMutex.Unlock()
	}
}
