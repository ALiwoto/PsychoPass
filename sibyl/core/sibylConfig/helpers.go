/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */

package sibylConfig

import (
	"encoding/json"
	"io/ioutil"
	"time"

	ws "github.com/AnimeKaizoku/ssg/ssg"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
)

func LoadConfig() error {
	return LoadConfigFromFile("config.ini")
}

func LoadTriggers() error {
	var t *sibylValues.Triggers
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
		logging.Info("Loading triggers from sample file")
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

	var config = &SibylSystemConfig{}

	err := ws.ParseConfig(config, fileName)
	if err != nil {
		return err
	}

	SibylConfig = config

	return nil
}

func GetMaxCacheTime() time.Duration {
	if SibylConfig != nil {
		d := time.Duration(SibylConfig.MaxCacheTime)
		if d < 40 {
			return 40 * time.Minute
		}
		return d * time.Minute
	}
	return 40 * time.Minute
}

func GetStatsCacheTime() time.Duration {
	if SibylConfig != nil {
		return time.Duration(SibylConfig.StatsCacheTime) * time.Minute
	}
	return 0
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

func IsOwner(id int64) bool {
	if SibylConfig != nil {
		for _, owner := range SibylConfig.Owners {
			if owner == id {
				return true
			}
		}
	}
	return false
}

func IsDevUser(id int64) bool {
	if IsOwner(id) {
		return true
	}

	if SibylConfig != nil {
		for _, dev := range SibylConfig.DevUsers {
			if dev == id {
				return true
			}
		}
	}
	return false
}

func GetBaseChatIds() []int64 {
	if SibylConfig != nil {
		return SibylConfig.BaseChats
	}
	return nil
}

func GetADIds() []int64 {
	if SibylConfig != nil {
		return SibylConfig.AssaultDominators
	}
	return nil
}

func GetAppealLogChatIds() []int64 {
	if SibylConfig != nil {
		return SibylConfig.AppealLogs
	}
	return nil
}

func GetCmdPrefixes() []rune {
	if SibylConfig != nil && len(SibylConfig.CmdPrefixes) > 0 {
		if SibylConfig.cmdPrefixes == nil {
			myStrs := ws.Split(SibylConfig.CmdPrefixes, " ")
			var myRunes []rune
			for _, current := range myStrs {
				myRunes = []rune(current)
				SibylConfig.cmdPrefixes = append(SibylConfig.cmdPrefixes, myRunes[0])
			}
		}

		return SibylConfig.cmdPrefixes
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
