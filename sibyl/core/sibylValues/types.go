package sibylValues

import (
	"time"
)

/*
Coefficients and Flags

==== Flags     -  ========
Range 0-100 (No bans) (Dominator Locked)
• Civilian     - 0-80
• Restored  - 81-100
==============
Range 100-300 (Auto-mute) (Non-lethal Paralyzer)
• TROLLING     - 101-125 - Trolling
• SPAM         - 126-200 - Spam/Unwanted Promotion
• EVADE        - 201-250 - Ban Evade using alts
x-------x
Manual Revert
• CUSTOM       - 251-300 - Any Custom reason
x-------x
==============
Range 300+ (Ban on Sight) (Lethal Eliminator)
• PSYCHOHAZARD - 301-350 - Bulk banned due to some bad users
• MALIMP       - 351-400 - Malicious Impersonation
• NSFW         - 401-450 - Sending NSFW Content in SFW
• RAID         - 451-500 - Bulk join raid to vandalize
• SPAMBOT       - 501-550 - SpamBot, crypto, btc, forex trading scams
• MASSADD      - 501-600 - Mass adding to group/channel
==============
*/

type UserPermission int
type BanFlag string
type ReportHandler func(r *Report)

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
	UserID           int64     `json:"user_id" gorm:"primaryKey"`
	Banned           bool      `json:"banned"`
	Reason           string    `json:"reason"`
	Message          string    `json:"message"`
	BanSourceUrl     string    `json:"ban_source_url"`
	Date             time.Time `json:"-"`
	BannedBy         int64     `json:"banned_by"`
	CrimeCoefficient int       `json:"crime_coefficient"`
	BanDate          string    `json:"date" gorm:"-" sql:"-"`
	SourceGroup      string    `json:"source_group"`
	BanFlags         []BanFlag `json:"ban_flags" gorm:"-" sql:"-"`
	IsBot            bool      `json:"is_bot"`
	BanCount         int       `json:"-"` // internal usage only; not meant to be seen by users.
	FlagTrolling     bool      `json:"-"`
	FlagSpam         bool      `json:"-"`
	FlagEvade        bool      `json:"-"`
	FlagCustom       bool      `json:"-"`
	FlagPsychoHazard bool      `json:"-"`
	FlagMalImp       bool      `json:"-"`
	FlagNsfw         bool      `json:"-"`
	FlagRaid         bool      `json:"-"`
	FlagSpamBot      bool      `json:"-"`
	FlagMassAdd      bool      `json:"-"`
	cacheDate        time.Time `json:"-"`
}

type Token struct {
	// UserId is the user id
	UserId int64 `json:"user_id" gorm:"primaryKey"`

	// Hash is the user's token hash
	Hash string `json:"hash"`

	// the user's permissions
	Permission UserPermission `json:"permission"`

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
	ReporterId         int64
	TargetUser         int64
	IsBot              bool
	ReportDate         string
	ReportReason       string
	ReportMessage      string
	ScanSourceLink     string
	ReporterPermission string
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
