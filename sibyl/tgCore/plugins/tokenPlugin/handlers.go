/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package tokenPlugin

import (
	"strconv"
	"strings"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/ssg/ssg"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/hashing"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func getTokenCallBackQuery(cq *gotgbot.CallbackQuery) bool {
	return cq.Data == GetTokenCbValue
}

func getTokenCallBackResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	var t *sv.Token
	var err error
	user := ctx.EffectiveUser
	kb := ctx.EffectiveMessage.ReplyMarkup
	if kb == nil || len(kb.InlineKeyboard) < 3 {
		// message doesn't have any reply markup; special situation which is
		// unlikely to happen, added this checker just in-case to prevent
		// from panic.
		return ext.EndGroups
	}

	t, err = database.GetTokenFromId(user.Id)
	if err != nil {
		logging.UnexpectedError(err)
		return ext.EndGroups
	}

	if t == nil {
		// should create a new token
		t, err = utils.CreateToken(user.Id, sv.NormalUser)
		if err != nil {
			logging.UnexpectedError(err)
			return ext.EndGroups
		}
	}

	md := mdparser.GetNormal("\u200DInformation:")
	md.Bold("\n • User").Normal(": ")
	md.Mention(user.FirstName, user.Id)
	md.Bold("\n • ID").Normal(": ")
	md.Mono(strconv.FormatInt(user.Id, 10))
	md.Bold("\n • Status").Normal(": ")
	md.Mono(t.GetTitleStringPermission())
	md.Bold("\n\nToken").Normal(":\n")
	md.Mono(t.Hash)
	md.Normal("\n\nPlease don't share your token with anyone else!")

	kb.InlineKeyboard[2][0].Text = "Revoke API token"
	kb.InlineKeyboard[2][0].CallbackData = RevokeTokenCbValue

	_, _, _ = ctx.EffectiveMessage.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:   sv.MarkDownV2,
		ReplyMarkup: *kb,
	})
	return ext.EndGroups
}

func revokeTokenCallBackQuery(cq *gotgbot.CallbackQuery) bool {
	return cq.Data == RevokeTokenCbValue
}

func revokeTokenCallBackResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	var t *sv.Token
	var err error
	user := ctx.EffectiveUser
	kb := ctx.EffectiveMessage.ReplyMarkup
	if kb == nil || len(kb.InlineKeyboard) < 3 {
		// message doesn't have any reply markup; special situation which is
		// unlikely to happen, added this checker just in-case to prevent
		// from panic.
		return ext.EndGroups
	}

	t, err = database.GetTokenFromId(user.Id)
	if err != nil {
		logging.UnexpectedError(err)
		return ext.EndGroups
	}

	if t == nil {
		// is user trying to invoke a token which doesn't even exist in the database?
		// seems impossible, unless someone who has direct access to database deleted it.
		// should create a new token
		t, err = utils.CreateToken(user.Id, sv.NormalUser)
		if err != nil {
			logging.UnexpectedError(err)
			return ext.EndGroups
		}
	} else {
		if !t.CanBeRevoked() {
			if ctx.CallbackQuery != nil {
				_, _ = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
					Text:      "You have revoked your token too many times!",
					ShowAlert: true,
					CacheTime: 5,
				})
			}
			return ext.EndGroups
		}
		_ = database.RevokeTokenHash(t, hashing.GetUserToken(user.Id))
	}

	md := mdparser.GetNormal("Your token has been revoked successfully!")
	md.Normal("\nInformation:")
	md.Bold("\n • User").Normal(": ")
	md.Mention(user.FirstName, user.Id)
	md.Bold("\n • ID").Normal(": ")
	md.Mono(strconv.FormatInt(user.Id, 10))
	md.Bold("\n • Status").Normal(": ")
	md.Mono(t.GetTitleStringPermission())
	md.Bold("\n\nToken").Normal(":\n")
	md.Mono(t.Hash)
	md.Normal("\n\nPlease don't share your token with anyone else!")

	kb.InlineKeyboard[2][0].Text = "Close"
	kb.InlineKeyboard[2][0].CallbackData = "ap_close" // let start package handle close button xD

	_, _, _ = ctx.EffectiveMessage.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
		ParseMode:   sv.MarkDownV2,
		ReplyMarkup: *kb,
	})
	return ext.EndGroups
}

// revokeHandler is the handler for the /revoke command.
// It will revoke the token of the user.
func revokeHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	t, err := database.GetTokenFromId(user.Id)
	if err != nil || t == nil {
		return ext.EndGroups
	}

	database.RevokeTokenHash(t, hashing.GetUserToken(user.Id))

	md := mdparser.GetNormal("Hi ").Mention(user.FirstName, user.Id)
	md.Normal(" !\nHere is your new token:\n")
	md.Mono(t.Hash).Normal("\n\n")
	md.Bold("Please don't share this token with anyone!")
	if t.HasRole() {
		md.Italic("You are a valid").Normal(" ")
		md.Italic(t.GetStringPermission()).AppendNormal(".")
	}

	b.SendMessage(user.Id, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode: sv.MarkDownV2,
	})

	return ext.EndGroups
}

