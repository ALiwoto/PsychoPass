/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
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
	scanDbMap.Add(scan.UniqueId, scan)
}

func AddMultiScan(data *sv.MultiScanRawData) {
	if data == nil {
		return
	}

	var scans []*sv.Report
	var tmpScan *sv.Report
	for _, current := range data.Users {
		tmpScan = sv.NewReport(
			current.Reason,
			current.Message,
			data.Source,
			current.UserId,
			data.ReporterId,
			data.ReporterPermission,
			current.TargetType,
		)

		// sv.NewReport function doesn't set `AssociationBanId` field
		// of the tmpScan variable. we have to set them manually here.
		// for more info please visit: https://github.com/MinistryOfWelfare/PsychoPass/issues/8
		tmpScan.AssociationBanId = data.AssociationBanId
		scans = append(scans, tmpScan)
	}

	// without doing this, none of the `Approve`, `Reject` and `Close` will work.
	// See also: https://github.com/MinistryOfWelfare/PsychoPass/issues/3
	// See also: https://github.com/MinistryOfWelfare/PsychoPass/issues/3#issuecomment-1098486411
	data.Origins = scans

	lockdb()
	tx := SESSION.Begin()
	tx.Create(scans)
	tx.Commit()
	unlockdb()
	associationScanMap.Add(data.AssociationBanId, data)
}

func GetMultiScan(associationBanId string) *sv.MultiScanRawData {
	data := associationScanMap.Get(associationBanId)
	if data == emptyAssociationData {
		// we have already done a check for this associationBan Id and we do know
		// that it doesn't exist, don't send a new query to database and return nil.
		return nil
	}

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
		// id is not found, cache the id in memory so we don't send new database
		// query in the future for this id.
		associationScanMap.Add(associationBanId, emptyAssociationData)
		return nil
	}

	data = toMultiScanRawData(scans)
	associationScanMap.Add(associationBanId, data)
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

		data.Origins = append(data.Origins, current)
	}

	return data
}

func GetScan(uniqueId string) *sv.Report {
	if uniqueId == "" {
		return nil
	}

	scan := scanDbMap.Get(uniqueId)
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

	scanDbMap.Add(uniqueId, scan)
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

func UpdateMultipleScan(data *sv.MultiScanRawData) {
	if data == nil || len(data.Origins) == 0 {
		return
	}

	for _, current := range data.Origins {
		UpdateScan(current)
	}
}
