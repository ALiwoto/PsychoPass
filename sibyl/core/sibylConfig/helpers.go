package sibylConfig

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/utils/logging"
	"github.com/bigkevmcd/go-configparser"
)

func LoadConfig() error {
	return LoadConfigFromFile("config.ini")
}

func LoadConfigFromFile(fileName string) error {
	if SibylConfig != nil {
		return nil
	}

	SibylConfig = &SibylSystemConfig{}
	configContent, err := configparser.NewConfigParserFromFile(fileName)
	if err != nil {
		return err
	}

	// general section variables:
	SibylConfig.Port, err = configContent.Get("general", "port")
	if err != nil {
		SibylConfig.Port = os.Getenv("PORT")
		if len(SibylConfig.Port) == 0 {
			logging.Error("No port specified in config file or environment variable." +
				"Using default port 8080")
			SibylConfig.Port = "8080"
		}
	}

	SibylConfig.TokenSize, err = configContent.GetInt64("general", "token_size")
	if err != nil {
		tokenSizeStr := os.Getenv("TOKEN_SIZE")
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

	SibylConfig.MasterId, err = configContent.GetInt64("general", "masterid")
	if err != nil {
		SibylConfig.MasterId, _ = strconv.ParseInt(os.Getenv("MASTER_ID"), 10, 64)
	}

	SibylConfig.MaxPanic, err = configContent.GetInt64("general", "max_panics")
	if err != nil {
		SibylConfig.MaxPanic, _ = strconv.ParseInt(os.Getenv("MAX_PANICS"), 10, 64)
		if SibylConfig.MaxPanic == 0 {
			SibylConfig.MaxPanic = -1
		}
	}

	SibylConfig.Debug, err = configContent.GetBool("general", "debug")
	if err != nil {
		debug := os.Getenv("SIBYL_DEBUG")
		SibylConfig.Debug = debug == "yes" || debug == "true"
	}

	// database section variables:
	SibylConfig.UseSqlite, err = configContent.GetBool("database", "use_sqlite")
	if err != nil {
		usesqlite := os.Getenv("USE_SQLITE")
		SibylConfig.UseSqlite = usesqlite == "yes" || usesqlite == "true"
	}

	SibylConfig.DbUrl, err = configContent.Get("database", "url")
	if err != nil || len(SibylConfig.DbUrl) == 0 {
		SibylConfig.DbUrl = os.Getenv("DB_URL")
		if len(SibylConfig.DbUrl) == 0 && !SibylConfig.UseSqlite {
			return errors.New("no database url is specified")
		}
	}

	SibylConfig.DbName, err = configContent.Get("database", "db_name")
	if err != nil || len(SibylConfig.DbUrl) == 0 {
		SibylConfig.DbName = os.Getenv("DB_NAME")
		if len(SibylConfig.DbName) == 0 {
			SibylConfig.DbName = "sibyldb"
		}
	}

	// telegram section variables
	SibylConfig.BotToken, err = configContent.Get("telegram", "bot_token")
	if err != nil || len(SibylConfig.BotToken) == 0 {
		SibylConfig.BotToken = os.Getenv("BOT_TOKEN")
	}

	baseStr, err := configContent.Get("telegram", "base_chats")
	if err != nil || len(SibylConfig.BotToken) == 0 {
		baseStr = os.Getenv("BASE_CHATS")
	}
	SibylConfig.BaseChats = parseBaseStr(strings.TrimSpace(baseStr))

	preStr, err := configContent.Get("telegram", "cmd_prefixes")
	if err != nil || len(SibylConfig.BotToken) == 0 {
		preStr = os.Getenv("CMD_PREFIXES")
	}
	SibylConfig.CmdPrefixes = parseCmdPrefixes(preStr)

	SibylConfig.MaxCacheTime, err = configContent.GetInt64("database", "max_cache_time")
	if err != nil {
		SibylConfig.MaxCacheTime, _ = strconv.ParseInt(os.Getenv("MAX_CACHE_TIME"), 10, 64)
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

func GetMasterId() int64 {
	if SibylConfig != nil {
		return SibylConfig.MasterId
	}
	return 0
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
