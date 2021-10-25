package entryPoints

import (
	"fmt"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/timeUtils"
	"net/http"
	"strconv"

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
	perm, _ := strconv.Atoi(utils.GetParam(c, "perm"))
	d, err := database.GetFromToken(token)
	if err != nil {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   err.Error(),
			Origin:    "createToken",
		}, http.StatusBadGateway)
		return
	}
	database.UpdateTokenLastUsage(d)
	if d.CanCreateToken() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			SendErrorToken(c, &sv.EndpointError{
				ErrorCode: 502,
				Message:   "invalid user-id",
				Origin:    "createToken",
			}, http.StatusBadGateway)
			return
		}
		u, _ := database.GetFromId(id)
		if u != nil {
			c.JSON(http.StatusOK, &CreateTokenResponse{
				Token:   u,
				Success: true,
			})
			return
		}

		t, err := utils.CreateToken(id, sv.UserPermission(perm))
		if err != nil {
			SendErrorToken(c, &sv.EndpointError{
				ErrorCode: 502,
				Message:   err.Error(),
				Origin:    c.Request.URL.Path,
			}, http.StatusBadGateway)
			return
		}
		c.JSON(http.StatusOK, &CreateTokenResponse{
			Token:   t,
			Success: true,
		})
	} else {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   "Permission Denied",
			Origin:    "CreateToken",
		}, http.StatusBadGateway)
	}
}

// RevokeToken function will revoke the specified token.
// you should pass the user-id of your target.
func RevokeToken(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "id")
	d, err := database.GetFromToken(token)
	if err != nil {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   err.Error(),
			Origin:    "revokeToken",
		}, http.StatusBadGateway)
		return
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   "invalid user-id",
			Origin:    "revokeToken",
		}, http.StatusBadGateway)
		return
	}
	u, _ := database.GetFromId(id)
	if u == nil {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   "user not found",
			Origin:    "revokeToken",
		}, http.StatusBadGateway)
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
			Message:   "Permission Denied",
			Origin:    "revokeToken",
		}, http.StatusBadGateway)
	}
}

// GetToken function will revoke the specified token.
func GetToken(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "id")
	d, err := database.GetFromToken(token)
	if err != nil {
		SendErrorToken(c, &sv.EndpointError{
			ErrorCode: 502,
			Message:   err.Error(),
			Origin:    "getToken",
		}, http.StatusBadGateway)
		return
	}
	if d.CanGetToken() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			SendErrorToken(c, &sv.EndpointError{
				ErrorCode: 502,
				Message:   "invalid user-id",
				Origin:    "getToken",
			}, http.StatusBadGateway)
			return
		}
		u, _ := database.GetFromId(id)

		if u == nil {
			SendErrorToken(c, &sv.EndpointError{
				ErrorCode: 502,
				Message:   "user not found",
				Origin:    "getToken",
			}, http.StatusBadGateway)
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
			Message:   "Permission Denied",
			Origin:    "getToken",
		}, http.StatusBadGateway)
	}
}

func SendErrorToken(c *gin.Context, err *sv.EndpointError, status int) {
	c.JSON(status, &CreateTokenResponse{
		Err:     err,
		Success: false,
	})
}

func AddBan(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "userId", "id")
	banReason := utils.GetParam(c, "reason", "banreason")
	banMsg := utils.GetParam(c, "message", "msg", "banmsg")
	d, err := database.GetFromToken(token)
	if err != nil {
		c.JSON(http.StatusBadGateway,
			&sv.SibylOperation{
				Success: false,
				Message: fmt.Sprintf("User wasn't banned due to %s", err.Error()),
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
					Message: fmt.Sprintf("User wasn't banned due to %s", err.Error()),
					Time:    timeUtils.GenerateCurrentDateTime(),
				})
			return
		}
		database.AddBan(id, banReason, banMsg)
		c.JSON(http.StatusOK, &sv.SibylOperation{Success: true, Message: "User was banned", Time: timeUtils.GenerateCurrentDateTime()})
		return
	} else {
		c.JSON(http.StatusForbidden,
			&sv.SibylOperation{
				Success: false,
				Message: "User wasn't banned as you lack permissions",
				Time:    timeUtils.GenerateCurrentDateTime(),
			})
		return
	}
}

func DeleteBan(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "userId", "id")
	d, err := database.GetFromToken(token)
	if err != nil {
		c.JSON(http.StatusBadGateway,
			&sv.SibylOperation{
				Success: false,
				Message: "User wasn't unbanned due to an issue with the token",
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
					Message: "User wasn't unbanned due to failure with parsing of user id",
					Time:    timeUtils.GenerateCurrentDateTime(),
				})
			return
		}
		database.DeleteBan(id)
		c.JSON(http.StatusOK, &sv.SibylOperation{Success: true, Message: "User was unbanned", Time: timeUtils.GenerateCurrentDateTime()})
		return
	} else {
		c.JSON(http.StatusForbidden,
			&sv.SibylOperation{
				Success: false,
				Message: "User wasn't unbanned as you lack permissions",
				Time:    timeUtils.GenerateCurrentDateTime(),
			})
		return
	}
}
