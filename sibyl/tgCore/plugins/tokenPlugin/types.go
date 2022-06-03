/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package tokenPlugin

import (
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

type AssignValue struct {
	targetChat   *gotgbot.Chat
	perm         string
	permValue    sv.UserPermission
	msg          *gotgbot.Message
	target       *sv.User // before accepting
	targetId     int64    // after accepting
	agentProfile *gotgbot.User
	agent        *sv.Token // before accepting
	agentId      int64     // after accepting
	src          string
}
