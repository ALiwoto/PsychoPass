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

	if t.IsOwner() {
		if agent.IsOwner() {
			i.Permission = t.Permission
		} else {
			i.Permission = sv.Inspector
		}
	}

	return i
}
