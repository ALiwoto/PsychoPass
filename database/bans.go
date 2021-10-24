package database

import "time"

type User struct {
	UserID int64 `json:"user_id" gorm:"primaryKey"`
	Banned bool `json:"banned"`
	Reason string `json:"reason"`
	Message string `json:"message"`
	Date time.Time `json:"date"`
}

func AddBan(UserID int64, Reason string, Message string) {
	tx := SESSION.Begin()
	ban := &User{UserID: UserID, Reason: Reason, Banned: true, Date: time.Now(), Message: Message}
	tx.Save(ban)
	tx.Commit()
}
