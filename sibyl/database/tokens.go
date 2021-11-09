package database

import (
	"time"

	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/hashing"
)

func GetTokenFromId(id int64) (*sv.Token, error) {
	if SESSION == nil {
		return nil, ErrNoSession
	}

	tokenMapMutex.Lock()
	t := tokenDbMap[id]
	tokenMapMutex.Unlock()
	if t != nil {
		t.SetCacheDate()
		return t, nil
	}

	p := &sv.Token{}
	lockdb()
	SESSION.Where("user_id = ?", id).Take(p)
	unlockdb()
	if len(p.Hash) == 0 || p.UserId == 0 || p.UserId != id {
		// not found
		return nil, nil
	}
	p.SetCacheDate()
	tokenMapMutex.Lock()
	tokenDbMap[p.UserId] = p
	tokenMapMutex.Unlock()

	return p, nil
}

func GetTokenFromString(token string) (*sv.Token, error) {
	id := hashing.GetIdFromToken(token)
	if id == 0 {
		return nil, ErrInvalidToken
	}

	u, err := GetTokenFromId(id)
	if err != nil {
		return nil, err
	}

	if u == nil || u.Hash != token {
		return nil, ErrInvalidToken
	}

	return u, nil
}

func GetAllRegistered(includeOwners bool) ([]int64, error) {
	if SESSION == nil {
		return nil, ErrNoSession
	}

	var tokens []sv.Token
	lockdb()
	if includeOwners {
		SESSION.Model(modelToken).Where("permission > ?", sv.NormalUser).Find(&tokens)
	} else {
		SESSION.Model(modelToken).Where("permission > ? AND NOT permission = ?",
			sv.NormalUser, sv.Owner).Find(&tokens)
	}
	unlockdb()

	return convertToIntArray(tokens), nil
}

func UpdateTokenLastUsageById(id int64) {
	u, err := GetTokenFromId(id)
	if err != nil || u == nil {
		return
	}

	u.LastUsage = time.Now()
	lockdb()
	tx := SESSION.Begin()
	tx.Save(u)
	tx.Commit()
	unlockdb()
}

func UpdateTokenLastUsage(t *sv.Token) {
	t.LastUsage = time.Now()
	lockdb()
	tx := SESSION.Begin()
	tx.Save(t)
	tx.Commit()
	unlockdb()
}

func UpdateTokenPermission(t *sv.Token, perm sv.UserPermission) {
	t.Permission = perm
	lockdb()
	tx := SESSION.Begin()
	tx.Save(t)
	tx.Commit()
	unlockdb()
}

func RevokeTokenHash(t *sv.Token, hash string) {
	t.Hash = hash
	lockdb()
	tx := SESSION.Begin()
	tx.Save(t)
	tx.Commit()
	unlockdb()
}

func NewToken(t *sv.Token) {
	lockdb()
	tx := SESSION.Begin()
	tx.Save(t)
	tx.Commit()
	unlockdb()
}

func convertToIntArray(tokens []sv.Token) []int64 {
	if len(tokens) == 0 {
		return nil
	}
	var ids []int64
	for _, t := range tokens {
		ids = append(ids, t.UserId)
	}
	return ids
}
