package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// BorrowUav 借用设备
func BorrowUav(c *gin.Context) {
	//模型定义
	var uavs []Model.Uav

	//结构体绑定
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态为审核中
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "Get under review")
		Model.UpdateBorrower(uav.Uid, uav.Borrower, uav.Phone)
	}
}

// BackUav 归还设备
func BackUav(c *gin.Context) {
	//模型
	var uavs []Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态为归还审核
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "Back under review")
	}
}
