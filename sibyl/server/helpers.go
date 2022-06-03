/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package server

import (
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylConfig"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues/whatColor"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
	"github.com/gin-gonic/gin"
)

func RunSibylSystem() {
	whatColor.LoadColors()
	port := sibylConfig.GetPort()

	if !sibylConfig.IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Starts a new Gin instance with no middle-ware
	ServerEngine = gin.New()
	LoadHandlers()
	// Listen and serve on defined port
	logging.Info("Listening on port ", port)

	err := ServerEngine.Run(":" + port)
	if err != nil {
		logging.Error(err)
	}
}
