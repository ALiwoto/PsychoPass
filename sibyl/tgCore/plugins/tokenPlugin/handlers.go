package tokenPlugin

import (
	"strconv"

	"github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/ALiwoto/mdparser/mdparser"
	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/hashing"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/database"
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
	md.AppendBoldThis("\n • User").AppendNormalThis(": ")
	md.AppendMentionThis(user.FirstName, user.Id)
	md.AppendBoldThis("\n • ID").AppendNormalThis(": ")
	md.AppendMonoThis(strconv.FormatInt(user.Id, 10))
	md.AppendBoldThis("\n • Status").AppendNormalThis(": ")
	md.AppendMonoThis(t.GetTitleStringPermission())
	md.AppendBoldThis("\n\nToken").AppendNormalThis(":\n")
	md.AppendMonoThis(t.Hash)
	md.AppendNormalThis("\n\nPlease don't share your token with anyone else!")

	kb.InlineKeyboard[2][0].Text = "Revoke API token"
	kb.InlineKeyboard[2][0].CallbackData = RevokeTokenCbValue

	_, _ = ctx.EffectiveMessage.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
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

	md := mdparser.GetNormal("\u200DYour token has been revoked successfully!")
	md.AppendNormalThis("\nInformation:")
	md.AppendBoldThis("\n • User").AppendNormalThis(": ")
	md.AppendMentionThis(user.FirstName, user.Id)
	md.AppendBoldThis("\n • ID").AppendNormalThis(": ")
	md.AppendMonoThis(strconv.FormatInt(user.Id, 10))
	md.AppendBoldThis("\n • Status").AppendNormalThis(": ")
	md.AppendMonoThis(t.GetTitleStringPermission())
	md.AppendBoldThis("\n\nToken").AppendNormalThis(":\n")
	md.AppendMonoThis(t.Hash)
	md.AppendNormalThis("\n\nPlease don't share your token with anyone else!")

	kb.InlineKeyboard[2][0].Text = "Close"
	kb.InlineKeyboard[2][0].CallbackData = "ap_close" // let start package handle close button xD

	_, _ = ctx.EffectiveMessage.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
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

	md := mdparser.GetNormal("Hi ").AppendMentionThis(user.FirstName, user.Id)
	md.AppendNormalThis(" !\nHere is your new token:\n")
	md.AppendMonoThis(t.Hash).AppendNormalThis("\n\n")
	md.AppendBoldThis("Please don't share this token with anyone!")
	if t.HasRole() {
		md.AppendItalicThis("You are a valid").AppendNormalThis(" ")
		md.AppendItalicThis(t.GetStringPermission()).AppendNormal(".")
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

	args := strongStringGo.Split(msg.Text, " ", "\n", "\t", ",", "-", "~")
	if len(args) < 2 {
		// show help.
		md := mdparser.GetNormal("Dear ").AppendMentionThis(user.FirstName, user.Id)
		md.AppendNormalThis(", this command lets you authorize dominator access for ")
		md.AppendHyperLinkThis("Sibyl", "http://t.me/SibylSystem")
		md.AppendNormalThis("\nRun command again in the following format")
		md.AppendBoldThis("\nYour options are:")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign inspector ID")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign enforcer ID")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign civilian ID")
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
		md.AppendMonoThis(args[1])
		md.AppendNormalThis("!\nHere is a list of possible permissions:")
		md.AppendNormalThis("\n- ").AppendMonoThis("inspector")
		md.AppendNormalThis("\n- ").AppendMonoThis("enforcer")
		md.AppendNormalThis("\n- ").AppendMonoThis("civilian")

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
		md.AppendBoldThis("\nYour options are:")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign inspector ID")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign enforcer ID")
		md.AppendNormalThis("\n- ").AppendMonoThis("/assign civilian ID")

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
			md.AppendMonoThis(args[2])
			md.AppendNormalThis("!\nPlease make sure the target's ID is valid.")

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
			md.AppendNormalThis("Please try another user ID.")
		} else if u.Permission == perm {
			md = mdparser.GetNormal("The user ")
			md.AppendMentionThis(strconv.FormatInt(targetId, 10), u.UserId)
			md.AppendNormalThis(" is already assigned as ")
			md.AppendMonoThis(perm.GetStringPermission()).AppendNormal(".")
		} else if !t.CanChangePermission(u.Permission, perm) {
			md = mdparser.GetNormal("Seems like you don't have enough privileges to ")
			md.AppendNormalThis("take this action.\nPlease try another user ID.")
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
			/*
				md = mdparser.GetNormal("The user ")
				md.AppendMentionThis(strconv.FormatInt(targetId, 10), u.UserId)
				md.AppendNormalThis(" has been assigned the permission ")
				md.AppendMonoThis(perm.GetStringPermission()).AppendNormal(".")
				md.AppendItalicThis("\n\nThe previous permission was ")
				md.AppendMonoThis(pre.GetStringPermission()).AppendNormal(".")
			*/
			invalid = false
		}
	} else {
		md = mdparser.GetUserMention(strconv.FormatInt(targetId, 10), targetId)
		md.AppendNormalThis(" needs to start me in PM to connect to Sibyl.")
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
		mdback.AppendMonoThis(perm.GetStringPermission())
		mdback.AppendNormalThis("!\n\nHere is your token:\n")
		mdback.AppendMonoThis(u.Hash).AppendNormalThis("\n\n")
		mdback.AppendBoldThis("Please don't share this token with anyone!")
		pm, err = b.SendMessage(targetId, mdback.ToString(), &gotgbot.SendMessageOpts{
			ParseMode:             sv.MarkDownV2,
			DisableWebPagePreview: true,
		})
		sendShouldStart := func() {
			md = mdparser.GetUserMention("\u200D", targetId)
			md.AppendMonoThis(strconv.FormatInt(targetId, 10))
			md.AppendNormalThis(" needs to start me in PM to connect to Sibyl.")
			if topMsg != nil {
				_, _ = topMsg.EditText(b, md.ToString(), &gotgbot.EditMessageTextOpts{
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
		database.UpdateTokenPermission(u, perm)
		assignValue := &AssignValue{
			targetChat: &pm.Chat,
			perm:       perm.GetStringPermission(),
			msg:        topMsg,
			targer:     targetUser,
			agent:      t,
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
