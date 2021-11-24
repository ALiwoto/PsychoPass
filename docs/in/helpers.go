package in

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func LoadDocs(e *gin.Engine) {
	p, _ := os.Getwd()
	p = filepath.Join(p, "docs", "in")
	e.Static("/docs", p)
}
