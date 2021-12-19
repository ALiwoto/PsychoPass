package sibylConfig

type SibylSystemConfig struct {
	// TokenSize the token size; it should be at least 20.
	TokenSize int64 `section:"general" key:"token_size" default:"20"`

	// Owners is the IDs of the owners.
	Owners []int64 `section:"general" key:"owners"`

	// DevUsers is the IDs of the owners.
	DevUsers []int64 `section:"general" key:"dev_users"`

	// AssaultDominators is the list of userbots known as
	// "assault dominators".
	AssaultDominators []int64 `section:"general" key:"assault_dominators"`

	// Port is the http port that we should listen to
	Port string `section:"general" key:"port" default:"8080"`

	// MaxPanic is the maximum amount of allowed panics to be caught.
	// -1 for unlimited.
	MaxPanic int64 `section:"general" key:"max_panics" default:"-1"`

	// Debug is the debug mode.
	Debug bool `section:"general" key:"debug" default:"false"`

	// DbUrl is the of your postgresql database.
	// if `use_sqlite` is set to `true`, this variable will be ignored.
	DbUrl string `section:"database" key:"url"`

	// MaxCacheTime is the max cache time for database.
	MaxCacheTime int64 `section:"database" key:"max_cache_time" default:"60"`
	// UseSqlite is a bool. set this to `true` if you want to use sqlite database.
	// this is not recommended for production version of Sibyl System.
	UseSqlite bool `section:"database" key:"use_sqlite" default:"true"`

	// DbName is the database name. if `use_sqlite` is true, this value is required;
	// in that case, if it's empty, it will be set to `sibyldb` by default.
	DbName string `section:"database" key:"db_name"`

	// StatsCacheTime is the amount of stats to be cached in memory in minutes.
	// set it to 0 or comment it out if you don't want it to be cached.
	StatsCacheTime int64 `section:"database" key:"stats_cache_time" default:"60"`

	// BotToken is the helper bot's token.
	// it can be commented out (or set to empty) if you don't want
	// the application to interact with telegram directly.
	BotToken string `section:"telegram" key:"bot_token"`

	// BaseChats is the base group's IDs. separate them using " " or ",".
	// values in base chats can be anything: a user's pm, a channel or a group.
	BaseChats []int64 `section:"telegram" key:"base_chats"`

	// AppealLogs is the chat id of the chat where the appeal logs will be sent.
	AppealLogs []int64 `section:"telegram" key:"appeal_logs"`

	// CmdPrefixes is the command prefixes of the bot.
	CmdPrefixes string `section:"telegram" key:"cmd_prefixes"`
	cmdPrefixes []rune

	// BotAPIUrl is the url of the bot-api server (optional).
	BotAPIUrl string `section:"telegram" key:"api_url"`

	DropUpdates bool `section:"telegram" key:"drop_updates" default:"true"`

	RateLimiterPunishmentTime int64 `section:"telegram" key:"ratelimiter_punishment_time"`
	RateLimiterTimeout        int64 `section:"telegram" key:"ratelimiter_timeout"`
	RateLimiterMaxMessages    int64 `section:"telegram" key:"ratelimiter_max_messages"`
	RateLimiterMaxCache       int64 `section:"telegram" key:"ratelimiter_max_cache"`
}
