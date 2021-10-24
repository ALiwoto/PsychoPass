package utils

import (
	"errors"
	"gitlab.com/Dank-del/SibylAPI-Go/core/sibylConfig"
	"gitlab.com/Dank-del/SibylAPI-Go/database"
	"gitlab.com/Dank-del/SibylAPI-Go/hashing"
	"gitlab.com/Dank-del/SibylAPI-Go/server"
)

func CreateToken(Id int64, Permission string) (*database.Token, error) {
	if Permission == server.AdminParam {
		h := hashing.NewSHA1Hash(int(sibylConfig.SibylConfig.TokenSize))
		data := database.Token{Permission: Permission, UserID: Id, Hash: h}
		database.NewToken(&data)
		return &data, nil
	}
	return nil, errors.New("permission not found")
}
