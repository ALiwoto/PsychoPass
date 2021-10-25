package database

import (
	"time"

	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
)

func AddBan(userID int64, reason string, message string) {
	user := &sv.User{
		UserID:  userID,
		Reason:  reason,
		Banned:  true,
		Date:    time.Now(),
		Message: message,
	}
	NewUser(user)
}

// DeleteUser will delete a user from the sibyl database.
func DeleteUser(userID int64) {
	tx := SESSION.Begin()
	u := tx.Where("user_id = ?", userID)
	if u != nil {
		u.Delete(&sv.User{})
	}
	tx.Commit()
}

// DeleteUserBan will unban a user from the sibyl database.
func DeleteUserBan(user *sv.User) {
	if user.Banned {
		user.Banned = false
		user.Reason = ""
		user.Message = ""
		user.Date = time.Now()
	} else {
		// user is not banned
		return
	}
	tx := SESSION.Begin()
	tx.Save(user)
	tx.Commit()
}
