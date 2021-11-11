package database

import (
	"time"

	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
)

func AddBan(userID, adder int64, reason, message, src string, isBot bool, count int) *sv.User {
	user := &sv.User{
		UserID:       userID,
		Banned:       true,
		Date:         time.Now(),
		Message:      message,
		BannedBy:     adder,
		BanSourceUrl: src,
		IsBot:        isBot,
		BanCount:     count,
	}
	user.SetAsBanReason(reason)
	user.IncreaseCrimeCoefficientAuto()
	NewUser(user)
	return user
}

func AddBanByInfo(info *sv.MultiBanUserInfo, adder int64, count int) *sv.User {
	return AddBan(info.UserId, adder, info.Reason, info.Message, info.Source, info.IsBot, count)
}

// DeleteUser will delete a user from the sibyl database.
func DeleteUser(userID int64) {
	lockdb()
	tx := SESSION.Begin()
	u := tx.Model(modelUser).Where("user_id = ?", userID)
	if u != nil {
		u.Delete(&sv.User{})
	}
	tx.Commit()
	unlockdb()
}

// RemoveUserBan will unban a user from the sibyl database.
func RemoveUserBan(user *sv.User, clearHistory bool) {
	if !user.Banned {
		// don't send any query to database if user is not banned.
		return
	}
	user.FormatBanDate()
	user.SetAsPastBan(clearHistory)
	lockdb()
	tx := SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockdb()
}

// ClearHistory will unban a user from the sibyl database.
func ClearHistory(user *sv.User) {
	if user.Banned {
		return
	}
	user.ClearHistory()
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
