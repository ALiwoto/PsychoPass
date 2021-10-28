package database

import "time"

type Scan ScanSource

type ScanSource struct {
	Id        uint64    `json:"id" gorm:"primaryKey"`
	InvokedBy int64     `json:"invoked_by"`
	UserID    int64     `json:"user_id"`
	Reason    string    `json:"reason"`
	Banned    bool      `json:"banned"`
	Date      time.Time `json:"date"`
	Message   string    `json:"message"`
	BannedBy  int64     `json:"banned_by"`
}

func AddScan(UserId int64, Reason string, Message string, invokedVy int64) {
	tx := SESSION.Begin()
	ss := &ScanSource{UserID: UserId, Reason: Reason, Message: Message, Date: time.Now(), InvokedBy: invokedVy}
	tx.Save(ss)
	tx.Commit()
}

func (s *ScanSource) Approve(by int64) {
	tx := SESSION.Begin()
	tx.Model(ScanSource{}).Update("banned", true).Update("banned_by", by)
	tx.Commit()
}

func (s *ScanSource) Reject() {
	if s == nil || s.UserID == 0 {
		return
	}
	tx := SESSION.Begin()
	tx.Delete(&s)
	tx.Commit()
}
