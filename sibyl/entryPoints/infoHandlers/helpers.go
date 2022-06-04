/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package infoHandlers

import (
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

func toGeneralInfoResult(t, agent *sv.Token) *GeneralInfoResult {
	i := &GeneralInfoResult{
		UserId:         t.UserId,
		Division:       t.DivisionNum,
		AssignedBy:     t.AssignedBy,
		AssignedReason: t.AssignedReason,
		AssignedAt:     t.GetFormattedCreatedDate(),
	}

	if t.IsOwner() && !agent.IsOwner() {
		i.Permission = sv.Inspector
	} else {
		i.Permission = t.Permission
	}

	return i
}

func shouldSendNotFound(token *sv.Token, user *sv.User) bool {
	return !user.IsPastBanned() && user.IsCivilian() && !token.IsRegistered()
}
