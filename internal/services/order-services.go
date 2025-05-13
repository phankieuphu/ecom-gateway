package services

import (
	"fmt"
	config "gateway/configs"

	"github.com/gin-gonic/gin"
)

func InitOrder(c *gin.Context) {
	dbInstance := config.GetDb()
	user := dbInstance.Table("user").Select("")
	fmt.Println(user)
	c.JSON(200, gin.H{
		"Message": "Init order success",
	})
}
