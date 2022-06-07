/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package pollingHandlers

import sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"

type StartPollingResult struct {
	PollingUniqueId   sv.PollingUniqueId `json:"polling_unique_id"`
	PollingAccessHash string             `json:"polling_access_hash"`
}
