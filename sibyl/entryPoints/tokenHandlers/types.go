/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package tokenHandlers

import sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"

type ChangePermResult struct {
	PreviousPerm sv.UserPermission `json:"previous_perm"`
	CurrentPerm  sv.UserPermission `json:"current_perm"`
}

type GetRegisteredResult struct {
	RegisteredUsers []int64 `json:"registered_users"`
}
