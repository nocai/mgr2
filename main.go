package main

import (
	"mgr2/routers"
	"github.com/gin-gonic/gin"
	"log"
)


func main() {
	r := gin.Default()
	r.Use(DebugInput())
	routers.SetupRouter(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}


func DebugInput() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("请求参数：", context.Params)
	}
}