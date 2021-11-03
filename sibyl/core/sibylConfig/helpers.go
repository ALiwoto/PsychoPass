package sibylConfig

import (
	"encoding/json"
	"errors"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"

	"github.com/bigkevmcd/go-configparser"
)

func LoadConfig() error {
	return LoadConfigFromFile("config.ini")
}

func LoadTriggers() error {
	var t *sibylValues.Triggers
	logging.Info("Loading triggers from triggers.json")
	f, err := ioutil.ReadFile("triggers.json")
	if err == nil {
		err = json.Unmarshal(f, &t)
		if err != nil {
			return err
		}
		sibylValues.ReasonEvade = t.Evade
		sibylValues.ReasonMalImp = t.MalImpersonation
		sibylValues.ReasonNSFW = t.Nsfw
		sibylValues.ReasonTrolling = t.Trolling
		sibylValues.ReasonMassAdd = t.MassAdd
		sibylValues.ReasonSpam = t.Spam
		sibylValues.ReasonPsychoHazard = t.PsychoHazard
		sibylValues.ReasonSpamBot = t.SpamBot
		sibylValues.ReasonRaid = t.Raid
	} else {
		logging.Info("Failed to load from triggers.json")
		logging.Info("Loading triggers from sample_triggers.json")
		f, err := ioutil.ReadFile("sample_triggers.json")
		if err != nil {
			return err
		}
		err = json.Unmarshal(f, &t)
		if err != nil {
			return err
		}
		sibylValues.ReasonEvade = t.Evade
		sibylValues.ReasonMalImp = t.MalImpersonation
		sibylValues.ReasonNSFW = t.Nsfw
		sibylValues.ReasonTrolling = t.Trolling
		sibylValues.ReasonMassAdd = t.MassAdd
		sibylValues.ReasonSpam = t.Spam
		sibylValues.ReasonPsychoHazard = t.PsychoHazard
		sibylValues.ReasonSpamBot = t.SpamBot
		sibylValues.ReasonRaid = t.Raid
	}
	return nil
}

