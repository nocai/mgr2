package routers

import (
	"github.com/gin-gonic/gin"
	"mgr2/routers/userouter"
)


func SetupRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	userouter.Route(r)
}

