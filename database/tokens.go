package database

import (
	"errors"
	"gitlab.com/Dank-del/SibylAPI-Go/server"
)

type Token struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Hash       string `json:"hash"`
	UserID     int64  `json:"user_id"`
	Permission string `json:"permission"`
}

func GetFromToken(token string) (*Token, error) {
	if SESSION == nil {
		return nil, errors.New("failed to Get token data as Session is nil")
	}
	print(SESSION.Find(&Token{}).RowsAffected)

	p := Token{}
	SESSION.Where("hash = ?", token).Take(&p)
	return &p, nil
}

func (t *Token) IsAdmin() bool {
	if t == nil {
		return false
	}
	return t.Permission == server.AdminParam
}

func NewToken(t *Token) {
	tx := SESSION.Begin()
	tx.Save(&t)
	tx.Commit()
}
