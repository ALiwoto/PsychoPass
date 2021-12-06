package database

import (
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

func AddScan(scan *sv.Report) {
	if scan == nil || scan.IsInvalid() {
		return
	}

	lockdb()
	tx := SESSION.Begin()
	tx.Save(scan)
	tx.Commit()
	unlockdb()
	scan.SetCacheDate()
	scanMapMutex.Lock()
	scanDbMap[scan.UniqueId] = scan
	scanMapMutex.Unlock()
}
