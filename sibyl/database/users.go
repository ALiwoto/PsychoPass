package database

import (
	"time"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

// GetUserFromId returns the user value from the database, returns error if any.
// please do notice that if the user can't be found in database, this function
// will return nil for both return variable. you should always have a checker for first
// value as well. This function use caching to speed up its operations, so most of the
// times it will return as fast as possible.
func GetUserFromId(id int64) (*sv.User, error) {
	if SESSION == nil {
		return nil, ErrNoSession
	}

	u := userDbMap.Get(id)
	if u == emptyUser {
		// an empty user is cached in our map, which means we have
		// already checked for this id, don't waste resources and don't
		// send any database queries again.
		// instead return nil.
		// this might use up a little bit memory, but instead will increase
		// the speed a lot.
		return nil, nil
	} else if u != nil {
		// map has returned a valid user, return it.
		return u, nil
	}

	// when we are at this point, it means we are checking this id for the first ever
	// time (in a while). we should send database query to get the user.
	u = &sv.User{}
	lockdb()
	SESSION.Where("user_id = ?", id).Take(u)
	unlockdb()
	if u.UserID != id {
		// if the user id is not the same as the one we are looking for,
		// which means the user doesn't exist in the database.
		// cache an empty user in our map, so we don't waste resources
		// again and again for checking the same id in future.
		userDbMap.Add(id, emptyUser)
		return nil, nil
	}

	// everything is fine, set parameters for the user, cache it in memory
	// and return the value.
	u.FormatBanDate()
	u.SetBanFlags()
	userDbMap.Add(u.UserID, u)

	return u, nil
}

// GetAllBannedUsers this function returns all banned users in the database.
// it doesn't use any cache at all, maybe we should implement it in future.
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

// GetBannedUsersCount returns all banned users count in the database.
// this function isn't using any caching, calling it all the times will result in
// slow operations. (the caller itself is supposed to cache the returned value)
func GetBannedUsersCount() (c int64) {
	lockdb()
	m := SESSION.Model(&sv.User{})
	m.Where("banned = ?", true).Count(&c)
	unlockdb()
	return
}

// FetchStat function will return a pointer to a sv.StatValue struct,
// representing the current stats of the database. this function is using caching
// for speeding up the operations it's using (the cache total time is obtained by calling
// sibylConfig.GetStatsCacheTime() function). calling this function repeatedly shouldn't cause
// any performance issues (in... normal situations I guess).
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

// NewUser saves the user to the database and caches it in memory.
func NewUser(u *sv.User) {
	u.FormatBanDate()
	u.SetBanFlags()

	if !u.ShouldSaveInDB() {
		u.CrimeCoefficient = 0
	}

	lockdb()
	tx := SESSION.Begin()
	tx.Save(u)
	tx.Commit()
	unlockdb()

	userDbMap.Add(u.UserID, u)
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
