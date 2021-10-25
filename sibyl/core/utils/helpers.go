package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/hashing"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/database"
)

func CreateToken(id int64, permission sv.UserPermission) (*sv.Token, error) {
	h := hashing.GetUserToken(id)
	data := &sv.Token{
		Permission: permission,
		Id:         id,
		Hash:       h,
	}

	database.NewToken(data)
	return data, nil
}

func GetParam(c *gin.Context, key ...string) string {
	var result string
	for _, k := range key {
		result = strings.TrimSpace(getParam(c, k))
		if len(result) > 0 {
			return result
		}
	}
	return result
}

func getParam(c *gin.Context, key string) string {
	v := c.GetHeader(key)
	if len(v) == 0 {
		v = c.Request.URL.Query().Get(key)
	}
	return v
}
