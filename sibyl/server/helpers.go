package server

import (
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylConfig"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/logging"
	"github.com/gin-gonic/gin"
)

func RunSibylSystem() {
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
