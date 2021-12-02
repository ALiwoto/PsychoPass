package database

import (
	"time"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
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
	u.FormatBanDate()
	u.SetBanFlags()
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

	if lastStats != nil && !lastStats.IsExpired(sibylConfig.GetStatsCacheTime()) {
		return lastStats, nil
	}

	lastStats = new(sv.StatValue)
	lastStats.SetCachedTime()
	var tmp int64
	lockdb()

	// users related stats
	SESSION.Model(modelUser).Where("banned = ?", true).Count(&tmp)
	lastStats.BannedCount = tmp

	SESSION.Model(modelUser).Where("flag_trolling = ?", true).Count(&tmp)
	lastStats.TrollingBanCount = tmp

	SESSION.Model(modelUser).Where("flag_spam = ?", true).Count(&tmp)
	lastStats.SpamBanCount = tmp

	SESSION.Model(modelUser).Where("flag_evade = ?", true).Count(&tmp)
	lastStats.EvadeBanCount = tmp

	SESSION.Model(modelUser).Where("flag_custom = ?", true).Count(&tmp)
	lastStats.CustomBanCount = tmp

	SESSION.Model(modelUser).Where("flag_psycho_hazard = ?", true).Count(&tmp)
	lastStats.PsychoHazardBanCount = tmp

	SESSION.Model(modelUser).Where("flag_mal_imp = ?", true).Count(&tmp)
	lastStats.MalImpBanCount = tmp

	SESSION.Model(modelUser).Where("flag_nsfw = ?", true).Count(&tmp)
	lastStats.NSFWBanCount = tmp

	SESSION.Model(modelUser).Where("flag_raid = ?", true).Count(&tmp)
	lastStats.RaidBanCount = tmp

	SESSION.Model(modelUser).Where("flag_mass_add = ?", true).Count(&tmp)
	lastStats.MassAddBanCount = tmp

	SESSION.Model(modelUser).Where("flag_spam_bot = ?", true).Count(&tmp)
	lastStats.SpamBotBanCount = tmp

	SESSION.Model(modelUser).Where("crime_coefficient < ? AND crime_coefficient > ? AND banned = ?",
		sv.UpperCloudyFactor, sv.LowerCloudyFactor, false).Count(&tmp)
	lastStats.CloudyCount = tmp

	// token related stats:
	SESSION.Model(modelToken).Count(&tmp)
	lastStats.TokenCount = tmp

	SESSION.Model(modelToken).Where("permission = ?", sv.Inspector).Count(&tmp)
	lastStats.InspectorsCount = tmp

	SESSION.Model(modelToken).Where("permission = ?", sv.Enforcer).Count(&tmp)
	lastStats.EnforcesCount = tmp

	unlockdb()

	if SESSION.Error != nil {
		return nil, SESSION.Error
	}

	return lastStats, nil
}

func NewUser(u *sv.User) {
	u.FormatBanDate()
	u.SetBanFlags()
	lockdb()
	tx := SESSION.Begin()
	tx.Save(u)
	tx.Commit()
	unlockdb()
	u.SetCacheDate()
	userMapMutex.Lock()
	userDbMap[u.UserID] = u
	userMapMutex.Unlock()
}

// ForceInsert function acts like `AddBan`, but it doesn't ban the user.
// it will calculate an average crime coefficient by the passed-by permission.
func ForceInsert(userID int64, perm sv.UserPermission) *sv.User {
	user := &sv.User{
		UserID: userID,
		Banned: false,
		Date:   time.Now(),
	}
	user.IncreaseCrimeCoefficientByPerm(perm)
	NewUser(user)
	return user
}