// assignHandler is the handler for the /assign command.
// It will change the permission of the token of the user.
func assignHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	if sv.IsInvalidID(user.Id) {
		return ext.EndGroups
	}

	t, err := database.GetTokenFromId(user.Id)
	if err != nil || t == nil || !t.CanTryChangePermission(false) {
		return ext.EndGroups
	}

	args := ssg.Split(msg.Text, " ", "\n", "\t", ",", "-", "~")
	if len(args) < 2 {
		// show help.
		md := mdparser.GetNormal("Dear ").Mention(user.FirstName, user.Id)
		md.Normal(", this command lets you authorize dominator access for ")
		md.Link("Sibyl", "http://t.me/SibylSystem")
		md.Normal("\nRun command again in the following format")
		md.Bold("\nYour options are:")
		md.Normal("\n- ").Mono("/assign inspector ID")
		md.Normal("\n- ").Mono("/assign enforcer ID")
		md.Normal("\n- ").Mono("/assign civilian ID")
		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	}

	perm, err := sv.ConvertToPermission(args[1])
	if err != nil {
		md := mdparser.GetNormal("Invalid permission provided: ")
		md.Mono(args[1])
		md.Normal("!\nHere is a list of possible permissions:")
		md.Normal("\n- ").Mono("inspector")
		md.Normal("\n- ").Mono("enforcer")
		md.Normal("\n- ").Mono("civilian")

		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	} else if perm.IsOwner() {
		md := mdparser.GetNormal("This decision is of Sibyl to make!")
		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	}

	var replied = msg.ReplyToMessage != nil && msg.ReplyToMessage.From != nil
	var preReplied = replied // reserved value
	var isBot bool
	var targetId int64

	if len(args) < 3 && !replied {
		// show help.
		md := mdparser.GetNormal("You need to provide a user ID for this command.")
		md.Bold("\nYour options are:")
		md.Normal("\n- ").Mono("/assign inspector ID")
		md.Normal("\n- ").Mono("/assign enforcer ID")
		md.Normal("\n- ").Mono("/assign civilian ID")

		_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
		}

		return ext.EndGroups
	} else if len(args) > 2 && replied {
		// ignore replied message if the message itself contains the user ID.
		replied = false
	}

	if replied {
		targetId = msg.ReplyToMessage.From.Id
		isBot = msg.ReplyToMessage.From.IsBot
	} else {
		targetId, err = strconv.ParseInt(args[2], 10, 64)
	}

	if err != nil || targetId == 0 {
		if !replied && !preReplied {
			md := mdparser.GetNormal("Invalid ID provided: ")
			md.Mono(args[2])
			md.Normal("!\nPlease make sure the target's ID is valid.")

			_, err := msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
				ParseMode:                sv.MarkDownV2,
				AllowSendingWithoutReply: true,
				DisableWebPagePreview:    true,
			})
			if err != nil {
				logging.UnexpectedError(err)
			}

			return ext.EndGroups
		} else if preReplied {
			targetId = msg.ReplyToMessage.From.Id
			isBot = msg.ReplyToMessage.From.IsBot
		}
	}

	// WARNING: this value may be nil; always check before using it.
	var targetUser *sv.User
	// target id validation section
	if targetId == user.Id {
		md := mdparser.GetNormal("You can't change your own permissions.")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		return ext.EndGroups
	} else if targetId == b.Id {
		md := mdparser.GetNormal("You can't change my permissions.")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		return ext.EndGroups
	} else if isBot || sv.IsInvalidID(targetId) {
		md := mdparser.GetNormal("Dominator authorization is only for humans.")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		return ext.EndGroups
	}

	invalid := true
	var md mdparser.WMarkDown
	var topMsg *gotgbot.Message
	u, _ := database.GetTokenFromId(targetId)
	var isDemote bool
	if u != nil {
		isDemote = u.Permission > perm
		if u.IsOwner() {
			md = mdparser.GetNormal("This decision is of Sibyl to make.\n")
			md.Normal("Please try another user ID.")
		} else if u.Permission == perm {
			md = mdparser.GetNormal("The user ")
			md.Mention(strconv.FormatInt(targetId, 10), u.UserId)
			md.Normal(" is already assigned as ")
			md.Mono(perm.GetStringPermission()).AppendNormal(".")
		} else if !t.CanChangePermission(u.Permission, perm) {
			md = mdparser.GetNormal("Seems like you don't have enough privileges to ")
			md.Normal("take this action.\nPlease try another user ID.")
		} else {
			targetUser, err = database.GetUserFromId(targetId)
			if err == nil && targetUser != nil && targetUser.Banned {
				go showUserIsBanned(b, ctx, targetUser, perm.GetStringPermission(), replied)
				return ext.EndGroups
			}

			//pre := u.Permission
			mmd := mdparser.GetNormal("Running a cymatic scan....")
			topMsg, _ = msg.Reply(b, mmd.ToString(), &gotgbot.SendMessageOpts{
				ParseMode: sv.MarkDownV2,
			})
			if topMsg == nil {
				return ext.EndGroups
			}
			invalid = false
		}
	} else {
		md = mdparser.GetUserMention(strconv.FormatInt(targetId, 10), targetId)
		md.Normal(" needs to start me in PM to connect to Sibyl.")
		_, err = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			ReplyMarkup:              *startCymaticScanButton,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
		if err != nil {
			logging.UnexpectedError(err)
			return ext.EndGroups
		}

		return ext.EndGroups
	}

	if !invalid {
		var pm *gotgbot.Message
		mdback := mdparser.GetNormal("Your permission has been changed to ")
		mdback.Mono(perm.GetStringPermission())
		mdback.Normal("!\n\nHere is your token:\n")
		mdback.Mono(u.Hash).Normal("\n\n")
		mdback.Bold("Please don't share this token with anyone!")
		pm, err = b.SendMessage(targetId, mdback.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:             sv.MarkDownV2,
			DisableWebPagePreview: true,
		})
		sendShouldStart := func() {
			md = mdparser.GetUserMention("\u200D", targetId)
			md.Mono(strconv.FormatInt(targetId, 10))
			md.Normal(" needs to start me in PM to connect to Sibyl.")
			if topMsg != nil {
				_, _, _ = topMsg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
					ParseMode:             sv.MarkDownV2,
					ReplyMarkup:           *startCymaticScanButton,
					DisableWebPagePreview: true,
				})
			} else {
				_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
					ParseMode:                sv.MarkDownV2,
					AllowSendingWithoutReply: false,
					DisableWebPagePreview:    true,
				})
			}
		}
		if err != nil || pm == nil {
			if !isDemote {
				sendShouldStart()
				return ext.EndGroups
			}
			// if we are going to demote the user, simply ignore the pm
			// part and do your own thing.
			chat, err := b.GetChat(targetId)
			if err != nil || chat == nil {
				chat = &gotgbot.Chat{
					Id:        targetId,
					FirstName: strconv.FormatInt(targetId, 10),
				}
			}
			pm = &gotgbot.Message{
				Chat: *chat,
			}
		}

		if targetUser == nil {
			targetUser, err = database.GetUserFromId(targetId)
			if err != nil || targetUser == nil {
				// means user is not found in the database;
				// so insert it by force.
				if t.CanTryChangePermission(true) {
					targetUser = database.ForceInsert(targetId, perm)
				} else {
					targetUser = database.ForceInsert(targetId, sv.NormalUser)
				}
			} else {
				if t.CanTryChangePermission(true) {
					// as database.ForceInsert will already set the crime coefficient
					// of the user by their perm, we should use else here, so it prevents
					// from sending useless queries to database.
					database.UpdateUserCrimeCoefficientByPerm(targetUser, perm)
				}
			}
		}

		if t.CanTryChangePermission(true) {
			database.UpdateTokenPermission(u, perm)
		}

		assignValue := &AssignValue{
			targetChat:   &pm.Chat,
			perm:         perm.GetStringPermission(),
			permValue:    perm,
			msg:          topMsg,
			target:       targetUser,
			agentProfile: user,
			agent:        t,
			src:          utils.GetLink(ctx),
		}
		go showUserAssigned(b, ctx, assignValue)
		return ext.EndGroups
	}

	if md != nil {
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:                sv.MarkDownV2,
			AllowSendingWithoutReply: true,
			DisableWebPagePreview:    true,
		})
	}

	return ext.EndGroups
}

