package database

import (
	"errors"
	"time"

	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/hashing"
)

func GetTokenFromId(id int64) (*sv.Token, error) {
	if SESSION == nil {
		return nil, errors.New("failed to Get token data as Session is nil")
	}

	p := sv.Token{}
	SESSION.Where("id = ?", id).Take(&p)
	if len(p.Hash) == 0 || p.Id == 0 || p.Id != id {
		// not found
		return nil, nil
	}

	return &p, nil
}

func GetTokenFromString(token string) (*sv.Token, error) {
	id := hashing.GetIdFromToken(token)
	if id == 0 {
		return nil, errors.New("token is invalid")
	}
	u, err := GetTokenFromId(id)
	if err != nil {
		return nil, err
	}
	if u == nil || u.Hash != token {
		return nil, errors.New("token is invalid")
	}

	return u, nil
}

func UpdateTokenLastUsageById(id int64) {
	u, err := GetTokenFromId(id)
	if err != nil || u == nil {
		return
	}

	u.LastUsage = time.Now()
	tx := SESSION.Begin()
	tx.Save(u)
	tx.Commit()
}

func UpdateTokenLastUsage(t *sv.Token) {
	t.LastUsage = time.Now()
	tx := SESSION.Begin()
	tx.Save(t)
	tx.Commit()
}

func RevokeTokenHash(t *sv.Token, hash string) {
	t.Hash = hash
	tx := SESSION.Begin()
	tx.Save(t)
	tx.Commit()
}

func NewToken(t *sv.Token) {
	tx := SESSION.Begin()
	tx.Save(t)
	tx.Commit()
}
