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

func GetScan(uniqueId string) *sv.Report {
	if uniqueId == "" {
		return nil
	}

	scanMapMutex.Lock()
	scan := scanDbMap[uniqueId]
	scanMapMutex.Unlock()
	if scan != nil {
		return scan
	}

	lockdb()
	scan = &sv.Report{}
	SESSION.Model(modelScan).Where("unique_id = ?", uniqueId).Take(scan)
	unlockdb()

	if scan.UniqueId == uniqueId {
		return scan
	}

	scanMapMutex.Lock()
	scanDbMap[scan.UniqueId] = scan
	scanMapMutex.Unlock()
	return scan
}

func UpdateScan(scan *sv.Report) {
	if scan == nil || scan.IsInvalid() {
		return
	}

	lockdb()
	tx := SESSION.Begin()
	tx.Save(scan)
	tx.Commit()
	unlockdb()
}
