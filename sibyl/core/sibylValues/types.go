/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package sibylValues

import (
	"context"
	"time"

	wc "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues/whatColor"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

type UserPermission int
type ScanStatus int
type EntityType int
type PollingUniqueId uint64
type BanFlag string
type SibylUpdateType string
type ReportHandler func(r *Report)
type MultiReportHandler func(data *MultiScanRawData)

type StatValue struct {
	BannedCount          int64     `json:"banned_count"`
	TrollingBanCount     int64     `json:"trolling_ban_count"`
	SpamBanCount         int64     `json:"spam_ban_count"`
	EvadeBanCount        int64     `json:"evade_ban_count"`
	CustomBanCount       int64     `json:"custom_ban_count"`
	PsychoHazardBanCount int64     `json:"psycho_hazard_ban_count"`
	MalImpBanCount       int64     `json:"mal_imp_ban_count"`
	NSFWBanCount         int64     `json:"nsfw_ban_count"`
	SpamBotBanCount      int64     `json:"spam_bot_ban_count"`
	RaidBanCount         int64     `json:"raid_ban_count"`
	MassAddBanCount      int64     `json:"mass_add_ban_count"`
	CloudyCount          int64     `json:"cloudy_count"`
	TokenCount           int64     `json:"token_count"`
	InspectorsCount      int64     `json:"inspectors_count"`
	EnforcesCount        int64     `json:"enforces_count"`
	cacheTime            time.Time `json:"-"`
}

type User struct {
	UserID           int64       `json:"user_id" gorm:"primaryKey"`
	Banned           bool        `json:"banned"`
	Reason           string      `json:"reason"`
	Message          string      `json:"message"`
	BanSourceUrl     string      `json:"ban_source_url"`
	Date             time.Time   `json:"-"`
	BannedBy         int64       `json:"banned_by"`
	TargetType       EntityType  `json:"target_type"`
	CrimeCoefficient int         `json:"crime_coefficient"`
	BanDate          string      `json:"date" gorm:"-" sql:"-"`
	SourceGroup      string      `json:"source_group"`
	HueColor         wc.HueColor `json:"hue_color" gorm:"-" sql:"-"`
	BanFlags         []BanFlag   `json:"ban_flags" gorm:"-" sql:"-"`
	BanCount         int         `json:"-"` // internal usage only; not meant to be seen by users.
	FlagTrolling     bool        `json:"-"`
	FlagSpam         bool        `json:"-"`
	FlagEvade        bool        `json:"-"`
	FlagCustom       bool        `json:"-"`
	FlagPsychoHazard bool        `json:"-"`
	FlagMalImp       bool        `json:"-"`
	FlagNsfw         bool        `json:"-"`
	FlagRaid         bool        `json:"-"`
	FlagSpamBot      bool        `json:"-"`
	FlagMassAdd      bool        `json:"-"`
}

type Token struct {
	// UserId is the user id
	UserId int64 `json:"user_id" gorm:"primaryKey"`

	// Hash is the user's token hash
	Hash string `json:"hash"`

	// the user's permissions
	Permission UserPermission `json:"permission"`

	AssignedBy     int64  `json:"assigned_by"`
	DivisionNum    int    `json:"division_num"`
	AssignedReason string `json:"assigned_reason"`

	// LastUsage is the user's last usage time
	LastUsage time.Time `json:"-"`

	// CreatedAt is the creation time of the token.
	CreatedAt time.Time `json:"created_at"`

	// LastRevokeDate is the last time this user has revoked the token.
	LastRevokeDate time.Time `json:"-"`

	// RevokeCount is the amount of token being revoked by this user since
	// `LastRevokeDate` field.
	RevokeCount int `json:"-"`

	// AcceptedReports is the count of accepted reports.
	AcceptedReports int `json:"accepted_reports"`

	// DeniedReports is the count of denied reports.
	DeniedReports int `json:"denied_reports"`

	// ForcedCount is the count of forced scan that this user has
	// sent (the ban count).
	ForcedCount int `json:"forced_count"`

	// RevertCount is the count of revert request that this user has
	// sent, inspector's revert request is direct, but enforcer's
	// revert request has to be approved by an inspector at NONA tower.
	RevertCount int `json:"revert_count"`
}

// PollingIdentifier represents a unique polling identifier.
type PollingIdentifier struct {
	PollingUniqueId   PollingUniqueId `json:"polling_unique_id"`
	PollingAccessHash string          `json:"polling_access_hash"`
}

type PollingUserUpdate struct {
	UpdateType SibylUpdateType `json:"update_type"`
	UpdateData any             `json:"update_data"`
}

type RegisteredPollingValue struct {
	OwnerId       int64
	Timeout       time.Duration
	UniqueId      PollingUniqueId
	AccessHash    string
	theChannel    chan *PollingUserUpdate
	ctx           context.Context
	cancelFunc    context.CancelFunc
	isPersistance bool
}

type AssaultDominatorData struct {
	Type         string    `json:"type"`
	TargetUser   int64     `json:"user"`
	ShortReasons []BanFlag `json:"short_reasons"`
	LongReason   string    `json:"long_reason"`
	ScannedBy    int64     `json:"scanned_by"`
	SrcUrl       string    `json:"src_url"`
}

type Report struct {
	// UniqueId is the unique id of the report.
	UniqueId string `json:"unique_id" gorm:"primaryKey"`

	// ReporterId is the reporter id.
	ReporterId int64 `json:"reporter_id"`

	// TargetUser is the id of the target.
	TargetUser int64 `json:"target_user"`

	// TargetType is the type of the target.
	TargetType EntityType `json:"target_type"`

	// ReportDate is the date of report.
	ReportDate string `json:"report_date"`

	// ReportReason is the reason of reporting.
	ReportReason string `json:"report_reason"`

	// ReportMessage is the message that the reported message.
	ReportMessage string `json:"report_message"`

	// ScanSourceLink is the link to the source of the scan.
	ScanSourceLink string `json:"scan_source_link"`

	// AgentId is the user id of the person who approved, rejected or closed the
	// scan.
	AgentId int64 `json:"agent_id"`

	// AgentReason is the reason that agent used to reject or close the
	// scan.
	AgentReason string `json:"agent_reason"`

	// AgentDate is the date that agent used to approve, reject or close the scan.
	AgentDate time.Time `json:"agent_date"`

	// ReporterPermission is the permission of the reporter.
	ReporterPermission UserPermission `json:"reporter_permission"`

	// ScanStatus is the status of this scan. it can either be approved, rejected or
	// closed.
	// please notice that only pending scans' details can be changed; if a
	// scan is approved, rejected or closed, its details cannot be changed anymore.
	ScanStatus ScanStatus `json:"scan_status"`

	// AssociationBanId is empty if the scan is not an association scan.
	AssociationBanId string `json:"association_ban_id"`

	// AgentUser is the agent user who has clicked the approve, reject or close
	// button (or sent the command).
	AgentUser *gotgbot.User `json:"-" gorm:"-" sql:"-"`

	// PollingId is the id of the polling that the result should be sent
	// to. it can be nil.
	PollingId *PollingIdentifier `json:"-" gorm:"-" sql:"-"`
}

// CrimeCoefficientRange is the range of crime coefficients.
type CrimeCoefficientRange struct {
	start int
	end   int
}

type MultiBanUserInfo struct {
	UserId      int64      `json:"user_id"`
	Reason      string     `json:"reason"`
	SourceGroup string     `json:"source_group"`
	Message     string     `json:"message"`
	Source      string     `json:"source"`
	TargetType  EntityType `json:"target_type"`
}

type MultiBanRawData struct {
	Users    []MultiBanUserInfo `json:"users"`
	IsSilent bool               `json:"is_silent"`
}

type MultiUnBanRawData struct {
	Users []int64 `json:"users"`
}

type MultiScanUserInfo struct {
	UserId     int64      `json:"user_id"`
	Reason     string     `json:"reason"`
	Message    string     `json:"message"`
	TargetType EntityType `json:"target_type"`
}

type MultiScanRawData struct {
	AssociationBanId string              `json:"-"`
	Users            []MultiScanUserInfo `json:"users"`
	Source           string              `json:"source"`
	GroupLink        string              `json:"group_link"`
	// ReporterPermission is the permission of the person who has sent this scan.
	// ignored by json.
	ReporterPermission UserPermission `json:"-"`
	AgentUser          *gotgbot.User  `json:"-"`
	AgentDate          time.Time      `json:"-"`
	ReporterId         int64          `json:"-"`
	Status             ScanStatus     `json:"-"`
	Origins            []*Report      `json:"-"`
}

type Triggers struct {
	Spam             []string `json:"spam"`
	Trolling         []string `json:"trolling"`
	Evade            []string `json:"evade"`
	MalImpersonation []string `json:"mal_impersonation"`
	PsychoHazard     []string `json:"psycho_hazard"`
	Nsfw             []string `json:"nsfw"`
	Raid             []string `json:"raid"`
	SpamBot          []string `json:"spam_bot"`
	MassAdd          []string `json:"mass_add"`
}
