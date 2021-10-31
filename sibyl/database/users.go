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
	SESSION.Model(&sv.User{}).Where("banned = ?", true).Count(&c)
	unlockdb()
	return
}

func NewUser(u *sv.User) {
	lockdb()
	tx := SESSION.Begin()
	tx.Save(u)
	tx.Commit()
	unlockdb()
}
