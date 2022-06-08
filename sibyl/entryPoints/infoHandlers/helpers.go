/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package infoHandlers

import (
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

func toGeneralInfoResult(targetToken, agent *sv.Token) *GeneralInfoResult {
	i := &GeneralInfoResult{
		UserId:             targetToken.UserId,
		Division:           targetToken.DivisionNum,
		AssignedBy:         targetToken.AssignedBy,
		AssignedReason:     targetToken.AssignedReason,
		AssignedAt:         targetToken.GetFormattedCreatedDate(),
		ApprovedScansCount: targetToken.AcceptedReports,
		RejectedScansCount: targetToken.DeniedReports,
		ForcedScansCount:   targetToken.ForcedCount,
		RevertCount:        targetToken.RevertCount,
	}

	if targetToken.IsOwner() && !agent.IsOwner() {
		i.Permission = sv.Inspector
	} else {
		i.Permission = targetToken.Permission
	}

	return i
}

func shouldSendNotFound(token *sv.Token, user *sv.User) bool {
	return !user.IsPastBanned() && user.IsCivilian() && !token.IsRegistered()
}