func LoadConfigFromFile(fileName string) error {
	if SibylConfig != nil {
		return nil
	}

	SibylConfig = &SibylSystemConfig{}
	env := os.Getenv
	configContent, err := configparser.NewConfigParserFromFile(fileName)
	if err != nil {
		return err
	}

	// general section variables:
	SibylConfig.Port, err = configContent.Get("general", "port")
	if err != nil {
		SibylConfig.Port = env("PORT")
		if len(SibylConfig.Port) == 0 {
			logging.Error("No port specified in config file or environment variable." +
				"Using default port 8080")
			SibylConfig.Port = "8080"
		}
	}

	SibylConfig.TokenSize, err = configContent.GetInt64("general", "token_size")
	if err != nil {
		tokenSizeStr := env("TOKEN_SIZE")
		SibylConfig.TokenSize, _ = strconv.ParseInt(tokenSizeStr, 10, 64)
		if SibylConfig.TokenSize == 0 {
			logging.Error("No token-size specified in config file or environment variable." +
				"Using default token-size 64")
			SibylConfig.TokenSize = 64
		}
	}

	if SibylConfig.TokenSize < 20 {
		logging.Fatal("Exiting cause token size is less than 20")
	}

	ownersStr, err := configContent.Get("general", "owners")
	if err != nil || len(ownersStr) == 0 {
		ownersStr = env("OWNERS")
	}
	SibylConfig.Owners = parseBaseStr(strings.TrimSpace(ownersStr))

	SibylConfig.MaxPanic, err = configContent.GetInt64("general", "max_panics")
	if err != nil {
		SibylConfig.MaxPanic, _ = strconv.ParseInt(env("MAX_PANICS"), 10, 64)
		if SibylConfig.MaxPanic == 0 {
			SibylConfig.MaxPanic = -1
		}
	}

	SibylConfig.Debug, err = configContent.GetBool("general", "debug")
	if err != nil {
		debug := env("SIBYL_DEBUG")
		SibylConfig.Debug = debug == "yes" || debug == "true"
	}

	// database section variables:
	SibylConfig.UseSqlite, err = configContent.GetBool("database", "use_sqlite")
	if err != nil {
		usesqlite := env("USE_SQLITE")
		SibylConfig.UseSqlite = usesqlite == "yes" || usesqlite == "true"
	}

	SibylConfig.DbUrl, err = configContent.Get("database", "url")
	if err != nil || len(SibylConfig.DbUrl) == 0 {
		SibylConfig.DbUrl = env("DB_URL")
		if len(SibylConfig.DbUrl) == 0 && !SibylConfig.UseSqlite {
			return errors.New("no database url is specified")
		}
	}

	SibylConfig.DbName, err = configContent.Get("database", "db_name")
	if err != nil || len(SibylConfig.DbUrl) == 0 {
		SibylConfig.DbName = env("DB_NAME")
		if len(SibylConfig.DbName) == 0 {
			SibylConfig.DbName = "sibyldb"
		}
	}

	SibylConfig.MaxCacheTime, err = configContent.GetInt64("database", "max_cache_time")
	if err != nil {
		SibylConfig.MaxCacheTime, _ = strconv.ParseInt(env("MAX_CACHE_TIME"), 10, 64)
	}

	// telegram section variables
	SibylConfig.BotToken, err = configContent.Get("telegram", "bot_token")
	if err != nil || len(SibylConfig.BotToken) == 0 {
		SibylConfig.BotToken = env("BOT_TOKEN")
	}

	SibylConfig.BotAPIUrl, err = configContent.Get("telegram", "api_url")
	if err != nil || len(SibylConfig.BotAPIUrl) == 0 {
		SibylConfig.BotAPIUrl = env("API_URL")
	}

	// database section variables:
	SibylConfig.DropUpdates, err = configContent.GetBool("telegram", "drop_updates")
	if err != nil {
		dropUpdates := env("DROP_UPDATES")
		SibylConfig.DropUpdates = dropUpdates == "yes" || dropUpdates == "true"
	}

	baseStr, err := configContent.Get("telegram", "base_chats")
	if err != nil || len(baseStr) == 0 {
		baseStr = env("BASE_CHATS")
	}
	SibylConfig.BaseChats = parseBaseStr(strings.TrimSpace(baseStr))

	preStr, err := configContent.Get("telegram", "cmd_prefixes")
	if err != nil || len(preStr) == 0 {
		preStr = env("CMD_PREFIXES")
	}
	SibylConfig.CmdPrefixes = parseCmdPrefixes(preStr)

	/*
		# ratelimiter's punishment (ignoring) time in minutes.
		ratelimiter_punishment_time = 40
		# ratelimiter's message sending timeout. (in seconds)
		ratelimiter_timeout = 4
		# ratelimiter's message sending interval. if user sends more than this amount
		# of messages per `ratelimiter_timeout` period, bot will ignore him for
		# `ratelimiter_punishment_time` minutes.
		ratelimiter_max_messages = 6
		# ratelimiter's maximum amount of caching for a user. (in minutes)
		# recommended to be more than `ratelimiter_punishment_time` +
		# `ratelimiter_timeout`; otherwise will be ignored by library itself.
		ratelimiter_max_cache = 50
	*/

	SibylConfig.RateLimiterPunishmentTime, err =
		configContent.GetInt64("telegram", "ratelimiter_punishment_time")
	if err != nil {
		SibylConfig.RateLimiterPunishmentTime, _ =
			strconv.ParseInt(env("RATELIMITER_PUNISHMENT_TIME"), 10, 64)
	}

	SibylConfig.RateLimiterTimeout, err =
		configContent.GetInt64("telegram", "ratelimiter_timeout")
	if err != nil {
		SibylConfig.MaxCacheTime, _ =
			strconv.ParseInt(env("RATELIMITER_TIMEOUT"), 10, 64)
	}

	SibylConfig.RateLimiterMaxMessages, err =
		configContent.GetInt64("telegram", "ratelimiter_max_messages")
	if err != nil {
		SibylConfig.MaxCacheTime, _ =
			strconv.ParseInt(env("RATELIMITER_MAX_MESSAGE"), 10, 64)
	}

	SibylConfig.RateLimiterMaxCache, err =
		configContent.GetInt64("telegram", "ratelimiter_max_cache")
	if err != nil {
		SibylConfig.MaxCacheTime, _ =
			strconv.ParseInt(env("RATELIMITER_MAX_CACHE"), 10, 64)
	}

	return nil
}

