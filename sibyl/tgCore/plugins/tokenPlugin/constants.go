/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package tokenPlugin

const (
	RevokeCmd          = "revoke"
	AssignCmd          = "assign"
	GetTokenCbValue    = "get_token"
	RevokeTokenCbValue = "revoke_token"
	SpecialChar        = "\u200D"
	AssignCbData       = "ag"
	AssignCbPrefix     = AssignCbData + CbSep
	RejectCbData       = "rej"
	CloseCbData        = "close"
	CbSep              = "_"
)
