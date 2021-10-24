package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/Dank-del/SibylAPI-Go/config"
	"log"
)

func SibylServer(c *config.ServerConfig) *gin.Engine {
	port := c.Port

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define handlers
	/*r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})*/

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	// r.Run(":" + port)
	return r
}
