/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */

package sibylBroadcast

import "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"

type ScanRequestApprovedUpdate struct {
	UniqueId    string                 `json:"unique_id"`
	TargetUser  int64                  `json:"target_user"`
	TargetType  sibylValues.EntityType `json:"target_type"`
	AgentReason string                 `json:"agent_reason"`
}

type ScanRequestRejectedUpdate struct {
	UniqueId    string                 `json:"unique_id"`
	TargetUser  int64                  `json:"target_user"`
	TargetType  sibylValues.EntityType `json:"target_type"`
	AgentReason string                 `json:"agent_reason"`
}
