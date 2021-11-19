package sibylValues

const (
	// NormalUser Can read from the Sibyl System.
	NormalUser UserPermission = iota
	// Enforcer Can only report to the Sibyl System.
	Enforcer
	// Inspector Can read/write directly to the Sibyl System.
	Inspector
	// Owner Can create/revoke tokens.
	Owner
)

const (
	MarkDownV2          = "markdownv2"
	OwnersTokenFileName = "owners.token"
	AppealLogDateFormat = "02-01-2006 3:04PM"
)

// Factor constants
const (
	// LowerCloudyFactor is the minimum possible value for a
	// user's crime coefficient to be counted as cloudy.
	LowerCloudyFactor = 80
	// UpperCloudyFactor is the maximum possible value for a
	// user's crime coefficient to be counted as cloudy.
	UpperCloudyFactor = 100
	// MaxAppealCoefficient is the maximum possible value for a user's crime coefficient
	// which makes the user able to appeal for unban.
	MaxAppealCoefficient = 600
	// MaxAppealCount is the maximum number of appeals a user can make.
	MaxAppealCount = 1

	// MaxTokenRevokeCount is the maximum possible times of a token being revoked
	// by a user per day.
	MaxTokenRevokeCount = 3
)

/*
	RangeCivilian     = &CrimeCoefficientRange{0, 80}
	RangeRestored   = &CrimeCoefficientRange{81, 100}
	RangeEnforcer   = &CrimeCoefficientRange{101, 125}
	RangeTrolling     = &CrimeCoefficientRange{126, 150}
	RangeSpam         = &CrimeCoefficientRange{151, 200}
	RangeEvade        = &CrimeCoefficientRange{201, 250}
	RangeCustom       = &CrimeCoefficientRange{251, 300}
	RangePsychoHazard = &CrimeCoefficientRange{301, 350}
	RangeMalImp       = &CrimeCoefficientRange{351, 400}
	RangeNSFW         = &CrimeCoefficientRange{401, 450}
	RangeRaid         = &CrimeCoefficientRange{451, 500}
	RangeSpamBot      = &CrimeCoefficientRange{501, 550}
	RangeMassAdd      = &CrimeCoefficientRange{551, 600}
*/

// flags constants
const (
	BanFlagTrolling     = "TROLLING"
	BanFlagSpam         = "SPAM"
	BanFlagEvade        = "EVADE"
	BanFlagCustom       = "CUSTOM"
	BanFlagPsychoHazard = "PSYCHOHAZARD"
	BanFlagMalImp       = "MALIMP"
	BanFlagNSFW         = "NSFW"
	BanFlagRaid         = "RAID"
	BanFlagSpamBot      = "SPAMBOT"
	BanFlagMassAdd      = "MASSADD"
)
