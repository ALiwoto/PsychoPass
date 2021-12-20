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
		AssignedAt:     t.GetFormatedCreatedDate(),
	}

	if t.IsOwner() && !agent.IsOwner() {
		i.Permission = sv.Inspector
	} else {
		i.Permission = t.Permission
	}

	return i
}

func shouldSendNotFound(token *sv.Token, user *sv.User) bool {
	if token == nil {
		return !user.IsPastBanned() && user.IsCivilian()
	}

	return !token.IsRegistered() && !user.IsPastBanned()
}
