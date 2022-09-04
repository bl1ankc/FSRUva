package Controller

import (
	"github.com/gin-gonic/gin"
	"main/Const"
	"main/Service"
)

// GetRecordsByUser 查询用户历史记录
func GetRecordsByUser(c *gin.Context) {

	//模型绑定
	stuid := c.Query("stuid")

	//查数据
	records := Service.GetRecordsByID(stuid)
	c.JSON(200, &records)
}

// GetRecordsByUva 查询设备历史记录
func GetRecordsByUva(c *gin.Context) {

	//模型绑定
	id := c.Query("uid")
	page := c.DefaultQuery("page", "0")

	//查数据
	records := Service.GetRecordsByUid(id, page, Const.PAGEMAX)
	c.JSON(200, gin.H{"code": "200", "desc": "获取成功", "data": &records})
}
