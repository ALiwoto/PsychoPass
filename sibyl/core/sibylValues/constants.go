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

const (
	// Range 0-100 (No bans) (Dominator Locked)
	// Civilian     - 0-80
	// Past Banned  - 81-100
	// Range 100-300 (Auto-mute) (Non-lethal Paralyzer)
	ReasonTrolling = "trolling" // - 101-125 - Trolling
	ReasonSpam     = "spam"     // - 126-200 - Spam/Unwanted Promotion
	ReasonEvade    = "evade"    // - 201-250 - Ban Evade using alts
	//ReasonCustom   = "evade" - 251-300 - Any Custom reason

	// Range 300+ (Ban on Sight) (Lethal Eliminator)
	ReasonMalImp       = "malimp"       // - 301-350 - Bulk banned due to some bad users
	ReasonPsychoHazard = "psychohazard" // - 351-400 - Malicious Impersonation
	ReasonNSFW         = "nsfw"         // - 401-450 - Sending NSFW Content in SFW
	ReasonRaid         = "raid"         // - 451-500 - Bulk join raid to vandalize
	ReasonMassAdd      = "massadd"      // - 501-600 - Mass adding to group/channel
)
