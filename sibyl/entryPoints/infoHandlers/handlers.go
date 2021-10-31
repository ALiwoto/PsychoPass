package infoHandlers

import (
	"strconv"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
	entry "github.com/AnimeKaizoku/PsychoPass/sibyl/entryPoints"
	"github.com/gin-gonic/gin"
)

func GetInfoHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userId", "id")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginGetInfo)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginGetInfo)
		return
	}

	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || id == 0 {
		entry.SendInvalidUserIdError(c, OriginGetInfo)
		return
	}

	u, _ := database.GetUserFromId(id)
	if u == nil {
		entry.SendUserNotFoundError(c, OriginGetInfo)
		return
	}

	entry.SendResult(c, u)
}

func GetAllBansHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginGetAllBans)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginGetAllBans)
		return
	}

	if !d.CanGetAllBans() {
		entry.SendPermissionDenied(c, OriginGetAllBans)
		return
	}

	bans, err := database.GetAllBannedUsers()
	if err != nil {
		entry.SendInternalServerError(c, OriginGetAllBans)
		return
	}

	entry.SendResult(c, &GetBansResult{
		Users: bans,
	})
}

func GetStatsHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginGetStats)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginGetStats)
		return
	}

	if !d.CanGetStats() {
		entry.SendPermissionDenied(c, OriginGetStats)
		return
	}

	stats, err := database.FetchStat()
	if err != nil {
		logging.UnexpectedError(err)
		entry.SendInternalServerError(c, OriginGetInfo)
		return
	}

	entry.SendResult(c, stats)
}

func CheckTokenHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginCheckToken)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendResult(c, false)
		return
	}

	entry.SendResult(c, true)
}
