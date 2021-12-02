package banHandlers

import sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"

type BanResult struct {
	PreviousBan *sv.User `json:"previous_ban"`
	CurrentBan  *sv.User `json:"current_ban"`
}
