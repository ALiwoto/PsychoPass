package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/ALiwoto/mdparser/mdparser"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/hashing"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gin-gonic/gin"
)

// SafeReply will reply the message safely, if output is longer than 4096 character, it will
// send it as file; otherwise it will send it as text using mdparser (monospace).
func SafeReply(b *gotgbot.Bot, ctx *ext.Context, output string) error {
	msg := ctx.EffectiveMessage
	if len(output) < 4096 {
		_, err := msg.Reply(b, mdparser.GetMono(output).ToString(),
			&gotgbot.SendMessageOpts{ParseMode: "Markdownv2"})
		if err != nil {
			logging.Error("got an error when trying to send results: ", err)
			return err
		}
	} else {
		_, err := b.SendDocument(ctx.EffectiveChat.Id, []byte(output), &gotgbot.SendDocumentOpts{
			ReplyToMessageId: msg.MessageId,
		})
		if err != nil {
			logging.Error("got an error when trying to send document: ", err)
			return err
		}
	}

	return nil
}

// SafeReply will reply the message safely, if output is longer than 4096 character, it will
// send it as file; otherwise it will send it as plain text without using any formating.
func SafeReplyNoFormat(b *gotgbot.Bot, ctx *ext.Context, output string) error {
	msg := ctx.EffectiveMessage
	if len(output) < 4096 {
		_, err := msg.Reply(b, output,
			&gotgbot.SendMessageOpts{ParseMode: "Markdownv2"})
		if err != nil {
			logging.Error("got an error when trying to send results: ", err)
			return err
		}
	} else {
		_, err := b.SendDocument(ctx.EffectiveChat.Id, []byte(output), &gotgbot.SendDocumentOpts{
			ReplyToMessageId: msg.MessageId,
		})
		if err != nil {
			logging.Error("got an error when trying to send document: ", err)
			return err
		}
	}

	return nil
}

func SafeEditNoFormat(b *gotgbot.Bot, ctx *ext.Context, topMsg *gotgbot.Message, output string) error {
	if topMsg == nil {
		return SafeReplyNoFormat(b, ctx, output)
	}

	msg := ctx.EffectiveMessage
	if len(output) < 4096 {
		_, err := topMsg.EditText(b, output,
			&gotgbot.EditMessageTextOpts{ParseMode: "Markdownv2"})

		if err != nil {
			logging.Error("got an error when trying to send results: ", err)
			return err
		}
	} else {
		_, err := b.SendDocument(ctx.EffectiveChat.Id, []byte(output), &gotgbot.SendDocumentOpts{
			ReplyToMessageId: msg.MessageId,
		})

		_, _ = topMsg.Delete(b)

		if err != nil {
			logging.Error("got an error when trying to send document: ", err)
			return err
		}
	}

	return nil
}

func GetLink(ctx *ext.Context) string {
	if ctx.EffectiveMessage == nil || ctx.EffectiveChat == nil {
		return ""
	}

	id := ctx.EffectiveMessage.MessageId

	var identifier string
	if ctx.EffectiveChat.Username != "" {
		identifier = ctx.EffectiveChat.Username
	} else {
		identifier = "c/" + strconv.FormatInt(ctx.EffectiveChat.Id, 10)[4:]
	}

	return fmt.Sprintf("https://t.me/%s/%d", identifier, id)
}

func GetLinkFromMessage(msg *gotgbot.Message) string {
	if msg == nil || msg.Chat.Id == 0 {
		return ""
	}

	id := msg.MessageId

	var identifier string
	if msg.Chat.Username != "" {
		identifier = msg.Chat.Username
	} else {
		identifier = "c/" + strconv.FormatInt(msg.Chat.Id, 10)[4:]
	}

	return fmt.Sprintf("https://t.me/%s/%d", identifier, id)
}

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
		if len(v) == 0 {
			for i, j := range c.Request.URL.Query() {
				if strings.EqualFold(i, key) && len(j) > 0 {
					v = j[0]
					break
				}
			}
		}
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

// ReadOneFile attempts to read one file from the specified paths. if file
// doesn't exist or it contains empty data, it will try to read next path.
// it will continue doing so until it reaches a file which exists and contains
// valid data. it will return empty string if none of the files are valid.
func ReadOneFile(paths ...string) string {
	if len(paths) == 0 {
		return ""
	}

	var str string

	for _, current := range paths {
		if len(current) == 0 {
			return ""
		}

		str = ReadFile(current)
		if len(str) != 0 {
			return str
		}
	}

	return ""
}

func GetNameFromUser(u *gotgbot.User, replacement string) string {
	if u == nil {
		return replacement
	}

	name := strings.TrimSpace(u.FirstName)
	if len(name) > 0 {
		return name
	}

	name = strings.TrimSpace(u.LastName)
	if len(name) > 0 {
		return name
	}

	name = strings.TrimSpace(u.Username)
	if len(name) > 0 {
		return name
	}

	return replacement
}
