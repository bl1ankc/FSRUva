package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// UploadUser 上传用户信息ljy
func UploadUser(c *gin.Context) {
	//模型定义
	var user Model.User

	//结构体绑定
	//绑定结构体
	if err := c.BindJSON(&user); err != nil {
		log.Fatal(err.Error())
		return
	}
	//数据插入

	Model.InsertUser(user.Name, user.Phone, user.StudentID)
	c.JSON(200, "OK")
	return
}
