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

func AddMultiScan(data *sv.MultiScanRawData) {
	if data == nil {
		return
	}

	var scans []*sv.Report
	var tmpScans *sv.Report
	for _, current := range data.Users {
		tmpScans = sv.NewReport(
			current.Reason,
			current.Message,
			current.Source,
			current.UserId,
			data.ReporterId,
			data.ReporterPermission,
			current.IsBot,
		)
		scans = append(scans, tmpScans)
	}

	lockdb()
	tx := SESSION.Begin()
	tx.Create(scans)
	tx.Commit()
	unlockdb()
	data.SetCacheDate()
	associationScanMutex.Lock()
	associationScanMap[data.AssociationBanId] = data
	associationScanMutex.Unlock()
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
