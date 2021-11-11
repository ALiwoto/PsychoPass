package banHandlers

import (
	"sync"
	"time"

	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
)

// applyMultiBan will apply the multi ban request. this function should
// be called in different goroutine than our http request goroutine.
// it uses a mutex to control the huge payload on database.
func applyMultiBan(data *sv.MultiBanRawData, by int64) {
	if data == nil || len(data.Users) == 0 {
		return
	}

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
