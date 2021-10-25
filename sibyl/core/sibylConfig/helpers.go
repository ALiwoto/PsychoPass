package sibylConfig

import (
	"errors"
	"os"
	"strconv"

	"github.com/bigkevmcd/go-configparser"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/logging"
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

	return nil
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

func IsDebug() bool {
	if SibylConfig == nil {
		return false
	}
	return SibylConfig.Debug
}
