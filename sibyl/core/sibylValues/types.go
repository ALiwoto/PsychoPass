package sibylValues

import "time"

/*
Coefficients and Flags

==== Flags     -  ========
Range 0-100 (No bans) (Dominator Locked)
• Civilian     - 0-80
• Past Banned  - 81-100
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
• MASSADD      - 501-600 - Mass adding to group/channel
==============
*/

type UserPermission int
type UserFlag int
type ReportHandler func(r *Report)

type StatValue struct {
	BannedCount          int64 `json:"banned_count"`
	TrollingBanCount     int64 `json:"trolling_ban_count"`
	SpamBanCount         int64 `json:"spam_ban_count"`
	EvadeBanCount        int64 `json:"evade_ban_count"`
	CustomBanCount       int64 `json:"custom_ban_count"`
	PsychoHazardBanCount int64 `json:"psycho_hazard_ban_count"`
	MalImpBanCount       int64 `json:"mal_imp_ban_count"`
	NSFWBanCount         int64 `json:"nsfw_ban_count"`
	RaidBanCount         int64 `json:"raid_ban_count"`
	MassAddBanCount      int64 `json:"mass_add_ban_count"`
	CloudyCount          int64 `json:"cloudy_count"`
	TokenCount           int64 `json:"token_count"`
	InspectorsCount      int64 `json:"inspectors_count"`
	EnforcesCount        int64 `json:"enforces_count"`
}

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
	FlagTrolling     bool      `json:"-"`
	FlagSpam         bool      `json:"-"`
	FlagEvade        bool      `json:"-"`
	FlagCustom       bool      `json:"-"`
	FlagPsychoHazard bool      `json:"-"`
	FlagMalImp       bool      `json:"-"`
	FlagNsfw         bool      `json:"-"`
	FlagRaid         bool      `json:"-"`
	FlagMassAdd      bool      `json:"-"`
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
