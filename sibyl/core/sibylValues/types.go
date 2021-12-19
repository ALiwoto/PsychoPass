package sibylValues

import (
	"time"

	wc "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues/whatColor"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

type UserPermission int
type ScanStatus int
type BanFlag string
type ReportHandler func(r *Report)
type MultiReportHandler func(data *MultiScanRawData)

type StatValue struct {
	BannedCount          int64     `json:"banned_count"`
	TrollingBanCount     int64     `json:"trolling_ban_count"`
	SpamBanCount         int64     `json:"spam_ban_count"`
	EvadeBanCount        int64     `json:"evade_ban_count"`
	CustomBanCount       int64     `json:"custom_ban_count"`
	PsychoHazardBanCount int64     `json:"psycho_hazard_ban_count"`
	MalImpBanCount       int64     `json:"mal_imp_ban_count"`
	NSFWBanCount         int64     `json:"nsfw_ban_count"`
	SpamBotBanCount      int64     `json:"spam_bot_ban_count"`
	RaidBanCount         int64     `json:"raid_ban_count"`
	MassAddBanCount      int64     `json:"mass_add_ban_count"`
	CloudyCount          int64     `json:"cloudy_count"`
	TokenCount           int64     `json:"token_count"`
	InspectorsCount      int64     `json:"inspectors_count"`
	EnforcesCount        int64     `json:"enforces_count"`
	cacheTime            time.Time `json:"-"`
}

type User struct {
	UserID           int64       `json:"user_id" gorm:"primaryKey"`
	Banned           bool        `json:"banned"`
	Reason           string      `json:"reason"`
	Message          string      `json:"message"`
	BanSourceUrl     string      `json:"ban_source_url"`
	Date             time.Time   `json:"-"`
	BannedBy         int64       `json:"banned_by"`
	CrimeCoefficient int         `json:"crime_coefficient"`
	BanDate          string      `json:"date" gorm:"-" sql:"-"`
	SourceGroup      string      `json:"source_group"`
	HueColor         wc.HueColor `json:"hue_color" gorm:"-" sql:"-"`
	BanFlags         []BanFlag   `json:"ban_flags" gorm:"-" sql:"-"`
	IsBot            bool        `json:"is_bot"`
	BanCount         int         `json:"-"` // internal usage only; not meant to be seen by users.
	FlagTrolling     bool        `json:"-"`
	FlagSpam         bool        `json:"-"`
	FlagEvade        bool        `json:"-"`
	FlagCustom       bool        `json:"-"`
	FlagPsychoHazard bool        `json:"-"`
	FlagMalImp       bool        `json:"-"`
	FlagNsfw         bool        `json:"-"`
	FlagRaid         bool        `json:"-"`
	FlagSpamBot      bool        `json:"-"`
	FlagMassAdd      bool        `json:"-"`
	cacheDate        time.Time   `json:"-" gorm:"-" sql:"-"`
}

type Token struct {
	// UserId is the user id
	UserId int64 `json:"user_id" gorm:"primaryKey"`

	// Hash is the user's token hash
	Hash string `json:"hash"`

	// the user's permissions
	Permission UserPermission `json:"permission"`

	AssignedBy     int64  `json:"assigned_by"`
	DivisionNum    int    `json:"division_num"`
	AssignedReason string `json:"assigned_reason"`

	// LastUsage is the user's last usage time
	LastUsage time.Time `json:"-"`

	// CreatedAt is the creation time of the token.
	CreatedAt time.Time `json:"created_at"`

	// LastRevokeDate is the last time this user has revoked the token.
	LastRevokeDate time.Time `json:"-"`

	// RevokeCount is the amount of token being revoked by this user since
	// `LastRevokeDate` field.
	RevokeCount int `json:"-"`

	// AcceptedReports is the count of accepted reports.
	AcceptedReports int `json:"accepted_reports"`

	// DeniedReports is the count of denied reports.
	DeniedReports int `json:"denied_reports"`

	// cacheDate is the starting date of value being cached in the memory.
	cacheDate time.Time `json:"-"`
}

type Report struct {
	UniqueId       string `json:"unique_id" gorm:"primaryKey"`
	ReporterId     int64  `json:"reporter_id"`
	TargetUser     int64  `json:"target_user"`
	IsBot          bool   `json:"is_bot"`
	ReportDate     string `json:"report_date"`
	ReportReason   string `json:"report_reason"`
	ReportMessage  string `json:"report_message"`
	ScanSourceLink string `json:"scan_source_link"`
	// AgentId is the user id of the person who approved, rejected or closed the
	// scan.
	AgentId int64 `json:"agent_id"`
	// AgentReason is the reason that agent used to reject or close the
	// scan.
	AgentReason string `json:"agent_reason"`
	// AgentDate is the date that agent used to approve, reject or close the scan.
	AgentDate time.Time `json:"agent_date"`
	// ReporterPermission is the permission of the reporter.
	ReporterPermission UserPermission `json:"reporter_permission"`
	// ScanStatus is the status of this scan. it can either be approved, rejected or
	// closed.
	// please notice that only pending scans' details can be changed; if a
	// scan is approved, rejected or closed, its details cannot be changed anymore.
	ScanStatus ScanStatus `json:"scan_status"`
	// AssociationBanId is empty is the scan is not an association scan.
	AssociationBanId string `json:"association_ban_id"`
	// AgentUser is the agent user who has clicked the approve, reject or close
	// button (or sent the command).
	AgentUser *gotgbot.User `json:"-" gorm:"-" sql:"-"`
	// cacheDate is the date of this scan being cached in the memory.
	// its for internal usage only; it shouldn't be present in json, gorm or sql
	// operations.
	cacheDate time.Time `json:"-" gorm:"-" sql:"-"`
}

// CrimeCoefficientRange is the range of crime coefficients.
type CrimeCoefficientRange struct {
	start int
	end   int
}

type MultiBanUserInfo struct {
	UserId      int64  `json:"user_id"`
	Reason      string `json:"reason"`
	SourceGroup string `json:"source_group"`
	Message     string `json:"message"`
	Source      string `json:"source"`
	IsBot       bool   `json:"is_bot"`
}

type MultiBanRawData struct {
	Users []MultiBanUserInfo `json:"users"`
}

type MultiUnBanRawData struct {
	Users []int64 `json:"users"`
}

type MultiScanUserInfo struct {
	UserId  int64  `json:"user_id"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	IsBot   bool   `json:"is_bot"`
}

type MultiScanRawData struct {
	AssociationBanId string              `json:"-"`
	Users            []MultiScanUserInfo `json:"users"`
	Source           string              `json:"source"`
	GroupLink        string              `json:"group_link"`
	// ReporterPermission is the permission of the person who has sent this scan.
	// ignored by json.
	ReporterPermission UserPermission `json:"-"`
	AgentUser          *gotgbot.User  `json:"-"`
	AgentDate          time.Time      `json:"-"`
	ReporterId         int64          `json:"-"`
	Status             ScanStatus     `json:"-"`
	cacheDate          time.Time      `json:"-" `
}

type Triggers struct {
	Spam             []string `json:"spam"`
	Trolling         []string `json:"trolling"`
	Evade            []string `json:"evade"`
	MalImpersonation []string `json:"mal_impersonation"`
	PsychoHazard     []string `json:"psycho_hazard"`
	Nsfw             []string `json:"nsfw"`
	Raid             []string `json:"raid"`
	SpamBot          []string `json:"spam_bot"`
	MassAdd          []string `json:"mass_add"`
}