func assignCallBackQuery(cq *gotgbot.CallbackQuery) bool {
	return strings.HasPrefix(cq.Data, AssignCbPrefix)
}

func assignCallBackResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	tgUser := ctx.EffectiveUser
	token, err := database.GetTokenFromId(tgUser.Id)
	if err != nil || token == nil || !token.CanTryChangePermission(true) {
		_, _ = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "This is not for you...",
			ShowAlert: true,
			CacheTime: 5,
		})
		return ext.EndGroups
	}
	data := strings.TrimPrefix(ctx.CallbackQuery.Data, AssignCbPrefix)
	switch data {
	case RejectCbData:
		txt := ctx.EffectiveMessage.Text + "\n\nRequest has been rejected."
		_, _, _ = ctx.EffectiveMessage.EditText(b, txt, &gotgbot.EditMessageTextOpts{
			Entities:              ctx.EffectiveMessage.Entities,
			DisableWebPagePreview: true,
		})

		return ext.EndGroups
	case CloseCbData:
		_, _ = ctx.EffectiveMessage.Delete(b)
		return ext.EndGroups
	}

	myStrs := strings.Split(data, CbSep)
	if len(myStrs) != 3 {
		return ext.EndGroups
	}

	perm, err := sv.ConvertToPermission(myStrs[2])
	if err != nil {
		return ext.EndGroups
	}

	assignValue := toAssignValue(ctx.EffectiveMessage, perm)
	print(assignValue)

	return ext.EndGroups
}
