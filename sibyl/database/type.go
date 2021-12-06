package database

type BanInfo struct {
	UserID,
	Adder int64
	Reason,
	SrcGroup,
	Message,
	Src string
	IsBot bool
	Count int
}
