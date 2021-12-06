package database

import (
	"time"

	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

func AddBan(info *BanInfo) *sv.User {
	user := &sv.User{
		UserID:       info.UserID,
		Banned:       true,
		Date:         time.Now(),
		Message:      info.Message,
		BannedBy:     info.Adder,
		BanSourceUrl: info.Src,
		SourceGroup:  info.SrcGroup,
		IsBot:        info.IsBot,
		BanCount:     info.Count,
	}

	user.SetAsBanReason(info.Reason)
	user.IncreaseCrimeCoefficientAuto()
	NewUser(user)
	return user
}

func AddBanByInfo(info *sv.MultiBanUserInfo, adder int64, count int) *sv.User {
	return AddBan(
		&BanInfo{
			UserID:   info.UserId,
			Adder:    adder,
			Reason:   info.Reason,
			SrcGroup: info.SourceGroup,
			Message:  info.Message,
			Src:      info.Source,
			IsBot:    info.IsBot,
			Count:    count,
		},
	)
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
	user.SetAsRestored(clearHistory)
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

// UpdateBanparameter will update a user's ban parameter into the database.
func UpdateBanparameter(user *sv.User) {
	lockdb()
	tx := SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockdb()
}

func UpdateUserCrimeCoefficientByPerm(user *sv.User, perm sv.UserPermission) {
	pre := user.CrimeCoefficient
	user.IncreaseCrimeCoefficientByPerm(perm)
	if pre == user.CrimeCoefficient {
		// don't send any query to database if user's crime coefficient
		// is not changed.
		return
	}

	lockdb()
	tx := SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockdb()
}
