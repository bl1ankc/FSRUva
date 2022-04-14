package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// GetRecordsByUser 查询历史记录
func GetRecordsByUser(c *gin.Context) {

	Name := c.Query("name")
	records := Model.GetRecordsByName(Name)
	c.JSON(200, &records)
	return
}

//查询历史记录（还没写好）
func GetRecordsByUva(c *gin.Context) {
	var user Model.User

	//结构体绑定
	//绑定结构体
	if err := c.BindJSON(&user); err != nil {
		log.Fatal(err.Error())
		return
	}
	return
}
