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
			data.Source,
			current.UserId,
			data.ReporterId,
			data.ReporterPermission,
			current.TargetType,
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

func GetMultiScan(associationBanId string) *sv.MultiScanRawData {
	associationScanMutex.Lock()
	data := associationScanMap[associationBanId]
	associationScanMutex.Unlock()
	if data != nil {
		return data
	}

	var scans []*sv.Report
	lockdb()
	SESSION.Model(modelScan).Where(
		"association_ban_id = ?", associationBanId,
	).Find(&scans)
	unlockdb()

	if len(scans) == 0 {
		return nil
	}

	data = toMultiScanRawData(scans)
	associationScanMutex.Lock()
	associationScanMap[associationBanId] = data
	associationScanMutex.Unlock()
	return data
}

func toMultiScanRawData(scans []*sv.Report) *sv.MultiScanRawData {
	data := &sv.MultiScanRawData{
		Status: scans[0].ScanStatus,
	}

	for _, current := range scans {
		data.Users = append(data.Users, sv.MultiScanUserInfo{
			UserId:     current.TargetUser,
			Reason:     current.ReportReason,
			Message:    current.ReportMessage,
			TargetType: current.TargetType,
		})

		if data.ReporterId == 0 && current.ReporterId != 0 {
			data.ReporterId = current.ReporterId
		}

		if data.ReporterPermission == 0 && current.ReporterPermission != 0 {
			data.ReporterPermission = current.ReporterPermission
		}

		if data.Source == "" && current.ScanSourceLink != "" {
			data.Source = current.ScanSourceLink
		}

	}

	data.SetCacheDate()

	return data
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

	if scan.IsApproved() {
		agent, err := GetTokenFromId(scan.ReporterId)
		if agent == nil || err != nil {
			return
		}
		agent.AcceptedReports++
		lockdb()
		tx := SESSION.Begin()
		tx.Save(agent)
		tx.Commit()
		unlockdb()
	} else if scan.IsRejected() {
		agent, err := GetTokenFromId(scan.ReporterId)
		if agent == nil || err != nil {
			return
		}
		agent.DeniedReports++
		lockdb()
		tx := SESSION.Begin()
		tx.Save(agent)
		tx.Commit()
		unlockdb()
	}
}
