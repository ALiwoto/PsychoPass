package sibylValues

import (
	"errors"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/ratelimiter/ratelimiter"
)

var (
	HelperBot              *gotgbot.Bot
	BotUpdater             *ext.Updater
	RateLimiter            *ratelimiter.Limiter
	SendReportHandler      ReportHandler
	SendToADHandler        AssaultDominatorHandler
	SendMultiReportHandler MultiReportHandler
	ServerStartTime        time.Time
)

var (
	ErrInvalidPerm = errors.New("invalid permission provided")
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

// crime coefficient increasement ranges
var (
	RangeCivilian     = &CrimeCoefficientRange{10, 80}
	RangeRestored     = &CrimeCoefficientRange{81, 100}
	RangeEnforcer     = &CrimeCoefficientRange{101, 150}
	RangeTrolling     = &CrimeCoefficientRange{151, 200}
	RangeSpam         = &CrimeCoefficientRange{201, 250}
	RangePsychoHazard = &CrimeCoefficientRange{251, 300}
	RangeSpamBot      = &CrimeCoefficientRange{301, 350}
	RangeCustom       = &CrimeCoefficientRange{351, 400}
	RangeNSFW         = &CrimeCoefficientRange{401, 450}
	RangeEvade        = &CrimeCoefficientRange{451, 500}
	RangeMalImp       = &CrimeCoefficientRange{501, 550}
	RangeRaid         = &CrimeCoefficientRange{551, 600}
	RangeMassAdd      = &CrimeCoefficientRange{601, 650}
)

var (

	/// This commentented out data is outdated, see line 25 for updated ones.
	// Range 0-100 (No bans) (Dominator Locked)
	// Civilian     - 0-80
	// Restored  - 81-100
	// Enforcer  - 81-100
	// Range 100-300 (Auto-mute) (Non-lethal Paralyzer)
	ReasonTrolling []string
	ReasonSpam     []string // - 151-200 - Spam/Unwanted Promotion
	ReasonEvade    []string // - 201-250 - Ban Evade using alts
	//ReasonCustom   = "custom" - 251-300 - Any Custom reason

	// Range 300+ (Ban on Sight) (Lethal Eliminator)

	ReasonMalImp       []string // - 351-400 - Malicious Impersonation
	ReasonPsychoHazard []string // - 301-350 - Bulk banned due to some bad users
	ReasonNSFW         []string // - 401-450 - Sending NSFW Content in SFW
	ReasonRaid         []string // - 451-500 - Bulk join raid to vandalize
	ReasonSpamBot      []string // - 501-550 -  Spambot, crypto, btc, forex trading scams
	ReasonMassAdd      []string // - 551-600 - Mass adding to group/channel
)

// _detailsString is the internal map for loading details string.
// as we are not using any mutex for this map, it should be read-only.
var _detailsString = map[BanFlag]string{
	BanFlagTrolling:     "Trolls aren't welcome on sane groups, when you go to groups just to annoy the admins to show much you are in control of mayhem, sibyl steps in. We do not welcome trolls and misbehavers.",
	BanFlagSpam:         "Users who post unwanted content with the aim to promote their own products or links aren't welcome around the communities we protect.",
	BanFlagEvade:        "Users that create more accounts to then evade a previously assigned ban are simply just as guilty, changing your account does not remove your previously caused drama from telegram.",
	BanFlagPsychoHazard: "You were blacklisted because you were either the owner of a group where someone was spam adding members or cause trouble for other groups and users, when you are in authority to change things and all you do is sit and watch, you are just as guilty as the person who cause the problem. Its about time you have some responsibility for the groups you were admin in.",
	BanFlagMalImp:       "You were fooling around with intentions to either harm or to either affect the profile of another established user, we do not welcome such users around our communities, you suck!",
	BanFlagNSFW:         "You were found posting pornographic or suggestively pornographic content in groups that do not welcome such content.",
	BanFlagRaid:         "You and your pals were engaged in raining a group/bot with the attempt to vandalize, this ban is unappealable and you are never welcome around our communities.",
	BanFlagSpamBot:      "You were behaving like a scam bot that attempts to ensnare users with falsified data in attempt to scam them.",
	BanFlagMassAdd:      "You were spam adding members from other groups to your own, not only this isn't welcome as platform terms of service this is also unwelcome around Sibyl, your ban will not be appealable.",
}
