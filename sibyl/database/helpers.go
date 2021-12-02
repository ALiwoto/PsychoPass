package database

import (
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/hashing"
)

func lockdb() {
	if sibylConfig.SibylConfig.UseSqlite {
		dbMutex.Lock()
	}
}

func unlockdb() {
	if sibylConfig.SibylConfig.UseSqlite {
		dbMutex.Unlock()
	}
}

func migrateOwners(owners []int64) {
	var tmpToken *sv.Token
	var err error

	for _, current := range owners {
		tmpToken, err = GetTokenFromId(current)
		if tmpToken == nil || err != nil {
			h := hashing.GetUserToken(current)
			tmpToken = &sv.Token{
				Permission: sv.Owner,
				UserId:     current,
				Hash:       h,
			}

			// insert the new token into database uwu
			NewToken(tmpToken)
			continue
		}

		if tmpToken.Permission < sv.Owner {
			UpdateTokenPermission(tmpToken, sv.Owner)
		}
	}
}
