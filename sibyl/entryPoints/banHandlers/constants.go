/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package banHandlers

// maximum constants
const (
	MaxMultiUsers = 50_000
)

const (
	OriginAddBan    = "AddBan"
	OriginMultiBan  = "MultiBan"
	OriginRemoveBan = "RemoveBan"
)

const (
	MessageUnbanned           = "User was unbanned"
	MessageHistoryCleared     = "User's history has been cleared"
	MessageApplyingMultiBan   = "Applying your multi-ban request in background"
	MessageApplyingMultiUnBan = "Applying your multi-unban request in background"
)
