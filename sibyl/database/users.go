package database

import (
	"errors"

	sv "github.com/AnimeKaizoku/sibylapi-go/sibyl/core/sibylValues"
)

func GetUserFromId(id int64) (*sv.User, error) {
	if SESSION == nil {
		return nil, errors.New("failed to Get token data as Session is nil")
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

func NewUser(u *sv.User) {
	lockdb()
	tx := SESSION.Begin()
	tx.Save(u)
	tx.Commit()
	unlockdb()
}
