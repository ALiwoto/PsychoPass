/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package banHandlers

import sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"

type BanResult struct {
	PreviousBan *sv.User `json:"previous_ban"`
	CurrentBan  *sv.User `json:"current_ban"`
}
