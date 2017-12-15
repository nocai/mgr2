package userouter

import (
	"github.com/gin-gonic/gin"
	"mgr2/service/userser"
	"net/http"
	"strconv"
	"fmt"
	"log"
)

func Route(r *gin.Engine) {
	r.GET("/user/:id", GetUser)
}


func GetUser(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	fmt.Println(userId)
	log.Println("aaaaaaaaaaaaaaaaa")
	user := userser.GetUserById(userId)
	c.JSON(http.StatusOK, user)
}
