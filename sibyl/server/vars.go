/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package server

import "github.com/gin-gonic/gin"

var ServerEngine *gin.Engine

var (
	getHandlers  = make(map[string]gin.HandlerFunc)
	postHandlers = make(map[string]gin.HandlerFunc)
)
