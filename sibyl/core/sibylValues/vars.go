package sibylValues

import (
	"errors"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/ratelimiter/ratelimiter"
)

var (
	HelperBot         *gotgbot.Bot
	BotUpdater        *ext.Updater
	RateLimiter       *ratelimiter.Limiter
	SendReportHandler ReportHandler
	ServerStartTime   time.Time
)

var (
	ErrInvalidPerm = errors.New("invalid permission provided")
)

/*
Coefficients and Flags

==== Flags     -  ========
Range 0-100 (No bans) (Dominator Locked)
• Civilian     - 0-80
• PastBanned  - 81-100
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
• MASSADD      - 551-600 - Mass adding to group/channel
==============

Trigger word aliases
EVADE   - evade, banevade
MALIMP  - impersonation, malimp, fake profile
NSFW    - porn, pornography, nsfw, cp
Crypto  - btc, crypto, forex, trading, binary
MASSADD - spam add, kidnapping, member scraping, member adding, mass adding, spam adding, bulk adding
*/

// crime coefficient increasement ranges
var (
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
)

var (
	// Range 0-100 (No bans) (Dominator Locked)
	// Civilian     - 0-80
	// Past Banned  - 81-100
	// Range 100-300 (Auto-mute) (Non-lethal Paralyzer)
	ReasonTrolling []string
	ReasonSpam     []string // - 126-200 - Spam/Unwanted Promotion
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
var _detailsString = map[string]string{
	BanFlagTrolling:     "Trolls aren't welcome on sane groups, when you go to groups just to annoy the admins to show much you are in control of mayhem, sibyl steps in. We do not welcome trolls and misbehavers.",
	BanFlagSpam:         "Users who post unwanted content with the aim to promote their own products or links aren't welcome around the communities we protect.",
	BanFlagEvade:        "Users that create more accounts to then evade a previously assigned ban are simply just as guilty, changing your account does not remove your previously caused drama from telegram.",
	BanFlagPsychoHazard: "You were blacklisted because you were either the owner of a group where someone was spam adding members or cause trouble for other groups and users, when you are in authority to change things and all you do is sit and watch, you are just as guilty as the person who cause the problem. Its about time you have some responsibility for the groups you were admin in.",
	BanFlagMalImp:       "You were fooling around with intentions to either harm or to either affect the profile of another established user, we do not welcome such users around our communities, you suck!",
	BanFlagNSFW:         "You were found posting pornographic or suggestively pornographic content in groups that do not welcome such content.",
	BanFlagRaid:         "You and your pals were engaged in raining a group/bot with the attempt to vandalize, this ban is unappealable and you are never welcome around our communities.",
	BanFlagSpamBot:      "You were behaving like a scam bot that attempts to ensnare users with falsified data in attempt to scam them.",
	BanFlagMassAdd:      "You were spam adding members from other groups to your own, not only this not welcome as platform terms of service this is also unwelcome around Sibyl, your ban will not be appealable.",
}
