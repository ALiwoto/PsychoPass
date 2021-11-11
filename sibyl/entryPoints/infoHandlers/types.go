package infoHandlers

import sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"

type InfoResult struct {
	UserId           string `json:"user_id"`
	IsBanned         bool   `json:"is_banned"`
	BanDate          string `json:"ban_date"`
	Reason           string `json:"reason"`
	Message          string `json:"message"`
	CrimeCoefficient int    `json:"crime_coefficient"`
}

type GetBansResult struct {
	Users []sv.User `json:"users"`
}
