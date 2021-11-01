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
)

// Factor constants
const (
	LowerCloudyFactor = 80
	UpperCloudyFactor = 100
)

/*
	RangeCivilian     = &CrimeCoefficientRange{0, 80}
	RangePastBanned   = &CrimeCoefficientRange{81, 100}
	RangeTrolling     = &CrimeCoefficientRange{101, 125}
	RangeSpam         = &CrimeCoefficientRange{126, 200}
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
