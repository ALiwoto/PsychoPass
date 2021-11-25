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
	MaxAppealCoefficient = 700
	// MaxAppealCount is the maximum number of appeals a user can make.
	MaxAppealCount = 1
	// MaxTokenRevokeCount is the maximum possible times of a token being revoked
	// by a user per day.
	MaxTokenRevokeCount = 3
)

/*
	RangeCivilian     = &CrimeCoefficientRange{0, 080}
	RangeRestored     = &CrimeCoefficientRange{81, 100}
	RangeEnforcer     = &CrimeCoefficientRange{101, 150}
	RangeTROLLING     = &CrimeCoefficientRange{151, 200}
	RangeSPAM         = &CrimeCoefficientRange{201, 250}
	RangePSYCHOHAZARD = &CrimeCoefficientRange{251, 300}
	RangeSPAMBOT      = &CrimeCoefficientRange{301, 350}
	RangeCUSTOM       = &CrimeCoefficientRange{351, 400}
	RangeNSFW         = &CrimeCoefficientRange{401, 450}
	RangeEVADE        = &CrimeCoefficientRange{451, 500}
	RangeMALIMP       = &CrimeCoefficientRange{501, 550}
	RangeRAID         = &CrimeCoefficientRange{551, 600}
	RangeMASSADD      = &CrimeCoefficientRange{601, 650}
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
