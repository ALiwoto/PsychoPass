// SRC: https://gist.github.com/nicklaw5/9d2d76b04d345152364d9b8cb4b554e9

package hashing

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var characterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

func GetToken(size, id int64) string {
	return FormatToken(RandomString(size), id)
}

func FormatToken(hash string, id int64) string {
	return strconv.FormatInt(id, 10) + ":" + hash
}

func GetUserToken(id int64) string {
	return GetToken(sibylConfig.GetMaxHashSize(), id)
}

func GetIdFromToken(value string) int64 {
	if !strings.Contains(value, ":") {
		return 0
	}

	id, _ := strconv.ParseInt(strings.Split(value, ":")[0], 10, 64)
	return id
}

func GetIdAndHashFromToken(value string) (int64, string) {
	if !strings.Contains(value, ":") {
		return 0, ""
	}

	strs := strings.Split(value, ":")
	id, _ := strconv.ParseInt(strs[0], 10, 64)
	return id, strs[1]
}

// RandomString generates a random string of n length
func RandomString(n int64) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
