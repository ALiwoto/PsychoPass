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
				if current.IsBot != tmpUser.IsBot {
					// check both conditions; if they don't match, update the field.
					tmpUser.IsBot = current.IsBot
				}
				tmpUser.BannedBy = by
				tmpUser.Message = current.Message
				tmpUser.Date = time.Now()
				tmpUser.BanSourceUrl = current.Source
				tmpUser.SetAsBanReason(current.Reason)
				tmpUser.IncreaseCrimeCoefficientAuto()
				database.UpdateBanparameter(tmpUser)
				continue
			}
			count = tmpUser.BanCount
		}

		_ = database.AddBanByInfo(&current, by, count)
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
