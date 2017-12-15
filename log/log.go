package log

import (
	"log"
	"github.com/gin-gonic/gin"
)

func New() *log.Logger {
	return log.New(gin.DefaultWriter, "[GIN]", log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}
