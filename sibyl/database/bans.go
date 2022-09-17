/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package database

import (
	"time"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylLogging"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

func AddBan(info *BanInfo) *sv.User {
	user := &sv.User{
		UserID:       info.UserID,
		Banned:       true,
		Date:         time.Now(),
		Message:      info.Message,
		BannedBy:     info.Adder,
		BanSourceUrl: info.Src,
		SourceGroup:  info.SrcGroup,
		TargetType:   info.TargetType,
		BanCount:     info.Count,
	}

	user.SetAsBanReason(info.Reason)
	user.IncreaseCrimeCoefficientAuto()
	NewUser(user)

	if sibylLogging.SendToADHandler != nil && !info.IsSilent {
		go sibylLogging.SendToADHandler(user.ToDominatorData(true))
	}

	return user
}

func AddBanByInfo(info *sv.MultiBanUserInfo, adder int64,
	count int, silent bool) *sv.User {
	return AddBan(
		&BanInfo{
			UserID:     info.UserId,
			Adder:      adder,
			Reason:     info.Reason,
			SrcGroup:   info.SourceGroup,
			Message:    info.Message,
			Src:        info.Source,
			TargetType: info.TargetType,
			Count:      count,
			IsSilent:   silent,
		},
	)
}

// DeleteUser will delete a user from the sibyl database.
// WARNING: this function will NOT check for user existence inside of db,
// it won't error out or panic if the user id is not found in the db, but it's
// still waste of resources, do check the user before using this function.
func DeleteUser(userID int64) {
	userDbMap.Delete(userID)
	lockdb()
	SESSION.Delete(&sv.User{}, "user_id = ?", userID)
	unlockdb()
}

// RemoveUserBan will unban a user from the sibyl database.
func RemoveUserBan(user *sv.User, clearHistory bool) {
	if !user.Banned {
		// don't send any query to database if user is not banned.
		return
	}
	user.FormatBanDate()
	user.SetAsRestored(clearHistory)
	lockdb()
	tx := SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockdb()

	if sibylLogging.SendToADHandler != nil {
		go sibylLogging.SendToADHandler(user.ToDominatorData(false))
	}

}

// ClearHistory will unban a user from the sibyl database.
func ClearHistory(user *sv.User) {
	if user.Banned {
		return
	}
	user.ClearHistory()
	lockdb()
	tx := SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockdb()
}

// UpdateBanParameter will update a user's ban parameter into the database.
func UpdateBanParameter(user *sv.User, silent bool) {
	lockdb()
	tx := SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockdb()

	if sibylLogging.SendToADHandler != nil && !silent {
		go sibylLogging.SendToADHandler(user.ToDominatorData(true))
	}
}

func UpdateUserCrimeCoefficientByPerm(user *sv.User, perm sv.UserPermission) {
	pre := user.CrimeCoefficient
	user.IncreaseCrimeCoefficientByPerm(perm)
	if pre == user.CrimeCoefficient {
		// don't send any query to database if user's crime coefficient
		// is not changed.
		return
	}

	lockdb()
	tx := SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockdb()
}
