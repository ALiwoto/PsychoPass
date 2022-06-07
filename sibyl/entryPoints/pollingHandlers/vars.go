/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package pollingHandlers

import (
	"github.com/AnimeKaizoku/ssg/ssg"
	sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

var (
	pollingNumGenerator = ssg.NewNumIdGenerator[sv.PollingUniqueId]()
)
