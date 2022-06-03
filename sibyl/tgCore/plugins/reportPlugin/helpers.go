/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package reportPlugin

import (
	"time"

	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
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

func getMultiReportButtons(uniqueId string) *gotgbot.InlineKeyboardMarkup {
	kb := &gotgbot.InlineKeyboardMarkup{}

	kb.InlineKeyboard = make([][]gotgbot.InlineKeyboardButton, 2)

	kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
		Text:         "✅ Approve",
		CallbackData: MultiReportPrefix + sepChar + ApproveData + sepChar + uniqueId,
	})
	kb.InlineKeyboard[0] = append(kb.InlineKeyboard[0], gotgbot.InlineKeyboardButton{
		Text:         "❌ Reject",
		CallbackData: MultiReportPrefix + sepChar + RejectData + sepChar + uniqueId,
	})

	kb.InlineKeyboard[1] = append(kb.InlineKeyboard[1], gotgbot.InlineKeyboardButton{
		Text:         "Close",
		CallbackData: MultiReportPrefix + sepChar + CloseData + sepChar + uniqueId,
	})

	return kb
}

// pushScanToDatabase converts a scan to a ban and pushes it to the database
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
			database.UpdateBanparameter(u, false)
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

// pushScanToDatabaseWithValidation converts a scan to a ban and pushes it
// to the database if and only if the user's crime coefficient will increase
// by doing so. this function is good only in multi-scan and multi-ban methods.
func pushScanToDatabaseWithValidation(scan *sv.Report) {
	u, err := database.GetUserFromId(scan.TargetUser)
	var count int
	if u != nil && err == nil {
		if u.Banned {
			tmpUser := u.Clone()
			if scan.TargetType != u.TargetType {
				// check both conditions; if they don't match, update the field.
				tmpUser.TargetType = scan.TargetType
			}
			tmpUser.BannedBy = scan.ReporterId
			tmpUser.Message = scan.ReportMessage
			tmpUser.Date = time.Now()
			tmpUser.BanSourceUrl = scan.ScanSourceLink
			tmpUser.SourceGroup = "" /* TODO */
			tmpUser.SetAsBanReason(scan.ReportReason)
			tmpUser.IncreaseCrimeCoefficientAuto()
			if tmpUser.CrimeCoefficient < u.CrimeCoefficient {
				return
			}

			*u = *tmpUser
			database.UpdateBanparameter(u, false)
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

func pushMultipleScanToDatabase(data *sv.MultiScanRawData) {
	if len(data.Origins) == 0 {
		return
	}

	for _, current := range data.Origins {
		pushScanToDatabaseWithValidation(current)
	}
}

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	sv.SendReportHandler = sendReportHandler
	sv.SendMultiReportHandler = sendMultiReportHandler
	sv.SendToADHandler = sendToADHandler

	scanCb := handlers.NewCallback(scanCallBackQuery, scanCallBackResponse)
	multiScanCb := handlers.NewCallback(multiScanCallBackQuery, multiScanCallBackResponse)
	approveCmd := handlers.NewCommand(ApproveCmd, approveHandler)
	aCmd := handlers.NewCommand(ACmd, approveHandler) // 'a' short for approve command
	rejectCmd := handlers.NewCommand(RejectCmd, rejectHandler)
	rCmd := handlers.NewCommand(RCmd, rejectHandler) // 'r' short for reject command
	closeCmd := handlers.NewCommand(CloseCmd, closeHandler)

	scanCb.AllowChannel = true
	multiScanCb.AllowChannel = true

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
	d.AddHandler(multiScanCb)
}
