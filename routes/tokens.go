package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/Dank-del/SibylAPI-Go/Hashing"
	"gitlab.com/Dank-del/SibylAPI-Go/database"
	"gitlab.com/Dank-del/SibylAPI-Go/server"
	"net/http"
)

func CreateToken(c *gin.Context) {
   key := c.GetHeader("Sibyl-Token")
   userId := c.GetHeader("User-ID")
   d, err := database.GetFromToken(key)
   if err != nil {
	   data := &CreateTokenResponse{Err: err.Error(), Success: false}
	   c.JSON(http.StatusBadGateway, interface{}(json.Marshal(data)))
   }
   if d.IsAdmin() {
	   h := Hashing.NewSHA1Hash(10)
	   var id int64
	   _, err := fmt.Sscan(userId, &id)
	   if err != nil {
		   data := &CreateTokenResponse{Err: err.Error(), Success: false}
		   c.JSON(http.StatusBadGateway, interface{}(json.Marshal(data)))
	   }
	   token := database.Token{Permission: server.AdminParam, Hash: h, UserID: id}
	   data := CreateTokenResponse{Token: token, Success: true}
	   c.JSON(http.StatusOK, interface{}(json.Marshal(data)))
   } else {
	   data := &CreateTokenResponse{Err: "Permission Denied", Success: false}
	   c.JSON(http.StatusBadRequest, interface{}(json.Marshal(data)))
   }
}
