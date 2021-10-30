package sibylValues

import (
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
