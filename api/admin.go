package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// UploadNewUav 上传新设备
func UploadNewUav(c *gin.Context) {
	//模型定义
	var uavs []Model.Uav

	//结构体绑定
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//数据插入
	for _, uav := range uavs {
		Model.InsertUva(uav.Name, uav.Type)
	}
}

// GetReviewUav 审核借用设备
func GetReviewUav(c *gin.Context) {
	//模型定义
	var uavs []Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态与借用时间
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "borrowing")
		Model.UpdateBorrowTime(uav.Uid, uav.Plan_time)
	}
}

// BackReviewUav 审核归还设备
func BackReviewUav(c *gin.Context) {
	//模型定义
	var uavs []Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态与归还时间
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "free")
		Model.UpdateBackTime(uav.Uid)
	}
}
