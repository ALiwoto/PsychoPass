package sibylValues

import (
	"strconv"
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
	value = strings.ToLower(strings.TrimSpace(value))
	// first of all check and see if value is an integer or not
	valueInt, err := strconv.Atoi(value)
	if err == nil {
		perm := UserPermission(valueInt)
		if perm.IsValid() {
			return perm, nil
		}
		// we already know that the value is a valid integer, so there is no
		// chance that the value is a valid permission in string format.
		return NormalUser, ErrInvalidPerm
	}

	switch value {
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

func IsInvalidID(id int64) bool {
	return id == 777000 || id == 1087968824
}
