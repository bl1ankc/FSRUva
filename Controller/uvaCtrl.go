package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/Const"
	"main/Model"
	"main/Service"
	"main/utils"
)

// GetNotUsedDrones 获取空闲的设备
func GetNotUsedDrones(c *gin.Context) {

	page := c.DefaultQuery("page", "0")
	typename := c.Query("type")

	//获取设备信息
	uav := Service.GetUavsByStatesWithPage("free", typename, page, Const.PAGEMAX)

	//JSON格式返回
	c.JSON(200, &uav)
}

/* 优化
// GetNotUsedBattery 获取空闲的电池设备
func GetNotUsedBattery(c *gin.Context) {

	page := c.DefaultQuery("page", "0")

	//获取设备信息
	Battery := Model.GetUavsByStatesWithPage("free", "Battery", page, PAGEMAX)

	//JSON格式返回
	c.JSON(200, &Battery)
}

// GetNotUsedControl 获取空闲的遥控设备
func GetNotUsedControl(c *gin.Context) {

	page := c.DefaultQuery("page", "0")

	//获取设备信息
	Control := Model.GetUavsByStatesWithPage("free", "Control", page, PAGEMAX)

	//JSON格式返回
	c.JSON(200, &Control)
}
*/

// GetUsingDevices 获取使用中的所有设备
func GetUsingDevices(c *gin.Context) {
	equipmentType := c.Query("type")
	device := Service.GetUavByStates("using", equipmentType)

	c.JSON(200, &device)
}

// GetAllDevices 获取所有设备
func GetAllDevices(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	typename := c.DefaultQuery("type", "")

	device := Service.GetUavsByStatesWithPage("", typename, page, Const.PAGEMAX)

	c.JSON(200, &device)
}

// GetDevices 获取所有设备(前端指定状态和类型)
func GetDevices(c *gin.Context) {
	var uavs Model.SearchUav
	//结构体绑定
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}
	device := Service.GetUavByAll(uavs)

	c.JSON(200, &device)
}

// GetDeviceByUid 获取对应uid设备信息(用户)
func GetDeviceByUid(c *gin.Context) {
	id := c.Query("uid")

	flag, uav := Service.GetBackUavByUid(id)
	if flag {
		c.JSON(200, &uav)
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"desc": "未查找到该设备",
		})
	}

}

// AdminGetDeviceByUid 获取对应uid设备信息(管理员)
func AdminGetDeviceByUid(c *gin.Context) {
	id := c.Query("uid")

	uav := Service.GetUavByUid(id)
	uav.Img, _ = utils.GetPicUrl(uav.Img)
	c.JSON(200, &uav)
}

// GetDeviceByUids 获取对应uid设备信息
func GetDeviceByUids(c *gin.Context) {
	type searchuid struct {
		Uid []string `json:"uid"`
	}
	var searchuids searchuid

	//结构体绑定
	if err := c.BindJSON(&searchuids); err != nil {
		fmt.Println("获取对应uid设备信息 绑定失败", err.Error())
		c.JSON(401, gin.H{"code": 400, "desc": "绑定失败"})
		return
	}
	uavs := Service.GetUavsByUids(searchuids.Uid)

	c.JSON(200, uavs)

}

// GetUavType 获取设备类型列表
func GetUavType(c *gin.Context) {

	flag, types := Service.GetUavType()

	if flag {
		c.JSON(200, gin.H{"code": 200, "desc": "查询成功", "data": types})
	} else {
		c.JSON(200, gin.H{"code": 200, "desc": "查询失败"})
	}

}

// AddUavType 添加设备类型
func AddUavType(c *gin.Context) {
	//定义结构体
	var typename Model.UavType

	//绑定数据
	err := c.BindJSON(&typename)
	if err != nil {
		fmt.Println("添加设备类型 绑定失败", err.Error())
		c.JSON(401, gin.H{"code": 400, "desc": "绑定失败"})
		return
	}

	//添加数据
	if Service.AddUavType(typename.TypeName, typename.Remark) {
		_, newType := Service.GetUavType()
		c.JSON(200, gin.H{"code": 200, "desc": "增加成功", "data": newType})
	} else {
		_, newType := Service.GetUavType()
		c.JSON(200, gin.H{"code": 200, "desc": "增加失败", "data": newType})
	}
}

// RemoveUavType 删除设备类型
func RemoveUavType(c *gin.Context) {

	var typename Model.UavType

	//绑定数据
	err := c.BindJSON(&typename)
	if err != nil {
		fmt.Println("删除设备类型 绑定失败", err.Error())
		c.JSON(401, gin.H{"code": 400, "desc": "绑定失败"})
		return
	}

	//查询数据
	if Service.RemoveUavType(typename.TypeName) {
		_, newType := Service.GetUavType()
		c.JSON(200, gin.H{"code": 200, "desc": "删除成功", "data": newType})
	} else {
		_, newType := Service.GetUavType()
		c.JSON(200, gin.H{"code": 200, "desc": "删除失败", "data": newType})
	}
}
