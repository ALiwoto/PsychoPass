package reportPlugin

import (
	"strconv"
	"strings"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylConfig"
	sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/logging"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/database"
)

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	sv.SetReportHandler(SendReportHandler)
	reportCB := handlers.NewCallback(reportCallBackQuery, reportCallBackResponse)
	d.AddHandler(reportCB)
}

func SendReportHandler(r *sv.Report) {
	// prevent from panic xD
	if sv.HelperBot == nil {
		return
	}

	bases := sibylConfig.GetBaseChatIds()
	if len(bases) == 0 {
		// there is no chat to send the report to...
		// ignore the report...
		return
	}

	var text string
	var opts *gotgbot.SendMessageOpts

	// let the report value generate a new unique id for itself.
	r.CreateUniqueId()

	md := mdparser.GetNormal("#Report Event:\n")
	md.AppendBoldThis("・User:").AppendNormalThis(" ")
	md.AppendMentionThis(strconv.FormatInt(r.ReporterId, 10), r.ReporterId)
	md.AppendNormalThis("\n")
	md.AppendBoldThis("・By " + r.ReporterPermission).AppendNormalThis(" ")
	md.AppendMentionThis(strconv.FormatInt(r.ReporterId, 10), r.ReporterId)
	md.AppendNormalThis("\n")
	md.AppendBoldThis("・Reason:").AppendNormalThis(" ")
	md.AppendMonoThis(r.ReportReason)
	md.AppendNormalThis("\n")
	md.AppendBoldThis("・Date:").AppendNormalThis(" ")
	md.AppendItalicThis(r.ReportDate)
	md.AppendNormalThis("\n\n")
	md.AppendBoldThis("・Message:").AppendNormalThis(" ")
	md.AppendNormalThis(r.ReportMessage)

	text = md.ToString()
	opts = &gotgbot.SendMessageOpts{
		ParseMode:   sv.MarkDownV2,
		ReplyMarkup: getReportButtons(r.GetUniqueId()),
	}

	for _, chat := range bases {
		sendReportMessage(chat, text, opts)
	}
}

//  func(cq *gotgbot.CallbackQuery)
func reportCallBackQuery(cq *gotgbot.CallbackQuery) bool {
	return strings.HasPrefix(cq.Data, reportCbPrefix+cbDataSeparator)
}

// type Response func(b *gotgbot.Bot, ctx *ext.Context) error
func reportCallBackResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveUser == nil {
		return ext.EndGroups
	}

	user := ctx.EffectiveUser
	cb := ctx.CallbackQuery
	mystrs := strings.Split(cb.Data, cbDataSeparator)
	if len(mystrs) != 3 {
		return ext.EndGroups
	}

	t, err := database.GetTokenFromId(user.Id)
	if err != nil {
		logging.UnexpectedError(err)
	}

	if t.CanBan() {
		go func() {
			_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text:      "Done!",
				ShowAlert: true,
				CacheTime: 5,
			})
			if err != nil {
				logging.Error(err)
			}
		}()
	} else {
		go func() {
			_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
				Text:      "Permission denied.",
				ShowAlert: true,
				CacheTime: 5,
			})
			if err != nil {
				logging.Error(err)
			}
		}()
		return ext.EndGroups
	}

	reportId := mystrs[2]
	uniqueId, err := strconv.ParseInt(reportId, reportIdBase, 64)
	if err != nil || uniqueId == 0 {
		return ext.EndGroups
	}

	report := sv.GetReportFromUniqueId(uniqueId, user.Id)

	//doActionForMessage(b, ctx, report)
	switch mystrs[1] {
	case acceptCbValue:
		err := editAcceptedMessage(b, ctx)
		if err != nil {
			return err
		}
		if report != nil {
			report.MarkAsAccepted()
		}
	case closeCbValue:
		err := editClosedMessage(b, ctx)
		if err != nil {
			return err
		}
		if report != nil {
			report.MarkAsClosed()
		}
	case deleteCbValue:
		_, err := ctx.EffectiveMessage.Delete(b)
		if err != nil {
			return err
		}
		if report != nil {
			report.MarkAsClosed()
		}
	default:
		return ext.EndGroups
	}

	if report == nil {
		return ext.EndGroups
	}

	database.ApplyReport(report)

	// destroy the report value, so it cannot be used the next time
	// someone clicks on one of the buttons in another chat.
	report.Destroy()

	return ext.EndGroups
}
