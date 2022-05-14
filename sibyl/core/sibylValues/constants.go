/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package sibylValues

// possible values for type `UserPermission`.
// please notice that any other value for this type, is considered
// as "invalid" by API.
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

// possible values for type `ScanStatus`.
// please notice that any other value for this type, is considered
// as "invalid" by API.
const (
	// ScanPending is the status of a pending scan.
	// it can be approved, rejected or closed.
	ScanPending ScanStatus = iota
	// ScanApproved is the status of an approved scan.
	// the details of an approved scan cannot be changed anymore.
	ScanApproved
	// ScanRejected is the status of a rejected scan.
	// the details of a rejected scan cannot be changed anymore.
	ScanRejected
	// ScanClosed is the status of a closed scan.
	// the details of a closed scan cannot be changed anymore.
	ScanClosed
)

const (
	// EntityTypeUser represents a normal user while being scanned.
	// please notice that "being normal", doesn't necessarily mean
	// not being criminal.
	EntityTypeUser EntityType = iota
	// EntityTypeBot represents an account which is considered as a bot.
	// as API has no idea what is a "bot account", the value "is_bot"
	// should be set by the enforcer/inspector while sending requests
	// to sibyl.
	EntityTypeBot
	// EntityTypeAdmin represents an account which is considered as an admin
	// in a psychohazard event. it's completely up to the person who is scanning
	// to decide what is an admin account.
	EntityTypeAdmin
	// EntityTypeOwner represents an account which is considered as an owner
	// in a psychohazard event. it's completely up to the person who is scanning
	// to decide what is an owner account.
	EntityTypeOwner
	// EntityTypeChannel represents an entity which is considered as a channel.
	EntityTypeChannel
	// EntityTypeGroup represents an entity which is considered as a group.
	EntityTypeGroup
)

const (
	MarkDownV2            = "markdownv2"
	AssociationScanPrefix = "AC-"
	OwnersTokenFileName   = "owners.token"
	AppealLogDateFormat   = "02-01-2006 3:04PM"
)

const (
	permaCheckerValue    = "perma"
	permaAppendingReason = " | perma ban"
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
	RangeCivilian     = &CrimeCoefficientRange{10, 080}
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
