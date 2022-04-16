package api

import (
	"github.com/gin-gonic/gin"
	"main/Model"
)

// UploadUser 上传用户信息ljy
func UploadUser(c *gin.Context) {
	//模型定义
	var user Model.User

	//结构体绑定
	//绑定结构体
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//数据插入
	response := Model.InsertUser(user.Name, user.Phone, user.StudentID)
	c.JSON(200, response)
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	var user Model.User

	//结构体绑定
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//数据插入
	response := ""
	if user.Phone != "" {
		Model.UpdatePhone(user.Name, user.Phone)
		response += "Phone Update! "
	}
	if user.StudentID != "" {
		Model.UpdateStudentId(user.Name, user.StudentID)
		response += "StudentID Update! "
	}

	c.JSON(200, &response)
}

// GetUser 获取单个用户信息
func GetUser(c *gin.Context) {
	//数据绑定
	Name := c.Query("name")

	//数据获取
	response := Model.GetUserByName(Name)
	c.JSON(200, &response)
}
