package infoHandlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/database"
	entry "gitlab.com/Dank-del/SibylAPI-Go/sibyl/entryPoints"
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
