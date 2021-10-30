package tokenHandlers

import (
	"strconv"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	entry "github.com/AnimeKaizoku/PsychoPass/sibyl/entryPoints"

	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/hashing"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
	"github.com/gin-gonic/gin"
)

// CreateTokenHandler function will create a new token for the specified
// user. if user already have a token in the db, it will just return that
// token.
func CreateTokenHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userid", "id")
	permInt, _ := strconv.Atoi(utils.GetParam(c, "perm", "permission"))
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginCreateToken)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginCreateToken)
		return
	}

	go database.UpdateTokenLastUsage(d)
	if d.CanCreateToken() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil || id == 0 {
			entry.SendInvalidUserIdError(c, OriginCreateToken)
			return
		}

		perm := sv.UserPermission(permInt)
		if !perm.IsValid() || perm.IsOwner() {
			entry.SendInvalidPermError(c, OriginCreateToken)
			return
		}

		u, _ := database.GetTokenFromId(id)
		if u != nil {
			if u.Permission != sv.UserPermission(perm) {
				database.UpdateTokenPermission(u, sv.UserPermission(perm))
			}
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

// ChangeTokenPermHandler function will change the permission of the specified
// token of the target user.
// users should have enough access to change the permission of a token.
// they need to pass the user-id of the target user and the new permission.
func ChangeTokenPermHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userid", "id")
	permInt, _ := strconv.Atoi(utils.GetParam(c, "perm", "permission"))
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginCreateToken)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginCreateToken)
		return
	}

	go database.UpdateTokenLastUsage(d)
	if d.CanCreateToken() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil || id == 0 {
			entry.SendInvalidUserIdError(c, OriginCreateToken)
			return
		}

		perm := sv.UserPermission(permInt)
		if !perm.IsValid() || perm.IsOwner() {
			entry.SendInvalidPermError(c, OriginCreateToken)
			return
		}

		u, _ := database.GetTokenFromId(id)
		if u != nil {
			if u.Permission != sv.UserPermission(perm) {
				database.UpdateTokenPermission(u, sv.UserPermission(perm))
			} else {
				entry.SendResult(c, MessagePermSame+u.GetTitleStringPermission())
				return
			}

			entry.SendResult(c, MessagePermChanged+u.GetTitleStringPermission())
			return
		}

		entry.SendUserNotFoundError(c, OriginCreateToken)
	} else {
		entry.SendPermissionDenied(c, OriginCreateToken)
	}
}

// RevokeTokenHandler function will revoke the specified token.
// you should pass the user-id of your target.
func RevokeTokenHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userid", "id")
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

// GetTokenHandler function will revoke the specified token.
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
