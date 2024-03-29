/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package banHandlers

import "sync"

var (
	multiBanMutex   *sync.Mutex
	multiUnBanMutex *sync.Mutex
)
