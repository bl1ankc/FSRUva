package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Const"
	"main/Model"
	"main/Service"
	"time"
)

// UploadNewUav 上传新设备
func UploadNewUav(c *gin.Context) {
	//模型定义
	var uav Model.Uav

	//结构体绑定
	if err := c.BindJSON(&uav); err != nil {
		fmt.Println("上传新设备数据绑定失败：", err.Error())
		c.JSON(400, gin.H{"msg": "参数格式错误"})
		return
	}

	//数据插入
	flag, response := Service.InsertUva(uav.Name, uav.Type, uav.Uid)
	//Model.CreateQRCode(uav.Uid)
	if flag {
		c.JSON(200, gin.H{"code": 200, "desc": "上传成功"})
	} else {
		c.JSON(200, gin.H{"code": 200, "desc": response})
	}
}

// UpdateDevice 更新设备 @2022/09/06
func UpdateDevice(c *gin.Context) {
	//模型定义
	var uav Model.Uav
	var code int

	//结构体绑定
	if err := c.BindJSON(&uav); err != nil {
		code = Const.FailToBindJson
		fmt.Println("上传新设备数据绑定失败：", err.Error())
		c.JSON(code, R(code, nil, "传输数据错误"))
		return
	}

	err := Service.UpdateDevice(uav)
	if err != nil {
		code = Const.FuncFail
		c.JSON(code, R(code, nil, "更新设备失败"))
		return
	}

	code = Const.Ok
	c.JSON(code, R(code, nil, "更新设备成功"))
	return
}

// DeleteDevice 删除设备 @2022/09/06
func DeleteDevice(c *gin.Context) {
	var code int

	ID := c.Query("uid")
	if ID == "" {
		code = Const.FuncFail
		c.JSON(code, R(code, nil, "传入uid错误"))
		return
	}

	_, Device := Service.GetUavByUid(ID)
	if err := Service.RemoveDevice(Device); err != nil {
		code = Const.FuncFail
		c.JSON(code, R(code, nil, "删除设备失败，检查函数与传入ID是否存在"))
		return
	}

	code = Const.Ok
	c.JSON(code, R(code, nil, "删除成功"))
	return
}

// GetReview 获取审核中设备
func GetReview(c *gin.Context) {
	var uavs []Model.Uav

	GetUav := Service.GetUavByStates("Get under review", "")
	Uav := Service.GetUavByStates("Back under review", "")

	for _, uav := range GetUav {
		uavs = append(uavs, uav)
	}
	for _, uav := range Uav {
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
		fmt.Println("审核通过借用设备数据绑定失败：", err.Error())
		c.JSON(400, gin.H{"msg": "参数格式错误"})
		return
	}

	if uav.Type == "Drone" {
		user := c.MustGet("user").(Model.User)
		if user.AdminType != 1 {
			c.JSON(301, R(301, nil, "非老师操作，无法通过无人机审核"))
			return
		}
	}

	//更新状态与借用时间
	BorrowTime := time.Now()
	err := Service.UpdateState(uav.Uid, "scheduled")
	err = Service.UpdateBorrowTime(uav.Uid, BorrowTime)
	err = Service.GetReviewRecord(uav.Uid, uav.Checker, "passed", uav.Comment, BorrowTime)
	err = Service.UpdateRecordState(uav.Uid, "scheduled")

	if err != nil {
		c.JSON(Const.FuncFail, R(Const.FuncFail, nil, "获取设备失败"))
		return
	}
	c.JSON(200, gin.H{"desc": "审核成功"})
	return
	//Model.UpdateUserCountByUid(uav.Uid, 1)
}

// BackPassedUav 审核通过归还设备
func BackPassedUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		fmt.Println("审核通过归还设备数据绑定失败：", err.Error())
		c.JSON(400, gin.H{"msg": "参数格式错误"})
		return
	}

	//更新状态与归还时间
	Service.UpdateBackRecord(uav.Uid)
	Service.UpdateState(uav.Uid, "free")
	Service.UpdateBackTime(uav.Uid)
	Service.UpdateRecordState(uav.Uid, "returned")
	Service.BackReviewRecord(uav.Uid, uav.Checker, "passed", uav.Comment)
	c.JSON(200, gin.H{"desc": "审核成功"})
}

// GetFailUav 审核不通过借用设备
func GetFailUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		fmt.Println("审核不通过借用设备数据绑定失败：", err.Error())
		c.JSON(400, gin.H{"msg": "参数格式错误"})
		return
	}

	//更新状态与借用时间
	Service.UpdateState(uav.Uid, "free")
	Service.UpdateRecordState(uav.Uid, "refuse")
	Service.GetReviewRecord(uav.Uid, uav.Checker, "fail", uav.Comment, time.Now())
	c.JSON(200, gin.H{"desc": "审核成功"})
}

// BackFailUav 审核不通过归还设备
func BackFailUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		fmt.Println("审核不通过归还设备数据绑定失败：", err.Error())
		c.JSON(400, gin.H{"msg": "参数格式错误"})
		return
	}

	//更新状态与归还时间
	Service.UpdateState(uav.Uid, "using")
	Service.UpdateRecordState(uav.Uid, "using")
	Service.BackReviewRecord(uav.Uid, uav.Checker, "fail", uav.Comment)
	c.JSON(200, gin.H{"desc": "审核成功"})
}

// GetAllUsers 获取所有用户
func GetAllUsers(c *gin.Context) {
	//查找数据
	response := Service.GetAllUsers()
	//返回数据
	c.JSON(200, response)
}

// GetAllRecords 获取所有历史记录
func GetAllRecords(c *gin.Context) {
	//查找数据
	response := Service.GetAllRecords()
	//返回数据
	c.JSON(200, response)
}

// ForceUpdateDevices 强制修改设备信息
func ForceUpdateDevices(c *gin.Context) {
	var uav Model.Uav
	//结构体绑定
	if err := c.BindJSON(&uav); err != nil {
		fmt.Println("强制修改设备信息数据绑定失败：", err.Error())
		c.JSON(400, gin.H{"msg": "参数格式错误"})
		return
	}
	Service.UpdateDevices(uav)
	_, device := Service.GetUavByUid(uav.Uid)

	c.JSON(200, &device)
}

//UpdateUavRemark 修改设备备注信息
func UpdateUavRemark(c *gin.Context) {
	var remark Model.Uav
	//结构体绑定
	if err := c.BindJSON(&remark); err != nil {
		fmt.Println("修改设备备注信息数据绑定失败：", err.Error())
		c.JSON(400, gin.H{"msg": "参数格式错误"})
	}

	Service.UpdateUavRemark(remark.Uid, remark.Remark)
	c.JSON(200, gin.H{"desc": "修改成功"})
}

// GetImgUrl 获取图片临时地址
//func GetImgUrl(c *gin.Context) {
//	imgName := c.Query("imgName")
//
//	url, flag := utils.GetPicUrl(imgName + ".png")
//
//	//绑定结构体
//	type T struct {
//		Url string `json:"url"`
//	}
//	resp := &T{Url: url}
//
//	if flag {
//		c.JSON(200, gin.H{"code": 200, "desc": "获取成功", "data": resp})
//	} else {
//		c.JSON(200, gin.H{"code": 200, "desc": "获取失败"})
//	}
//}
