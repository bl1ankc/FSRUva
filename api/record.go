package api

import (
	"github.com/gin-gonic/gin"
	"main/Model"
)

// GetRecordsByUser 查询人历史记录
func GetRecordsByUser(c *gin.Context) {

	Name := c.Query("name")
	records := Model.GetRecordsByName(Name)
	c.JSON(200, &records)
	return
}

// GetRecordsByUva 查询设备历史记录
func GetRecordsByUva(c *gin.Context) {
	id := c.Query("uid")
	records := Model.GetRecordsByUid(id)
	c.JSON(200, &records)
	return
}
