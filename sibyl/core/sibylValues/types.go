package sibylValues

import "time"

type UserPermission int
type ReportHandler func(r *Report)

type User struct {
	UserID           int64     `json:"user_id" gorm:"primaryKey"`
	Banned           bool      `json:"banned"`
	Reason           string    `json:"reason"`
	Message          string    `json:"message"`
	BanSourceUrl     string    `json:"ban_source_url"`
	Date             time.Time `json:"date"`
	BannedBy         int64     `json:"banned_by"`
	CrimeCoefficient int       `json:"crime_coefficient"`
	cacheDate        time.Time `json:"-"`
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

	cacheDate time.Time `json:"-"`
}

type Report struct {
	ReporterId         int64
	TargetUser         int64
	ReportDate         string
	ReportReason       string
	ReportMessage      string
	ReporterPermission string
}

// CrimeCoefficientRange is the range of crime coefficients.
type CrimeCoefficientRange struct {
	start int
	end   int
}
