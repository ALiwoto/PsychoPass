package banHandlers

import (
	"strconv"

	entry "gitlab.com/Dank-del/SibylAPI-Go/sibyl/entryPoints"

	"github.com/gin-gonic/gin"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/hashing"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/database"
)

func AddBanHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "userId", "id", "user-id")
	banReason := utils.GetParam(c, "reason", "banReason", "ban-reason")
	banMsg := utils.GetParam(c, "message", "msg", "banMsg", "ban-msg")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginAddBan)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginAddBan)
		return
	}

	if d.CanBan() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			entry.SendInvalidUserIdError(c, OriginAddBan)
			return
		}

		by := hashing.GetIdFromToken(token)
		database.AddBan(id, by, banReason, banMsg)
		entry.SendResult(c, MessageBanned)
		return
	} else {
		entry.SendPermissionDenied(c, OriginAddBan)
		return
	}
}

func RemoveBanHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userId", "id")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginRemoveBan)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginRemoveBan)
		return
	}

	if d.CanBan() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			entry.SendInvalidUserIdError(c, OriginRemoveBan)
			return
		}

		u, _ := database.GetUserFromId(id)
		if u == nil {
			entry.SendUserNotFoundError(c, OriginRemoveBan)
			return
		}

		database.RemoveUserBan(u)

		entry.SendResult(c, MessageUnbanned)
		return
	} else {
		entry.SendPermissionDenied(c, OriginRemoveBan)
		return
	}
}
