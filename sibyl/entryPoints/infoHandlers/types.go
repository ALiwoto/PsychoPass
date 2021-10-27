package infoHandlers

type InfoResult struct {
	UserId          string `json:"user_id"`
	IsBanned        bool   `json:"is_banned"`
	BanDate         string `json:"ban_date"`
	Reason          string `json:"reason"`
	Message         string `json:"message"`
	CrimeCoeficient int    `json:"crime_coeficient"`
}
