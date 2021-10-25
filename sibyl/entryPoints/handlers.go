package entryPoints

import (
	"net/http"
	"strconv"

	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/logging"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/timeUtils"

	"github.com/gin-gonic/gin"
	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/hashing"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/database"
)

// CreateToken function will create a new token for the specified
// user. if user already have a token in the db, it will just return that
// token.
func CreateToken(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "id")
	perm, _ := strconv.Atoi(utils.GetParam(c, "perm", "permission"))
	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   ErrInvalidToken,
			Origin:    "createToken",
		}, http.StatusBadGateway)
		return
	}
	database.UpdateTokenLastUsage(d)
	if d.CanCreateToken() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil || id == 0 {
			SendErrorToken(c, &sv.EndpointError{
				ErrorCode: 502,
				Message:   ErrInvalidUserId,
				Origin:    "createToken",
			}, http.StatusBadGateway)
			return
		}
		u, _ := database.GetTokenFromId(id)
		if u != nil {
			c.JSON(http.StatusOK, &CreateTokenResponse{
				Token:   u,
				Success: true,
			})
			return
		}

		t, err := utils.CreateToken(id, sv.UserPermission(perm))
		if err != nil {
			// this error is supposed to be unexpected;
			// in our tests, I couldn't see any case where we reached here,
			// but we should log it just in case.
			SendErrorToken(c, &sv.EndpointError{
				ErrorCode: 500,
				Message:   ErrInternalServerError,
				Origin:    "CreateToken",
			}, http.StatusInternalServerError)
			logging.UnexpectedError(err)
			return
		}
		c.JSON(http.StatusOK, &CreateTokenResponse{
			Token:   t,
			Success: true,
		})
	} else {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   ErrPermissionDenied,
			Origin:    "CreateToken",
		}, http.StatusBadGateway)
	}
}

// RevokeToken function will revoke the specified token.
// you should pass the user-id of your target.
func RevokeToken(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "id")
	d, err := database.GetTokenFromString(token)
	if err != nil {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   ErrInvalidToken,
			Origin:    "RevokeToken",
		}, http.StatusBadGateway)
		return
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || id == 0 {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   ErrInvalidUserId,
			Origin:    "RevokeToken",
		}, http.StatusBadGateway)
		return
	}
	u, _ := database.GetTokenFromId(id)
	if u == nil {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 404,
			Message:   ErrUserNotFound,
			Origin:    "RevokeToken",
		}, http.StatusNotFound)
		return
	}
	if d.CanRevokeToken() || token == u.Hash {
		database.RevokeTokenHash(u, hashing.GetUserToken(id))
		c.JSON(http.StatusOK, &CreateTokenResponse{
			Token:   u,
			Success: true,
		})
		return
	} else {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   ErrPermissionDenied,
			Origin:    "RevokeToken",
		}, http.StatusBadGateway)
	}
}

// GetToken function will revoke the specified token.
func GetToken(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userId", "id")
	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   ErrInvalidToken,
			Origin:    "getToken",
		}, http.StatusBadGateway)
		return
	}
	if d.CanGetToken() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil || id == 0 {
			SendErrorToken(c, &sv.EndpointError{
				ErrorCode: 502,
				Message:   ErrInvalidUserId,
				Origin:    "getToken",
			}, http.StatusBadGateway)
			return
		}

		u, _ := database.GetTokenFromId(id)
		if u == nil {
			SendErrorToken(c, &sv.EndpointError{
				ErrorCode: 404,
				Message:   ErrUserNotFound,
				Origin:    "getToken",
			}, http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, &CreateTokenResponse{
			Token:   u,
			Success: true,
		})
		return
	} else {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   ErrPermissionDenied,
			Origin:    "getToken",
		}, http.StatusBadGateway)
	}
}

func AddBan(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "userId", "id", "user-id")
	banReason := utils.GetParam(c, "reason", "banreason", "ban-reason")
	banMsg := utils.GetParam(c, "message", "msg", "banmsg", "ban-msg")
	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		c.JSON(http.StatusBadGateway,
			&sv.SibylOperation{
				Success: false,
				Message: ErrInvalidToken,
				Time:    timeUtils.GenerateCurrentDateTime(),
			})
		return
	}
	if d.CanBan() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadGateway,
				&sv.SibylOperation{
					Success: false,
					Message: ErrInvalidUserId,
					Time:    timeUtils.GenerateCurrentDateTime(),
				})
			return
		}
		database.AddBan(id, banReason, banMsg)
		c.JSON(http.StatusOK, &sv.SibylOperation{
			Success: true,
			Message: MessageBanned,
			Time:    timeUtils.GenerateCurrentDateTime(),
		})
		return
	} else {
		c.JSON(http.StatusForbidden,
			&sv.SibylOperation{
				Success: false,
				Message: ErrPermissionDenied,
				Time:    timeUtils.GenerateCurrentDateTime(),
			})
		return
	}
}

func DeleteBan(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userId", "id")
	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		c.JSON(http.StatusBadGateway,
			&sv.SibylOperation{
				Success: false,
				Message: ErrInvalidToken,
				Time:    timeUtils.GenerateCurrentDateTime(),
			})
		return
	}
	if d.CanBan() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadGateway,
				&sv.SibylOperation{
					Success: false,
					Message: ErrInvalidUserId,
					Time:    timeUtils.GenerateCurrentDateTime(),
				})
			return
		}

		u, _ := database.GetUserFromId(id)
		if u == nil {
			c.JSON(http.StatusNotFound,
				&sv.SibylOperation{
					Success: false,
					Message: ErrUserNotFound,
					Time:    timeUtils.GenerateCurrentDateTime(),
				})
			return
		}

		database.DeleteUserBan(u)
		c.JSON(http.StatusOK, &sv.SibylOperation{
			Success: true,
			Message: MessageUnbanned,
			Time:    timeUtils.GenerateCurrentDateTime()})
		return
	} else {
		c.JSON(http.StatusForbidden,
			&sv.SibylOperation{
				Success: false,
				Message: ErrPermissionDenied,
				Time:    timeUtils.GenerateCurrentDateTime(),
			})
		return
	}
}

func GetInfo(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userId", "id")
	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		SendErrorUser(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   ErrInvalidToken,
			Origin:    "GetInfo",
		}, http.StatusBadGateway)
		return
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || id == 0 {
		SendErrorUser(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   ErrInvalidUserId,
			Origin:    "GetInfo",
		}, http.StatusBadGateway)
		return
	}
	u, _ := database.GetUserFromId(id)
	if u == nil {
		SendErrorUser(c, &sv.EndpointError{
			ErrorCode: 404,
			Message:   ErrUserNotFound,
			Origin:    "GetInfo",
		}, http.StatusBadGateway)
		return
	}

	c.JSON(http.StatusOK, &UserResponse{
		Success: true,
		User:    u,
	})
}
