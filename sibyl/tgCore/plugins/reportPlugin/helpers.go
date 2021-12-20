package reportPlugin

import (
	"time"

	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func getReportButtons(uniqueId string) *gotgbot.InlineKeyboardMarkup {
	kb := &gotgbot.InlineKeyboardMarkup{}

	kb.InlineKeyboard = make([][]gotgbot.InlineKeyboardButton, 2)

	kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
		Text:         "✅ Approve",
		CallbackData: ReportPrefix + sepChar + ApproveData + sepChar + uniqueId,
	})
	kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
		Text:         "❌ Reject",
		CallbackData: ReportPrefix + sepChar + RejectData + sepChar + uniqueId,
	})

	kb.InlineKeyboard[1] = append(kb.InlineKeyboard[1], gotgbot.InlineKeyboardButton{
		Text:         "Close",
		CallbackData: ReportPrefix + sepChar + CloseData + sepChar + uniqueId,
	})

	return kb
}

func sendReportMessage(chat int64, text string, opts *gotgbot.SendMessageOpts) {
	_, err := sv.HelperBot.SendMessage(chat, text, opts)
	if err != nil {
		logging.Debug("Tried to send message to ", chat, err)
	}
}

func pushScanToDatabase(scan *sv.Report) {
	u, err := database.GetUserFromId(scan.TargetUser)
	var count int
	if u != nil && err == nil {
		if u.Banned {
			if scan.TargetType != u.TargetType {
				// check both conditions; if they don't match, update the field.
				u.TargetType = scan.TargetType
			}
			u.BannedBy = scan.ReporterId
			u.Message = scan.ReportMessage
			u.Date = time.Now()
			u.BanSourceUrl = scan.ScanSourceLink
			u.SourceGroup = "" /* TODO */
			u.SetAsBanReason(scan.ReportReason)
			u.IncreaseCrimeCoefficientAuto()
			database.UpdateBanparameter(u)
			return
		}
		count = u.BanCount
	}

	info := &database.BanInfo{
		UserID:     scan.TargetUser,
		Adder:      scan.ReporterId,
		Reason:     scan.ReportReason,
		SrcGroup:   "", /* TODO */
		Src:        scan.ScanSourceLink,
		Message:    scan.ReportMessage,
		TargetType: scan.TargetType,
		Count:      count,
	}

	database.AddBan(info)
}

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	sv.SendReportHandler = sendReportHandler
	sv.SendMultiReportHandler = sendMultiReportHandler

	scanCb := handlers.NewCallback(scanCallBackQuery, scanCallBackResponse)
	approveCmd := handlers.NewCommand(ApproveCmd, approveHandler)
	aCmd := handlers.NewCommand(ACmd, approveHandler)
	rejectCmd := handlers.NewCommand(RejectCmd, rejectHandler)
	rCmd := handlers.NewCommand(RCmd, rejectHandler)
	closeCmd := handlers.NewCommand(CloseCmd, closeHandler)
	scanCb.AllowChannel = true
	approveCmd.Triggers = triggers
	aCmd.Triggers = triggers
	rejectCmd.Triggers = triggers
	rCmd.Triggers = triggers
	closeCmd.Triggers = triggers
	d.AddHandler(approveCmd)
	d.AddHandler(aCmd)
	d.AddHandler(rejectCmd)
	d.AddHandler(rCmd)
	d.AddHandler(closeCmd)
	d.AddHandler(scanCb)
}
