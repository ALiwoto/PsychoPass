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
