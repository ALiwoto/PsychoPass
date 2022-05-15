/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package pollingHandlers

import (
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
	entry "github.com/MinistryOfWelfare/PsychoPass/sibyl/entryPoints"
	"github.com/gin-gonic/gin"
)

func StartPollingHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginStartPolling)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginStartPolling)
		return
	}

	if !d.CanStartPolling() {
		entry.SendPermissionDenied(c, OriginStartPolling)
		return
	}

	// TODO: generate and add config

	entry.SendResult(c, true)
}
