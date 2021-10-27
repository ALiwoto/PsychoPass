package tokenHandlers

import (
	"strconv"

	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/logging"
	entry "gitlab.com/Dank-del/SibylAPI-Go/sibyl/entryPoints"

	"github.com/gin-gonic/gin"
	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/hashing"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/database"
)

// CreateToken function will create a new token for the specified
// user. if user already have a token in the db, it will just return that
// token.
func CreateTokenHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "id")
	perm, _ := strconv.Atoi(utils.GetParam(c, "perm", "permission"))
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginCreateToken)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginCreateToken)
		return
	}

	database.UpdateTokenLastUsage(d)
	if d.CanCreateToken() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil || id == 0 {
			entry.SendInvalidUserIdError(c, OriginCreateToken)
			return
		}

		u, _ := database.GetTokenFromId(id)
		if u != nil {
			entry.SendResult(c, u)
			return
		}

		t, err := utils.CreateToken(id, sv.UserPermission(perm))
		if err != nil {
			// this error is supposed to be unexpected;
			// in our tests, I couldn't see any case where we reached here,
			// but we should log it just in case.
			entry.SendInternalServerError(c, OriginCreateToken)
			logging.UnexpectedError(err)
			return
		}

		entry.SendResult(c, t)
	} else {
		entry.SendPermissionDenied(c, OriginCreateToken)
	}
}

// RevokeToken function will revoke the specified token.
// you should pass the user-id of your target.
func RevokeTokenHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "id")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginRevokeToken)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginRevokeToken)
		return
	}

	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || id == 0 {
		entry.SendInvalidUserIdError(c, OriginRevokeToken)
		return
	}

	u, _ := database.GetTokenFromId(id)
	if u == nil {
		entry.SendUserNotFoundError(c, OriginRevokeToken)
		return
	}

	if d.CanRevokeToken() || token == u.Hash {
		database.RevokeTokenHash(u, hashing.GetUserToken(id))
		entry.SendResult(c, u)
		return
	} else {
		entry.SendPermissionDenied(c, OriginRevokeToken)
	}
}

// GetToken function will revoke the specified token.
func GetTokenHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userId", "id")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginGetToken)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginGetToken)
		return
	}

	if d.CanGetToken() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil || id == 0 {
			entry.SendInvalidUserIdError(c, OriginGetToken)
			return
		}

		u, _ := database.GetTokenFromId(id)
		if u == nil {
			entry.SendUserNotFoundError(c, OriginGetToken)
			return
		}

		entry.SendResult(c, u)
		return
	} else {
		entry.SendPermissionDenied(c, OriginGetToken)
	}
}
