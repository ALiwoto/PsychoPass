package database

import (
	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
)

func GetUserFromId(id int64) (*sv.User, error) {
	if SESSION == nil {
		return nil, ErrNoSession
	}

	userMapMutex.Lock()
	u := userDbMap[id]
	userMapMutex.Unlock()
	if u != nil {
		u.SetCacheDate()
		return u, nil
	}

	u = &sv.User{}
	lockdb()
	SESSION.Where("user_id = ?", id).Take(u)
	unlockdb()
	if u.UserID != id {
		// not found
		return nil, nil
	}

	u.SetCacheDate()
	userMapMutex.Lock()
	userDbMap[u.UserID] = u
	userMapMutex.Unlock()

	return u, nil
}

func GetAllBannedUsers() ([]sv.User, error) {
	if SESSION == nil {
		return nil, ErrNoSession
	}

	var users []sv.User
	lockdb()
	SESSION.Where("banned = ?", true).Find(&users)
	unlockdb()

	return users, nil
}

func GetBannedUsersCount() (c int64) {
	lockdb()
	m := SESSION.Model(&sv.User{})
	m.Where("banned = ?", true).Count(&c)
	unlockdb()
	return
}

func FetchStat() (*sv.StatValue, error) {
	if SESSION == nil {
		return nil, ErrNoSession
	}

	var stat = new(sv.StatValue)
	var tmp int64
	lockdb()

	// users related stats
	m := SESSION.Model(&sv.User{})

	m.Where("banned = ?", true).Count(&tmp)
	stat.BannedCount = tmp

	m.Where("flag_trolling = ?", true).Count(&tmp)
	stat.TrollingBanCount = tmp

	m.Where("flag_spam = ?", true).Count(&tmp)
	stat.SpamBanCount = tmp

	m.Where("flag_evade = ?", true).Count(&tmp)
	stat.EvadeBanCount = tmp

	m.Where("flag_custom = ?", true).Count(&tmp)
	stat.CustomBanCount = tmp

	m.Where("flag_psycho_hazard = ?", true).Count(&tmp)
	stat.PsychoHazardBanCount = tmp

	m.Where("flag_mal_imp = ?", true).Count(&tmp)
	stat.MalImpBanCount = tmp

	m.Where("flag_nsfw = ?", true).Count(&tmp)
	stat.NSFWBanCount = tmp

	m.Where("flag_raid = ?", true).Count(&tmp)
	stat.RaidBanCount = tmp

	m.Where("flag_mass_add = ?", true).Count(&tmp)
	stat.MassAddBanCount = tmp

	m.Where("crime_coefficient < ? AND crime_coefficient > ?",
		sv.UpperCloudyFactor, sv.LowerCloudyFactor).Count(&tmp)
	stat.CloudyCount = tmp

	// token related stats:

	m = SESSION.Model(&sv.Token{})
	m.Count(&tmp)
	stat.TokenCount = tmp

	m.Where("permission = ?", sv.Inspector).Count(&tmp)
	stat.InspectorsCount = tmp

	m.Where("permission = ?", sv.Enforcer).Count(&tmp)
	stat.EnforcesCount = tmp

	unlockdb()

	if SESSION.Error != nil {
		return nil, SESSION.Error
	} else if m.Error != nil {
		return nil, m.Error
	}

	return stat, nil
}

func NewUser(u *sv.User) {
	lockdb()
	tx := SESSION.Begin()
	tx.Save(u)
	tx.Commit()
	unlockdb()
}
