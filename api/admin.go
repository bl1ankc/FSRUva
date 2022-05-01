package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
	"time"
)

// UploadNewUav 上传新设备
func UploadNewUav(c *gin.Context) {
	//模型定义
	var uav Model.Uav

	//结构体绑定
	if err := c.BindJSON(&uav); err != nil {
		c.JSON(400, gin.H{"msg": "参数格式错误"})
		return
	}

	//数据插入
	Model.InsertUva(uav.Name, uav.Type, uav.Uid)
	Model.CreateQRCode(uav.Uid)

}

// GetReview 获取审核中设备
func GetReview(c *gin.Context) {
	var uavs []Model.Uav

	GetUav := Model.GetUavByStates("Get under review", "")
	BackUav := Model.GetUavByStates("Back under review", "")

	for _, uav := range GetUav {
		uavs = append(uavs, uav)
	}
	for _, uav := range BackUav {
		uavs = append(uavs, uav)
	}

	c.JSON(200, &uavs)
}

// GetPassedUav 审核通过借用设备
func GetPassedUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态与借用时间
	BorrowTime := time.Now()
	Model.UpdateState(uav.Uid, "scheduled")
	Model.UpdateBorrowTime(uav.Uid, BorrowTime)
	Model.GetReviewRecord(uav.Uid, uav.Checker, "passed", uav.Comment, BorrowTime)
	Model.UpdateRecordState(uav.Uid, "scheduled")
	//Model.UpdateUserCountByUid(uav.Uid, 1)
}

// BackPassedUav 审核通过归还设备
func BackPassedUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态与归还时间
	Model.UpdateBackRecord(uav.Uid)
	Model.UpdateState(uav.Uid, "free")
	Model.UpdateBackTime(uav.Uid)
	Model.UpdateRecordState(uav.Uid, "returned")
	Model.BackReviewRecord(uav.Uid, uav.Checker, "passed", uav.Comment)

}

// GetFailUav 审核不通过借用设备
func GetFailUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态与借用时间
	Model.UpdateState(uav.Uid, "free")
	Model.UpdateRecordState(uav.Uid, "refuse")
	Model.GetReviewRecord(uav.Uid, uav.Checker, "fail", uav.Comment, time.Now())

}

// BackFailUav 审核不通过归还设备
func BackFailUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态与归还时间
	Model.UpdateState(uav.Uid, "using")
	Model.UpdateRecordState(uav.Uid, "using")
	Model.BackReviewRecord(uav.Uid, uav.Checker, "fail", uav.Comment)
}

// GetAllUsers 获取所有用户
func GetAllUsers(c *gin.Context) {
	//查找数据
	response := Model.GetAllUsers()
	//返回数据
	c.JSON(200, response)
}

// GetAllRecords 获取所有历史记录
func GetAllRecords(c *gin.Context) {
	//查找数据
	response := Model.GetAllRecords()
	//返回数据
	c.JSON(200, response)
}

// ForceUpdateDevices 强制修改设备信息
func ForceUpdateDevices(c *gin.Context) {
	var uav Model.ChangeUav
	//结构体绑定
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}
	Model.UpdateDevices(uav)
	device := Model.GetUavByUid(uav.Uid)

	c.JSON(200, &device)
}

//UpdateUavRemark 修改设备备注信息
func UpdateUavRemark(c *gin.Context) {
	var remark Model.RemarkUav
	//结构体绑定
	if err := c.BindJSON(&remark); err != nil {
		log.Fatal(err.Error())
	}

	Model.UpdateUavRemark(remark.Uid, remark.Remark)
	c.JSON(200, "OK")
}
