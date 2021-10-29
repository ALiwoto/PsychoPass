package utils

import (
	"io/ioutil"
	"strings"

	sv "github.com/AnimeKaizoku/sibylapi-go/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/utils/hashing"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/database"
	"github.com/gin-gonic/gin"
)

// CreateToken creates a token for the given user.
func CreateToken(id int64, permission sv.UserPermission) (*sv.Token, error) {
	h := hashing.GetUserToken(id)
	data := &sv.Token{
		Permission: permission,
		UserId:     id,
		Hash:       h,
	}

	database.NewToken(data)
	return data, nil
}

// GetParam returns the value of the param with the given key.
// You can pass multiple keys to this function; it will check them
// in order until it finds a non-empty value.
// If no value is found, it returns an empty string.
func GetParam(c *gin.Context, key ...string) string {
	// prevent nil pointer dereference.
	if len(key) == 0 || c == nil {
		return ""
	}
	var result string
	for _, k := range key {
		result = strings.TrimSpace(getParam(c, k))
		if len(result) > 0 {
			return result
		}
	}
	return result
}

// getParam returns the value of the param with the given key.
// If the key is not found, it returns an empty string.
// This function will first check the key in header, then url query.
// Internal usage only; as it doesn't check for the passed parameters.
func getParam(c *gin.Context, key string) string {
	v := c.GetHeader(key)
	if len(v) == 0 {
		v = c.Request.URL.Query().Get(key)
	}
	return v
}

// ReadFile reads a file and returns its content.
// it returns an empty string if there is any problem in reading
// the file. If you want to do error handling, consider not using
// this function.
func ReadFile(path string) string {
	b, _ := ioutil.ReadFile(path)
	if len(b) == 0 {
		return ""
	}
	return string(b)
}
