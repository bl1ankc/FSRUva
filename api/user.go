package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// BorrowUav 借用设备
func BorrowUav(c *gin.Context) {
	//模型定义
	var uavs []Model.BorrowUav

	//结构体绑定
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//表单中提交不可使用的无人机
	flag := false
	var erruav []Model.Uav
	var Uids []string

	//更新状态为审核中
	for _, uav := range uavs {
		//再次验证是否能被借用
		if uav.State != "free" {
			flag = true
			Uids = append(Uids, uav.Uid)
			continue
		}
		Model.UpdateState(uav.Uid, "Get under review")
		Model.UpdateBorrower(uav.Uid, uav.Borrower, uav.Phone)
		Model.UpdatePlanTime(uav.Uid, uav.Plan_time)
		Model.RecordBorrow(uav.Uid, uav.Borrower, uav.Get_time, uav.Plan_time, "uav.Usage") //用途
	}
	erruav = Model.GetUavsByUids(Uids)
	//返回错误信息
	if flag {
		c.JSON(200, erruav)
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

// CancelBorrow 取消借用
func CancelBorrow(c *gin.Context) {
	//模型定义
	var uavs []Model.Uav

	//结构体绑定
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态为审核中
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "free")
		Model.UpdateRecordState(uav.Uid, "cancelled")
	}

}

// CancelBack 取消归还
func CancelBack(c *gin.Context) {
	//模型
	var uavs []Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态为归还审核
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "using")
	}
}
