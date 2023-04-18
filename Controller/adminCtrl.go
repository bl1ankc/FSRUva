package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Model"
	"main/Service"
	"main/Service/Status"
	"main/Service/Trans"
)

// UploadNewUav 上传新设备
func UploadNewUav(c *gin.Context) {
	//模型定义
	var uav Model.Uav

	//结构体绑定
	if err := c.ShouldBindJSON(&uav); err != nil {
		fmt.Println("上传新设备数据绑定失败：", err.Error())
		c.JSON(400, gin.H{"msg": "参数格式错误"})
		return
	}

	//数据插入
	flag, response := Service.InsertUva(uav)
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
		code = Status.FailToBindJson
		fmt.Println("上传新设备数据绑定失败：", err.Error())
		c.JSON(code, R(code, nil, "传输数据错误"))
		return
	}

	err := Service.UpdateDevice(uav)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "更新设备失败"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "更新设备成功"))
	return
}

// DeleteDevice 删除设备 @2022/09/06
func DeleteDevice(c *gin.Context) {
	var code int

	ID := c.Query("uid")
	if ID == "" {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "传入uid错误"))
		return
	}

	_, Device := Service.GetUavByUid(ID)
	if err := Service.RemoveDevice(Device); err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "删除设备失败，检查函数与传入ID是否存在"))
		return
	}

	code = Status.OK
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

	c.JSON(200, uavs)
	return
}

// GetPassedUav 审核通过借用设备 @2023/4/15更新
func GetPassedUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.ShouldBindJSON(&uav); err != nil {
		fmt.Println("审核通过借用设备数据绑定失败：", err.Error())
		c.JSON(Status.FailToBindJson, R(Status.FailToBindJson, nil, "数据绑定失败"))
		return
	}

	if uav.Type == "Drone" {
		adminType := c.MustGet("adminType").(int)
		if adminType != 1 {
			c.JSON(301, R(301, nil, "非老师操作，无法通过无人机审核"))
			return
		}
	}

	//更新
	if err := Trans.Review(&uav, "GetPass"); err != nil {
		c.JSON(Status.FuncFail, R(Status.FuncFail, nil, "更新函数失败"))
		return
	}

	c.JSON(Status.OK, R(Status.OK, nil, "审核成功"))
	return
}

// BackPassedUav 审核通过归还设备 @2023/4/15更新
func BackPassedUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		fmt.Println("审核通过归还设备数据绑定失败：", err.Error())
		c.JSON(Status.FailToBindJson, R(Status.FailToBindJson, nil, "绑定数据失败"))
		return
	}

	//更新
	if err := Trans.Review(&uav, "BackPass"); err != nil {
		c.JSON(Status.FuncFail, R(Status.FuncFail, nil, "更新函数失败"))
		return
	}

	c.JSON(Status.OK, R(Status.OK, nil, "审核成功"))
	return
}

// GetFailUav 审核不通过借用设备 @2023/4/15更新
func GetFailUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.ShouldBindJSON(&uav); err != nil {
		fmt.Println("审核通过归还设备数据绑定失败：", err.Error())
		c.JSON(Status.FailToBindJson, R(Status.FailToBindJson, nil, "绑定数据失败"))
		return
	}

	if uav.Type == "Drone" {
		adminType := c.MustGet("adminType").(int)
		if adminType != 1 {
			c.JSON(301, R(301, nil, "非老师操作，无法通过无人机审核"))
			return
		}
	}

	//更新
	if err := Trans.Review(&uav, "GetRefuse"); err != nil {
		c.JSON(Status.FuncFail, R(Status.FuncFail, nil, "更新函数失败"))
		return
	}

	c.JSON(Status.OK, R(Status.OK, nil, "审核成功"))
	return
}

// BackFailUav 审核不通过归还设备 @2023/4/15更新
func BackFailUav(c *gin.Context) {
	//模型定义
	var uav Model.CheckUav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		fmt.Println("审核通过归还设备数据绑定失败：", err.Error())
		c.JSON(Status.FailToBindJson, R(Status.FailToBindJson, nil, "绑定数据失败"))
		return
	}

	//更新
	if err := Trans.Review(&uav, "BackRefuse"); err != nil {
		c.JSON(Status.FuncFail, R(Status.FuncFail, nil, "更新函数失败"))
		return
	}

	c.JSON(Status.OK, R(Status.OK, nil, "审核成功"))
	return
}

// ForcedGet 强制取走 @2023/4/16更新
func ForcedGet(c *gin.Context) {
	var code int

	id := c.Query("uid")

	//实例获取
	exist, uav := Service.GetUavByUid(id)
	if !exist {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "设备查询失败,检查是否有该设备记录"))
		return
	}

	//强制归还事务
	if err := Trans.ForcedGet(&uav); err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "数据库操作错误"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "强制修改成功"))
	return
}

// ForcedBack 强制归还 @2023/4/16 更新
func ForcedBack(c *gin.Context) {
	var code int

	id := c.Query("uid")

	//实例获取
	exist, uav := Service.GetUavByUid(id)
	if exist == false {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "无人机不存在"))
		return
	}

	//设备状态检测
	if uav.State != "using" && uav.State != "scheduled" {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "设备不在使用中"))
		return
	}

	//强制归还
	if err := Trans.ForcedBack(&uav); err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "数据库操作错误"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "强制归还成功"))
	return
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
