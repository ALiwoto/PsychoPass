/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package database

import sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"

type BanInfo struct {
	UserID     int64
	Adder      int64
	Reason     string
	SrcGroup   string
	Message    string
	Src        string
	TargetType sv.EntityType
	Count      int
	IsSilent   bool
}
