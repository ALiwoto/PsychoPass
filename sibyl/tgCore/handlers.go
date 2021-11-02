package tgCore

import (
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylConfig"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/tgCore/plugins/infoPlugin"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/tgCore/plugins/reportPlugin"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/tgCore/plugins/tokenPlugin"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/ratelimiter/ratelimiter"
)

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	loadLimiter(d)
	infoPlugin.LoadAllHandlers(d, triggers)
	reportPlugin.LoadAllHandlers(d, triggers)
	tokenPlugin.LoadAllHandlers(d, triggers)
}

func loadLimiter(d *ext.Dispatcher) {
	limiter := ratelimiter.NewLimiter(d, false, false)
	limiter.TextOnly = true
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
	pt := sibylConfig.GetRateLimiterPunishmentTime()
	timeout := sibylConfig.GetRateLimiterTimeout()
	maxMessages := sibylConfig.GetRateLimiterMaxMessages()
	maxCache := sibylConfig.GetRateLimiterMaxCache()
	if pt != 0 {
		limiter.SetPunishmentDuration(pt)
	}
	if timeout != 0 {
		limiter.SetFloodWaitTime(timeout)
	}
	if maxMessages != 0 {
		limiter.SetMaxMessageCount(int(maxMessages))
	}
	if maxCache != 0 {
		limiter.SetMaxCacheDuration(maxCache)
	}

	limiter.Start()
}
