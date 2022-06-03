/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package pollingHandlers

import (
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
	entry "github.com/MinistryOfWelfare/PsychoPass/sibyl/entryPoints"
	"github.com/gin-gonic/gin"
)

func GetUpdatesHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	timeout := utils.GetParam(c, "timeout", "time-out")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginGetUpdates)
		return
	}

	timeoutValue := sibylValues.DefaultPollingTimeout
	if timeout != "" {
		timeoutInt := ssg.ToInt64(timeout)
		if sibylValues.IsPollingTimeoutInvalid(timeoutInt) {
			timeoutValue = time.Duration(timeoutInt) * time.Second
		}
	}

	requesterToken, err := database.GetTokenFromString(token)
	if err != nil || requesterToken == nil {
		entry.SendInvalidTokenError(c, OriginGetUpdates)
		return
	}

	if !requesterToken.CanStartPolling() {
		entry.SendPermissionDenied(c, OriginGetUpdates)
		return
	}

	pValue := sibylValues.RegisterNewPollingValue(
		requesterToken.UserId,
		pollingNumGenerator.Next(),
		timeoutValue,
	)

	select {
	case <-pValue.Done():
		entry.SendResult(c, nil)
		// we got timed out here, so we have to unregister the registered pValue.
		sibylValues.UnregisterPollingValue(false, pValue)
	case pUpdate := <-pValue.GotUpdate():
		// there is no need to unregister the pValue here, everything is done on caller's side
		// (the one that broadcasted this update, we just have to send the response to the user and
		// we are done!)
		entry.SendResult(c, pUpdate)
	}
}
