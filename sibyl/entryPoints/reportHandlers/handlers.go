package reportHandlers

import (
	"strconv"

	entry "github.com/AnimeKaizoku/sibylapi-go/sibyl/entryPoints"

	sv "github.com/AnimeKaizoku/sibylapi-go/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/utils"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/utils/hashing"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/database"
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
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			entry.SendInvalidUserIdError(c, OriginReport)
			return
		}

		if len(reason) == 0 {
			entry.SendNoReasonError(c, OriginReport)
			return
		}

		by := hashing.GetIdFromToken(token)
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
