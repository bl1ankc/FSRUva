package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// BorrowerUav 借用设备
func BorrowerUav(c *gin.Context) {
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
		//Model.UpdateBorrowTime(uav.Uid, uav.Get_time, uav.Plan_time)
	}
}

// BackUav 归还设备
func BackUav(c *gin.Context) {
	//
	var uavs []Model.Uav

	//
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "Back under review")
		//Model.UpdateBackTime(uav.Uid,uav.Back_time)
	}
}
