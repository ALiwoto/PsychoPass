package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylConfig"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/logging"
)

func RunSibylSystem() {
	port := sibylConfig.GetPort()
	// Starts a new Gin instance with no middle-ware
	ServerEngine = gin.New()

	// Define handlers
	/*r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})*/
	LoadHandlers()
	// Listen and serve on defined port
	logging.Info("Listening on port ", port)

	ServerEngine.Run(":" + port)
}
