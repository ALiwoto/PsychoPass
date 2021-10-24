package routes

import "gitlab.com/Dank-del/SibylAPI-Go/database"

type CreateTokenResponse struct {
	Token   *database.Token `json:"token"`
	Success bool            `json:"success"`
	Err     string          `json:"err"`
}
