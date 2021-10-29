package sibylConfig

type SibylSystemConfig struct {
	TokenSize        int64    `json:"toke_size"`
	MasterId         int64    `json:"masterid"`
	MaxPanic         int64    `json:"max_panic"`
	MaxCacheTime     int64    `json:"max_cache_time"`
	DbUrl            string   `json:"db_url"`
	DbName           string   `json:"db_name"`
	UseSqlite        bool     `json:"use_sqlite"`
	Port             string   `json:"port"`
	BotToken         string   `json:"bot_token"`
	Debug            bool     `json:"debug"`
	OrdinaryPrefixes []string `json:"cmd_prefixes"`
	CmdPrefixes      []rune   `json:"-"`
	BaseChats        []int64  `json:"base_chats"`
}
