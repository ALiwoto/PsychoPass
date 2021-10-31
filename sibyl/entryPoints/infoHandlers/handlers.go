package infoHandlers

import (
	"strconv"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"
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
	database.FetchStat()
	entry.SendResult(c, u)
}

func GetAllBansHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginGetInfo)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginGetInfo)
		return
	}

	if !d.CanGetToken() {
		entry.SendPermissionDenied(c, OriginGetInfo)
		return
	}

	bans, err := database.GetAllBannedUsers()
	if err != nil {
		entry.SendInternalServerError(c, OriginGetInfo)
		return
	}

	entry.SendResult(c, &GetBansResult{
		Users: bans,
	})
}

func CheckTokenHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginGetInfo)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendResult(c, false)
		return
	}

	entry.SendResult(c, true)
}
