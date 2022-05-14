/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package infoHandlers

import sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"

type GeneralInfoResult struct {
	UserId         int64             `json:"user_id"`
	Division       int               `json:"division"`
	AssignedBy     int64             `json:"assigned_by"`
	AssignedReason string            `json:"assigned_reason"`
	AssignedAt     string            `json:"assigned_at"`
	Permission     sv.UserPermission `json:"permission"`
}

type GetBansResult struct {
	Users []sv.User `json:"users"`
}
