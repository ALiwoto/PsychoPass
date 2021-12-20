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
}
