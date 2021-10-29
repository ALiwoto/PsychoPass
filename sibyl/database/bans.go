package database

import (
	"time"

	sv "github.com/AnimeKaizoku/sibylapi-go/sibyl/core/sibylValues"
)

func AddBan(userID, adder int64, reason, message, src string) {
	user := &sv.User{
		UserID:       userID,
		Reason:       reason,
		Banned:       true,
		Date:         time.Now(),
		Message:      message,
		BannedBy:     adder,
		BanSourceUrl: src,
	}

	NewUser(user)
}

// DeleteUser will delete a user from the sibyl database.
func DeleteUser(userID int64) {
	lockdb()
	tx := SESSION.Begin()
	u := tx.Where("user_id = ?", userID)
	if u != nil {
		u.Delete(&sv.User{})
	}
	tx.Commit()
	unlockdb()
}

// RemoveUserBan will unban a user from the sibyl database.
func RemoveUserBan(user *sv.User) {
	if user.Banned {
		user.Banned = false
		user.Reason = ""
		user.Message = ""
		user.BannedBy = 0
		user.Date = time.Now()
	} else {
		// user is not banned
		return
	}
	lockdb()
	tx := SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockdb()
}

// RemoveUserBan will unban a user from the sibyl database.
func UpdateBanparameter(user *sv.User) {
	lockdb()
	tx := SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockdb()
}
