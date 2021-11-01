package reportHandlers

import (
	"strconv"

	entry "github.com/AnimeKaizoku/PsychoPass/sibyl/entryPoints"

	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/hashing"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
	"github.com/gin-gonic/gin"
)

func ReportUserHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "userId", "id", "user-id")
	reason := utils.GetParam(c, "reason", "reportReason", "report-reason")
	msg := utils.GetParam(c, "message", "msg", "reportMsg", "report-msg")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginReport)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginReport)
		return
	}

	if d.CanReport() {
		by := hashing.GetIdFromToken(token)
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil || sv.IsInvalidID(id) {
			entry.SendInvalidUserIdError(c, OriginReport)
			return
		}

		if by == id {
			entry.SendCannotReportYourselfError(c, OriginReport)
			return
		} else if sv.IsInvalidID(id) {
			entry.SendCannotBeReportedError(c, OriginReport)
			return
		}

		if len(reason) == 0 {
			entry.SendNoReasonError(c, OriginReport)
			return
		}

		u, err := database.GetTokenFromId(id)
		if err == nil && u != nil {
			if !u.CanBeReported() {
				entry.SendCannotBeReportedError(c, OriginReport)
				return
			}
		}

		if sv.SendReportHandler != nil {
			r := sv.NewReport(reason, msg, id, by, d.Permission)
			go sv.SendReportHandler(r)
		}

		entry.SendResult(c, MessageReported)
		return
	} else {
		entry.SendPermissionDenied(c, OriginReport)
		return
	}
}
