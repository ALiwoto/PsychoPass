package banHandlers

import (
	"strconv"
	"time"

	entry "github.com/AnimeKaizoku/PsychoPass/sibyl/entryPoints"

	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/hashing"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
	"github.com/gin-gonic/gin"
)

func AddBanHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "userId", "id", "user-id")
	banReason := utils.GetParam(c, "reason", "banReason", "ban-reason")
	banMsg := utils.GetParam(c, "message", "msg", "banMsg", "ban-msg")
	srcUrl := utils.GetParam(c, "srcUrl", "source",
		"source-url", "ban-src", "src")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginAddBan)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginAddBan)
		return
	}

	if d.CanBan() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil || sv.IsInvalidID(id) {
			entry.SendInvalidUserIdError(c, OriginAddBan)
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
		if u != nil && err == nil && u.Banned {
			if u.Reason == banReason && u.Message == banMsg &&
				u.BanSourceUrl == srcUrl {
				entry.SendUserAlreadyBannedError(c, OriginAddBan)
				return
			}

			// make a copy of the current struct value.
			pre := *u
			by := hashing.GetIdFromToken(token)
			u.BannedBy = by
			u.Message = banMsg
			u.Date = time.Now()
			u.BanSourceUrl = srcUrl
			u.SetAsBanReason(banReason)
			u.IncreaseCrimeCoefficientAuto()
			database.UpdateBanparameter(u)
			entry.SendResult(c, &BanResult{
				PreviousBan: &pre,
				CurrentBan:  u,
			})
			return
		}

		u = database.AddBan(id, by, banReason, banMsg, srcUrl)
		entry.SendResult(c, &BanResult{
			CurrentBan: u,
		})
		return
	} else {
		entry.SendPermissionDenied(c, OriginAddBan)
		return
	}
}

func RemoveBanHandler(c *gin.Context) {
	token := utils.GetParam(c, "token", "hash")
	userId := utils.GetParam(c, "user-id", "userId", "id")
	if len(token) == 0 {
		entry.SendNoTokenError(c, OriginRemoveBan)
		return
	}

	d, err := database.GetTokenFromString(token)
	if err != nil || d == nil {
		entry.SendInvalidTokenError(c, OriginRemoveBan)
		return
	}

	if d.CanBan() {
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil || sv.IsInvalidID(id) {
			entry.SendInvalidUserIdError(c, OriginRemoveBan)
			return
		}

		u, _ := database.GetUserFromId(id)
		if u == nil {
			entry.SendUserNotFoundError(c, OriginRemoveBan)
			return
		}

		if !u.Banned {
			entry.SendUserNotBannedError(c, OriginRemoveBan)
			return
		}

		database.RemoveUserBan(u)
		entry.SendResult(c, MessageUnbanned)
		return
	} else {
		entry.SendPermissionDenied(c, OriginRemoveBan)
		return
	}
}
