package main

import (
	"mgr2/routers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"io"
)


func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	l := log.New(gin.DefaultWriter, "[GIN]", log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	l.Println("aaaaaaaaaaaaaa")

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
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