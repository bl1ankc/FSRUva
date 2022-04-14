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

// GetPassedUav 审核通过借用设备
func GetPassedUav(c *gin.Context) {
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
		Model.UpdateBorrowTime(uav.Uid)
		Model.UpdateRecordState(uav.Uid, uav.Borrower, uav.Get_time, "borrowing")
	}
}

// BackPassedUav 审核通过归还设备
func BackPassedUav(c *gin.Context) {
	//模型定义
	var uavs []Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态与归还时间
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "using")
		Model.UpdateBackTime(uav.Uid)
		Model.UpdateRecordState(uav.Uid, uav.Borrower, uav.Get_time, "free")
	}
}

// GetFailUav 审核不通过借用设备
func GetFailUav(c *gin.Context) {
	//模型定义
	var uavs []Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态与借用时间
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "free")
		Model.UpdateRecordState(uav.Uid, uav.Borrower, uav.Get_time, "free")
	}

	//返回数据
}

// BackFailUav 审核不通过归还设备
func BackFailUav(c *gin.Context) {
	//模型定义
	var uavs []Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态与归还时间
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "using")
		Model.UpdateRecordState(uav.Uid, uav.Borrower, uav.Get_time, "using")
	}

	//返回数据表示不成功
}
