package sibylValues

import "time"

type UserPermission int
type ReportHandler func(r *Report)
type reportState int

type User struct {
	UserID   int64     `json:"user_id" gorm:"primaryKey"`
	Banned   bool      `json:"banned"`
	Reason   string    `json:"reason"`
	Message  string    `json:"message"`
	Date     time.Time `json:"date"`
	BannedBy int64     `json:"banned_by"`
}

type Token struct {
	// the user id
	UserId int64 `json:"user_id" gorm:"primaryKey"`

	// the user hash
	Hash string `json:"hash"`

	// the user's permissions
	Permission UserPermission `json:"permission"`

	// the user's last usage time
	LastUsage time.Time `json:"-"`

	// Creation time
	CreatedAt time.Time `json:"created_at"`

	AcceptedReports int `json:"accepted_reports"`

	DeniedReports int `json:"denied_reports"`
}

type Report struct {
	ReporterId         int64
	TargetUser         int64
	ReportDate         string
	ReportReason       string
	ReportMessage      string
	ReporterPermission string
	uniqueId           int64
	date               time.Time
	state              reportState
}
