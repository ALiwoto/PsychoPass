/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package infoHandlers

import (
	"strconv"

	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
	entry "github.com/MinistryOfWelfare/PsychoPass/sibyl/entryPoints"
	"github.com/gin-gonic/gin"
)

func GetInfoHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userId", "id")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginGetInfo)
		return
	}

	agent, err := database.GetTokenFromString(token)
	if err != nil || agent == nil {
		entry.SendInvalidTokenError(c, OriginGetInfo)
		return
	}

	targetId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || sv.IsInvalidID(targetId) {
		entry.SendInvalidUserIdError(c, OriginGetInfo)
		return
	}

	if sv.IsForbiddenID(targetId) {
		entry.SendPermissionDenied(c, OriginGetInfo)
		return
	}

	target, _ := database.GetUserFromId(targetId)
	if target == nil {
		entry.SendUserNotFoundError(c, OriginGetInfo)
		return
	}

	targetToken, err := database.GetTokenFromId(targetId)
	if err == nil && targetToken != nil {
		if !target.IsCCValid(targetToken) {
			database.UpdateUserCrimeCoefficientByPerm(target, targetToken.Permission)
		}
	}

	if shouldSendNotFound(targetToken, target) {
		entry.SendUserNotFoundError(c, OriginGetInfo)
		return
	}

	entry.SendResult(c, target)
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
		entry.SendInternalServerError(c, OriginGetStats)
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

func GeneralInfoHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userId", "id")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginGeneralInfo)
		return
	}

	requesterToken, err := database.GetTokenFromString(token)
	if err != nil || requesterToken == nil {
		entry.SendInvalidTokenError(c, OriginGeneralInfo)
		return
	}

	if !requesterToken.CanGetGeneralInfo() {
		entry.SendPermissionDenied(c, OriginGeneralInfo)
		return
	}

	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || sv.IsInvalidID(id) {
		entry.SendInvalidUserIdError(c, OriginGeneralInfo)
		return
	}

	if sv.IsForbiddenID(id) {
		entry.SendPermissionDenied(c, OriginGeneralInfo)
		return
	}

	targetToken, err := database.GetTokenFromId(id)
	if targetToken == nil || err != nil {
		entry.SendUserNotFoundError(c, OriginGeneralInfo)
		return
	}

	if !targetToken.IsRegistered() {
		entry.SendUserNotRegisteredError(c, OriginGeneralInfo)
		return
	}

	entry.SendResult(c, toGeneralInfoResult(targetToken, requesterToken))
}
