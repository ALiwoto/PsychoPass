/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package banHandlers

import (
	"sync"
	"time"

	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
)

// applyMultiBan will apply the multi ban request. this function should
// be called in different goroutine than our http request goroutine.
// it uses a mutex to control the huge payload on database.
func applyMultiBan(data *sv.MultiBanRawData, by int64) {
	if multiBanMutex == nil {
		multiBanMutex = new(sync.Mutex)
	}

	// calling this method here will cause all of the users of this
	// multi-ban get perma-ban if and only if one of the users have
	// got perma-ban.
	data.SyncPermaBans()

	multiBanMutex.Lock()

	var tmpToken *sv.Token
	var tmpUser *sv.User
	var err error

	for _, current := range data.Users {
		if current.IsInvalid(by) {
			continue
		}

		tmpToken, _ = database.GetTokenFromId(current.UserId)
		if tmpToken != nil && !tmpToken.CanBeBanned() {
			continue
		}

		tmpUser, err = database.GetUserFromId(current.UserId)
		var count int
		if tmpUser != nil && err == nil {
			if tmpUser.Banned {
				cloneUser := tmpUser.Clone()
				if current.TargetType != tmpUser.TargetType {
					// check both conditions; if they don't match, update the field.
					cloneUser.TargetType = current.TargetType
				}
				cloneUser.BannedBy = by
				cloneUser.Message = current.Message
				cloneUser.Date = time.Now()
				cloneUser.BanSourceUrl = current.Source
				cloneUser.SetAsBanReason(current.Reason)
				cloneUser.IncreaseCrimeCoefficientAuto()
				if cloneUser.CrimeCoefficient < tmpUser.CrimeCoefficient {
					continue
				}

				*tmpUser = *cloneUser
				database.UpdateBanParameter(tmpUser, data.IsSilent)
				continue
			}
			count = tmpUser.BanCount
		}

		_ = database.AddBanByInfo(&current, by, count, data.IsSilent)
	}

	multiBanMutex.Unlock()
}

// applyMultiUnBan will apply the multi unban request. this function should
// be called in different goroutine than our http request goroutine.
// it uses a mutex to control the huge payload on database.
func applyMultiUnBan(data *sv.MultiUnBanRawData, by int64) {
	if multiUnBanMutex == nil {
		multiUnBanMutex = new(sync.Mutex)
	}

	var tmpUser *sv.User
	var err error

	multiUnBanMutex.Lock()

	for _, current := range data.Users {
		tmpUser, err = database.GetUserFromId(current)
		if tmpUser == nil || err != nil {
			// user not found or there is an issue in database package.
			continue
		}

		if !tmpUser.Banned && len(tmpUser.Reason) == 0 && len(tmpUser.BanFlags) == 0 {
			// user is not banned at all.
			continue
		}

		database.RemoveUserBan(tmpUser, false)
	}

	multiUnBanMutex.Unlock()
}
