package banHandlers

import sv "github.com/AnimeKaizoku/sibylapi-go/sibyl/core/sibylValues"

type BanResult struct {
	PreviousBan *sv.User `json:"previous_ban"`
	CurrentBan  *sv.User `json:"current_ban"`
}
