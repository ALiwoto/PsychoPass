package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/Dank-del/SibylAPI-Go/database"
	"gitlab.com/Dank-del/SibylAPI-Go/hashing"
	"gitlab.com/Dank-del/SibylAPI-Go/server"
)

func CreateToken(c *gin.Context) {
	key := c.GetHeader("Sibyl-Token")
	userId := c.GetHeader("User-ID")
	d, err := database.GetFromToken(key)
	if err != nil {
		c.JSON(http.StatusBadGateway, &CreateTokenResponse{
			Err:     err.Error(),
			Success: false,
		})
	}
	if d.IsAdmin() {
		h := hashing.NewSHA1Hash(10)
		var id int64
		_, err := fmt.Sscan(userId, &id)
		if err != nil {
			c.JSON(http.StatusBadGateway, &CreateTokenResponse{
				Err:     err.Error(),
				Success: false,
			})
		}
		token := &database.Token{Permission: server.AdminParam, Hash: h, UserID: id}

		c.JSON(http.StatusOK,
			CreateTokenResponse{
				Token:   token,
				Success: true,
			})
	} else {
		c.JSON(http.StatusBadRequest, &CreateTokenResponse{
			Err:     "Permission Denied",
			Success: false,
		})
	}
}
