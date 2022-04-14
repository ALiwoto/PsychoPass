package reportHandlers

import (
	"encoding/json"
	"strconv"

	entry "github.com/MinistryOfWelfare/PsychoPass/sibyl/entryPoints"

	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/hashing"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
	"github.com/gin-gonic/gin"
)

func ReportUserHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "userId", "id", "user-id")
	reason := utils.GetParam(c, "reason", "reportReason", "report-reason")
	msg := utils.GetParam(c, "message", "msg", "reportMsg", "report-msg")
	msgLink := utils.GetParam(c, "src", "source", "report-src",
		"message-src", "msg-src")
	targetType := utils.GetEntityType(c)

	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginReport)
		return
	}

	agent, err := database.GetTokenFromString(token)
	if err != nil || agent == nil {
		entry.SendInvalidTokenError(c, OriginReport)
		return
	}

	if !agent.CanReport() {
		entry.SendPermissionDenied(c, OriginReport)
		return
	}

	by := hashing.GetIdFromToken(token)
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || sv.IsInvalidID(id) {
		entry.SendInvalidUserIdError(c, OriginReport)
		return
	}

	if sv.IsForbiddenID(id) {
		entry.SendPermissionDenied(c, OriginReport)
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
		if !u.CanBeReported(agent.Permission) {
			entry.SendCannotBeReportedError(c, OriginReport)
			return
		}
	}

	if sv.SendReportHandler != nil {
		r := sv.NewReport(
			reason,
			msg,
			msgLink,
			id,
			by,
			agent.Permission,
			targetType,
		)
		database.AddScan(r)
		go sv.SendReportHandler(r)
	}

	entry.SendResult(c, MessageReported)
}

func MultiReportHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")

	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginMultiScan)
		return
	}

	agent, err := database.GetTokenFromString(token)
	if err != nil || agent == nil {
		entry.SendInvalidTokenError(c, OriginMultiScan)
		return
	}

	if !agent.CanBan() {
		entry.SendPermissionDenied(c, OriginMultiScan)
		return
	}

	var rawData []byte
	multiScanData := new(sv.MultiScanRawData)

	rawData, err = c.GetRawData()
	if err != nil || len(rawData) < 2 {
		entry.SendNoDataError(c, OriginMultiScan)
		return
	}

	err = json.Unmarshal(rawData, multiScanData)
	if err != nil {
		entry.SendBadDataError(c, OriginMultiScan)
		return
	}

	if multiScanData != nil && len(multiScanData.Users) > 0 {
		if len(multiScanData.Users) > MaxMultiUsers {
			entry.SendTooManyError(c, OriginMultiScan, MaxMultiUsers)
			return
		}
		if sv.SendMultiReportHandler != nil {
			multiScanData.ReporterPermission = agent.Permission
			multiScanData.ReporterId = agent.UserId
			// prevent from spawning new goroutine if there is no handler
			go applyMultiScan(multiScanData)
		}
	}

	entry.SendResult(c, MessageApplyingMultiScan)
}
