package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// GetNotUsedDrones 获取空闲的无人机设备
func GetNotUsedDrones(c *gin.Context) {
	//获取设备信息
	uav := Model.GetUavByStates("free", "Drone")

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
	uav := Model.GetUavByStates("", "Drone")

	c.JSON(200, &uav)
}

// GetBattery 获取所有电池设备
func GetBattery(c *gin.Context) {
	Battery := Model.GetUavByStates("", "Battery")

	c.JSON(200, &Battery)
}

// GetControl 获取所有遥控设备
func GetControl(c *gin.Context) {
	Control := Model.GetUavByStates("", "Control")

	c.JSON(200, &Control)
}

// GetUsingDevices 获取使用中的所有设备
func GetUsingDevices(c *gin.Context) {
	device := Model.GetUavByStates("using", "")

	c.JSON(200, &device)
}

// GetAllDevices 获取所有设备
func GetAllDevices(c *gin.Context) {
	device := Model.GetUavByStates("", "")

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
	device := Model.GetUavByAll(uavs)

	c.JSON(200, &device)
}

// GetDeviceByUid 获取对应uid设备信息
func GetDeviceByUid(c *gin.Context) {
	id := c.Query("uid")

	uav := Model.GetUavByUid(id)
	uav.Img, _ = Model.GetPicUrl(uav.Img)
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
	uavs, flag := Model.GetUavsByUids(searchuids.Uid)
	if flag {
		c.JSON(200, uavs)
	} else {
		c.JSON(200, gin.H{"code": 200, "desc": "查找失败"})
	}
}
