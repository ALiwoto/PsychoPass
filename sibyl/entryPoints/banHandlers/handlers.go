package banHandlers

import (
	"encoding/json"
	"strconv"
	"time"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
	entry "github.com/MinistryOfWelfare/PsychoPass/sibyl/entryPoints"

	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/hashing"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
	"github.com/gin-gonic/gin"
)

func AddBanHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "userId", "id", "user-id")
	banReason := utils.GetParam(c, "reason", "banReason", "ban-reason")
	banMsg := utils.GetParam(c, "message", "msg", "banMsg", "ban-msg")
	srcUrl := utils.GetParam(c, "srcUrl", "source",
		"source-url", "ban-src", "src")
	srcGroup := utils.GetParam(c, "source-group", "src-group")
	targetType := utils.GetEntityType(c)

	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginAddBan)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginAddBan)
		return
	}

	if !d.CanBan() {
		entry.SendPermissionDenied(c, OriginAddBan)
		return
	}

	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || sv.IsInvalidID(id) {
		entry.SendInvalidUserIdError(c, OriginAddBan)
		return
	}

	if sv.IsForbiddenID(id) {
		entry.SendPermissionDenied(c, OriginAddBan)
		return
	}

	by := hashing.GetIdFromToken(token)
	if by == id {
		entry.SendCannotBanYourselfError(c, OriginAddBan)
		return
	}

	tu, err := database.GetTokenFromId(id)
	if err == nil && tu != nil {
		if !tu.CanBeBanned() {
			entry.SendCannotBeBannedError(c, OriginAddBan)
			return
		}
	}

	if len(banReason) == 0 {
		entry.SendNoReasonError(c, OriginAddBan)
		return
	}

	u, err := database.GetUserFromId(id)
	var count int
	if u != nil && err == nil {
		if u.Banned {
			// make a copy of the current struct value.
			pre := *u
			by := hashing.GetIdFromToken(token)
			if targetType != u.TargetType {
				// check both conditions; if they don't match, update the field.
				u.TargetType = targetType
			}
			u.BannedBy = by
			u.Message = banMsg
			u.Date = time.Now()
			u.BanSourceUrl = srcUrl
			u.SourceGroup = srcGroup
			u.SetAsBanReason(banReason)
			u.IncreaseCrimeCoefficientAuto()
			database.UpdateBanparameter(u)
			entry.SendResult(c, &BanResult{
				PreviousBan: &pre,
				CurrentBan:  u,
			})
			return
		}
		count = u.BanCount
	}

	info := &database.BanInfo{
		UserID:     id,
		Adder:      by,
		Reason:     banReason,
		SrcGroup:   srcGroup,
		Message:    banMsg,
		Src:        srcUrl,
		TargetType: targetType,
		Count:      count,
	}
	u = database.AddBan(info)

	entry.SendResult(c, &BanResult{
		CurrentBan: u,
	})
}

func MultiBanHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")

	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginMultiBan)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginMultiBan)
		return
	}

	if !d.CanBan() {
		entry.SendPermissionDenied(c, OriginMultiBan)
		return
	}

	by := hashing.GetIdFromToken(token)

	var rawData []byte
	multiBanData := new(sv.MultiBanRawData)

	rawData, err = c.GetRawData()
	if err != nil || len(rawData) < 2 {
		entry.SendNoDataError(c, OriginMultiBan)
		return
	}

	err = json.Unmarshal(rawData, multiBanData)
	if err != nil {
		entry.SendBadDataError(c, OriginMultiBan)
		return
	}

	if multiBanData != nil && len(multiBanData.Users) > 0 {
		if len(multiBanData.Users) > MaxMultiUsers {
			entry.SendTooManyError(c, OriginMultiBan)
			return
		}
		go applyMultiBan(multiBanData, by)
	}

	entry.SendResult(c, MessageApplyingMultiBan)
}

func RemoveBanHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userId", "id")
	clearHistory := ws.ToBool(utils.GetParam(c, "clear-history"))

	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginRemoveBan)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginRemoveBan)
		return
	}

	if !d.CanBan() {
		entry.SendPermissionDenied(c, OriginRemoveBan)
		return
	}

	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || sv.IsInvalidID(id) {
		entry.SendInvalidUserIdError(c, OriginRemoveBan)
		return
	}

	if sv.IsForbiddenID(id) {
		entry.SendPermissionDenied(c, OriginRemoveBan)
		return
	}

	u, _ := database.GetUserFromId(id)
	if u == nil {
		entry.SendUserNotFoundError(c, OriginRemoveBan)
		return
	}

	if !u.Banned && len(u.Reason) == 0 && len(u.BanFlags) == 0 {
		if clearHistory {
			database.ClearHistory(u)
			entry.SendResult(c, MessageHistoryCleared)
			return
		}
		entry.SendUserNotBannedError(c, OriginRemoveBan)
		return
	}

	database.RemoveUserBan(u, clearHistory)
	entry.SendResult(c, MessageUnbanned)
}

func MultiUnBanHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")

	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginMultiBan)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginMultiBan)
		return
	}

	if !d.CanBan() {
		entry.SendPermissionDenied(c, OriginMultiBan)
		return
	}

	by := hashing.GetIdFromToken(token)

	var rawData []byte
	multiUnBanData := new(sv.MultiUnBanRawData)

	rawData, err = c.GetRawData()
	if err != nil || len(rawData) < 2 {
		entry.SendNoDataError(c, OriginMultiBan)
		return
	}

	err = json.Unmarshal(rawData, multiUnBanData)
	if err != nil {
		entry.SendBadDataError(c, OriginMultiBan)
		return
	}

	if multiUnBanData != nil && len(multiUnBanData.Users) > 0 {
		if len(multiUnBanData.Users) > MaxMultiUsers {
			entry.SendTooManyError(c, OriginMultiBan)
			return
		}
		go applyMultiUnBan(multiUnBanData, by)
	}

	entry.SendResult(c, MessageApplyingMultiUnBan)
}
