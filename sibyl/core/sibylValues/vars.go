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
	ReasonTrolling = []string{
		"troll",
	} // - 101-125 - Trolling
	ReasonSpam = []string{
		"spam",
		"refer",
		"promo",
	} // - 126-200 - Spam/Unwanted Promotion
	ReasonEvade = []string{
		"evade",
		"banevade",
	} // - 201-250 - Ban Evade using alts
	//ReasonCustom   = "custom" - 251-300 - Any Custom reason

	// Range 300+ (Ban on Sight) (Lethal Eliminator)
	ReasonMalImp = []string{
		"malimp",
		"impersonation",
		"impersonating",
		"impersonate",
	} // - 351-400 - Malicious Impersonation
	ReasonPsychoHazard = []string{
		"psychohazard",
		"team",
		"cult",
		"gang",
		"association",
		"network", // https://t.me/SibylSystem/976
	} // - 301-350 - Bulk banned due to some bad users
	ReasonNSFW = []string{
		"nsfw",
		"cp",
		"pornography",
		"porn",
	} // - 401-450 - Sending NSFW Content in SFW
	ReasonRaid = []string{
		"raid",
		"joinraid",
	} // - 451-500 - Bulk join raid to vandalize
	ReasonSpamBot = []string{
		"crypto",
		"btc",
		"bitcoin",
		"forex",
		"trading",
		"binary",
		"thotbot",
		"joinspam",
		"binance",
		"advertis",
	} // - 501-550 -  Spambot, crypto, btc, forex trading scams
	ReasonMassAdd = []string{
		"massadd",
		"kidnap",
		"memberscraping",
	} // - 551-600 - Mass adding to group/channel
)
