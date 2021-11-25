package infoHandlers

import (
	"time"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
)

func toGeneralInfoResult(t *sibylValues.Token) *GeneralInfoResult {
	return &GeneralInfoResult{
		UserId:         t.UserId,
		Division:       t.DivisionNum,
		AssignedBy:     t.AssignedBy,
		AssignedReason: t.AssignedReason,
		AssignedAt:     t.CreatedAt.Format(time.RFC3339),
		Permission:     t.Permission,
	}
}
