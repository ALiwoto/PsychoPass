package tokenPlugin

import (
	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

type AssignValue struct {
	targetChat   *gotgbot.Chat
	perm         string
	permValue    sv.UserPermission
	msg          *gotgbot.Message
	targer       *sv.User
	agentProfile *gotgbot.User
	agent        *sv.Token
}
