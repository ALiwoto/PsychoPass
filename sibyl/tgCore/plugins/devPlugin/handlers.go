package devPlugin

import (
	"os"
	"strings"
	"time"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/shellUtils"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func gitpullHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	if !sibylConfig.IsDevUser(user.Id) {
		return ext.ContinueGroups
	}

	isWindows := os.PathListSeparator == '\\'

	msg := ctx.EffectiveMessage
	isForced := strings.Contains(msg.Text, "restart")

	topMsg, err := msg.Reply(b, "Trying to pull changes...", nil)
	if err != nil || topMsg == nil {
		return ext.EndGroups
	}

	var output, errout string
	branch := strings.Join(ws.SplitN(msg.Text, 2, " ", "\n", "\r", "\t")[1:], "")
	branch = strings.TrimSpace(branch)
	if branch == "" {
		// find the correct branch
		output, errout, err = shellUtils.Shellout(`git branch`)
		if err != nil {
			return utils.SafeReply(b, ctx, err.Error())
		}

		output += errout
		myStrs := strings.Split(output, "\n")
		for _, str := range myStrs {
			if strings.HasPrefix(str, "*") {
				branch = strings.TrimSpace(strings.TrimLeft(str, "*"))
				break
			}
		}
	}
	whole := "git pull origin " + branch

	output, errout, err = shellUtils.Shellout(whole)

	mdResult := mdparser.GetEmpty()
	if err != nil {
		mdResult.Bold("\n\nError:\n").Mono(err.Error())
		if len(errout) != 0 {
			mdResult.Mono("\n" + errout)
		}
	} else {
		if len(output) == 0 && len(errout) == 0 {
			mdResult.Mono("None")
		} else {
			mdResult.Mono(
				strings.ReplaceAll(output, b.Token, "$TOKEN"))
			mdResult.Mono(
				strings.ReplaceAll("\n"+errout, b.Token, "$TOKEN"))
		}
	}
	_ = utils.SafeEditNoFormat(b, ctx, topMsg, mdResult.ToString())

	if strings.Contains(output+errout, "Already up to date.") && !isForced {
		return ext.EndGroups
	}

	_, _ = b.SendMessage(msg.Chat.Id, "Restarting...", nil)

	go shellUtils.RestartBot(isWindows)

	if !isWindows {
		time.Sleep(3500 * time.Millisecond)

		os.Exit(0x0)
	}

	return ext.EndGroups
}
