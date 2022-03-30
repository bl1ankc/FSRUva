package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// GetNotUsedDrones 获取空闲的无人机设备
func GetNotUsedDrones(c *gin.Context) {
	//获取设备信息
	uav := Model.GetUavByStates("free", "uav")

	//JSON格式返回
	c.JSON(200, &uav)
}

// GetNotUsedBattery 获取空闲的电池设备
func GetNotUsedBattery(c *gin.Context) {
	//获取设备信息
	Battery := Model.GetUavByStates("free", "Battery")

	//JSON格式返回
	c.JSON(200, &Battery)
}

// GetNotUsedControl 获取空闲的遥控设备
func GetNotUsedControl(c *gin.Context) {
	//获取设备信息
	Control := Model.GetUavByStates("free", "Control")

	//JSON格式返回
	c.JSON(200, &Control)
}

// GetDrones 获取所有无人机设备
func GetDrones(c *gin.Context) {
	uav := Model.GetUavByStates("", "uav")

	c.JSON(200, &uav)
}

// UploadNewUav 上传新设备
func UploadNewUav(c *gin.Context) {
	//模型定义
	var uav Model.Uav

	//结构体绑定
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}

	//数据插入
	Model.InsertUva(uav.Name, uav.Type, uav.Uid)

}

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
		Model.UpdateState(uav.Uid, "under review")
		//Model.UpdateBorrower(uav.Uid, uav.Borrower, uav.Phone)
		//Model.UpdateBorrowTime(uav.Uid, uav.Plan_time, uav.Plan_time)
	}
}
