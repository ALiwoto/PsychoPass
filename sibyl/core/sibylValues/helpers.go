package sibylValues

import (
	"strings"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/timeUtils"
)

func NewReport(reason, message string, target, reporter int64,
	reporterPerm UserPermission) *Report {

	return &Report{
		ReportReason:       reason,
		ReportMessage:      message,
		TargetUser:         target,
		ReporterId:         reporter,
		ReportDate:         timeUtils.GenerateCurrentDateTime(),
		ReporterPermission: reporterPerm.GetStringPermission(),
	}
}

func ConvertToPermission(value string) (UserPermission, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "user", "civilian":
		return NormalUser, nil
	case "enforcer":
		return Enforcer, nil
	case "inspector":
		return Inspector, nil
	case "owner":
		return Owner, nil
	default:
		return NormalUser, ErrInvalidPerm
	}
}
