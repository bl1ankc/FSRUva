package api

import (
	"github.com/gin-gonic/gin"
	"main/Model"
)

// GetRecordsByUser 查询用户历史记录
func GetRecordsByUser(c *gin.Context) {

	//模型绑定
	stuid := c.Query("stuid")

	//查数据
	records := Model.GetRecordsByID(stuid)
	c.JSON(200, &records)
}

// GetRecordsByUva 查询设备历史记录
func GetRecordsByUva(c *gin.Context) {

	//模型绑定
	id := c.Query("uid")

	//查数据
	records := Model.GetRecordsByUid(id)
	c.JSON(200, &records)
}
