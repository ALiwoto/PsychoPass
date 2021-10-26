package entryPoints

import (
	"github.com/gin-gonic/gin"
	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
)

func SendErrorToken(c *gin.Context, err *sv.EndpointError, status int) {
	c.JSON(status, &CreateTokenResponse{
		Err:     err,
		Success: false,
	})
}

func SendErrorUser(c *gin.Context, err *sv.EndpointError, status int) {
	c.JSON(status, &UserResponse{
		Err:     err,
		Success: false,
	})
}
