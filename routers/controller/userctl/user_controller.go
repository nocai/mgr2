package userctl

import (
	"github.com/gin-gonic/gin"
	"mgr2/models/service/userser"
	"net/http"
	"strconv"
	"fmt"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	fmt.Println(userId)
	user := userser.GetUserById(userId)
	c.JSON(http.StatusOK, user)
}