func parseCmdPrefixes(value string) []rune {
	if len(value) == 0 {
		return []rune{'!', '/'}
	}

	value = strings.TrimSpace(value)
	if strings.Contains(value, " ") {
		var all []rune
		mystrs := ws.FixSplitWhite(strings.Split(value, " "))
		for _, str := range mystrs {
			all = append(all, rune(str[0]))
		}
		return all
	} else {
		if len(value) > 0 {
			return []rune(value)
		}
		return nil
	}
}

func parseBaseStr(value string) []int64 {
	if !strings.Contains(value, " ") && !strings.Contains(value, ",") {
		value = strings.TrimSpace(value)
		tmp, err := strconv.ParseInt(value, 10, 64)
		if err != nil || tmp == 0 {
			return nil
		}
		return []int64{tmp}
	}

	mystrs := ws.Split(value, " ", ",")
	if len(mystrs) == 0 {
		return nil
	}

	var tmp int64
	var err error
	var all []int64
	for _, str := range mystrs {
		tmp, err = strconv.ParseInt(str, 10, 64)
		if err != nil || tmp == 0 {
			continue
		}
		all = append(all, tmp)
	}

	return all
}

func GetMaxCacheTime() time.Duration {
	if SibylConfig != nil {
		return time.Duration(SibylConfig.MaxCacheTime) * time.Minute
	}
	return 40 * time.Minute
}

func GetPort() string {
	if SibylConfig != nil {
		return SibylConfig.Port
	}
	return "8080" // default port is set to 8080
}

func DropUpdates() bool {
	if SibylConfig != nil {
		return SibylConfig.DropUpdates
	}
	return true
}

func GetMaxPanics() int64 {
	if SibylConfig != nil {
		return SibylConfig.MaxPanic
	}
	return 0
}

func GetMaxHashSize() int64 {
	if SibylConfig != nil {
		return SibylConfig.TokenSize
	}
	return 0
}

func GetBotToken() string {
	if SibylConfig != nil {
		return SibylConfig.BotToken
	}
	return ""
}

func GetAPIUrl() string {
	if SibylConfig != nil {
		return SibylConfig.BotAPIUrl
	}
	return ""
}

func GetOwnersID() []int64 {
	if SibylConfig != nil {
		return SibylConfig.Owners
	}
	return nil
}

func GetBaseChatIds() []int64 {
	if SibylConfig != nil {
		return SibylConfig.BaseChats
	}
	return nil
}

func GetCmdPrefixes() []rune {
	if SibylConfig != nil {
		return SibylConfig.CmdPrefixes
	}
	return []rune{'/', '!', '?'}
}

func IsDebug() bool {
	if SibylConfig == nil {
		return false
	}
	return SibylConfig.Debug
}

func GetRateLimiterPunishmentTime() time.Duration {
	if SibylConfig == nil {
		return 40 * time.Minute
	}
	return time.Duration(SibylConfig.RateLimiterPunishmentTime) * time.Minute
}

func GetRateLimiterTimeout() time.Duration {
	if SibylConfig == nil {
		return 4 * time.Second
	}
	return time.Duration(SibylConfig.RateLimiterTimeout) * time.Second
}

func GetRateLimiterMaxMessages() int64 {
	if SibylConfig == nil {
		return 6
	}
	return SibylConfig.RateLimiterMaxMessages
}

func GetRateLimiterMaxCache() time.Duration {
	if SibylConfig == nil {
		return 50 * time.Minute
	}
	return time.Duration(SibylConfig.RateLimiterMaxCache) * time.Minute
}
