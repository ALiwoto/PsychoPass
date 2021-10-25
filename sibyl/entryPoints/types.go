package entryPoints

import (
	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
)

type CreateTokenResponse struct {
	Token   *sv.Token         `json:"token"`
	Success bool              `json:"success"`
	Err     *sv.EndpointError `json:"error"`
}

type UserResponse struct {
	User    *sv.User          `json:"user"`
	Success bool              `json:"success"`
	Err     *sv.EndpointError `json:"error"`
}
