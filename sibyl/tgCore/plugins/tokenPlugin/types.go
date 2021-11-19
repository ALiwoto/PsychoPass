package tokenPlugin

import (
	sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

type AssignValue struct {
	targetChat *gotgbot.Chat
	perm       string
	msg        *gotgbot.Message
	targer     *sv.User
	agent      *sv.Token
}
