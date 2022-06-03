/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */

package sibylBroadcast

import "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"

func SendScanRequestApproved(r sibylValues.Report) {
	sibylValues.BroadcastUpdate(&sibylValues.PollingUserUpdate{
		UpdateType: UpdateTypeScanRequestApproved,
		UpdateData: &ScanRequestApprovedUpdate{
			UniqueId:    r.UniqueId,
			TargetUser:  r.TargetUser,
			TargetType:  r.TargetType,
			AgentReason: r.AgentReason,
		},
	})
}

func SendScanRequestRejected(r sibylValues.Report) {
	sibylValues.BroadcastUpdate(&sibylValues.PollingUserUpdate{
		UpdateType: UpdateTypeScanRequestRejected,
		UpdateData: &ScanRequestRejectedUpdate{
			UniqueId:    r.UniqueId,
			TargetUser:  r.TargetUser,
			TargetType:  r.TargetType,
			AgentReason: r.AgentReason,
		},
	})
}
