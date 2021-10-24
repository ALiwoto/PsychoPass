package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/Dank-del/SibylAPI-Go/database"
	"gitlab.com/Dank-del/SibylAPI-Go/hashing"
	"gitlab.com/Dank-del/SibylAPI-Go/server"
	"net/http"
)

func CreateToken(c *gin.Context) {
	key := c.GetHeader("Sibyl-Token")
	userId := c.GetHeader("User-ID")
	d, err := database.GetFromToken(key)
	if err != nil {
		data := &CreateTokenResponse{Err: err.Error(), Success: false}
		e, _ := json.Marshal(data)
		c.JSON(http.StatusBadGateway, e)
	}
	if d.IsAdmin() {
		h := hashing.NewSHA1Hash(10)
		var id int64
		_, err := fmt.Sscan(userId, &id)
		if err != nil {
			data := &CreateTokenResponse{Err: err.Error(), Success: false}
			e, _ := json.Marshal(data)
			c.JSON(http.StatusBadGateway, e)
		}
		token := database.Token{Permission: server.AdminParam, Hash: h, UserID: id}
		data := CreateTokenResponse{Token: token, Success: true}
		e, _ := json.Marshal(data)
		c.JSON(http.StatusOK, e)
	} else {
		data := &CreateTokenResponse{Err: "Permission Denied", Success: false}
		e, _ := json.Marshal(data)
		c.JSON(http.StatusBadRequest, e)
	}
}
