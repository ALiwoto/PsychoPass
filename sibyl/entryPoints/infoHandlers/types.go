/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package infoHandlers

import sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"

type GeneralInfoResult struct {
	// UserId is the id of the user.
	UserId int64 `json:"user_id"`

	// Division is the division number of the user.
	Division int `json:"division"`

	// AssignedBy is the user id of the user who assigned the user.
	AssignedBy int64 `json:"assigned_by"`

	// AssignedReason is the reason why the user is assigned to this permission.
	AssignedReason string `json:"assigned_reason"`

	// AssignedAt is the assignment date of the user.
	AssignedAt string `json:"assigned_at"`

	// Permission is the permission of the user.
	Permission sv.UserPermission `json:"permission"`

	// ApprovedScansCount is the count of accepted reports.
	ApprovedScansCount int `json:"approved_scans_count"`

	// RejectedScansCount is the count of denied reports.
	RejectedScansCount int `json:"rejected_scans_count"`

	// ForcedScansCount is the count of forced scan that this user has
	// sent (the ban count).
	ForcedScansCount int `json:"forced_scans_count"`

	// RevertCount is the count of revert request that this user has
	// sent, inspector's revert request is direct, but enforcer's
	// revert request has to be approved by an inspector at NONA tower.
	RevertCount int `json:"revert_count"`
}

type GetBansResult struct {
	Users []sv.User `json:"users"`
}
