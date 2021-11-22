package banHandlers

// Update ErrTooManyUsers in sibyl/entryPoints/constants.go if this is edited
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
