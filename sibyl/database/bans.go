package database

import (
	"time"

	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
)

func AddBan(UserID int64, Reason string, Message string) {
	tx := SESSION.Begin()
	ban := &sv.User{
		UserID: UserID,
		Reason: Reason, Banned: true,
		Date:    time.Now(),
		Message: Message,
	}
	tx.Save(ban)
	tx.Commit()
}
