package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
	"main/Service"
	"main/utils"
)

// GetNotUsedDrones 获取空闲的设备 @2023/3/3(paginate update)
func GetNotUsedDrones(c *gin.Context) {

	typename := c.Query("type")

	//获取设备信息
	uav := Service.GetUavsByStatesWithPage("free", typename, c.Request)

	//JSON格式返回
	c.JSON(200, &uav)
}

// GetUsingDevices 获取使用中的所有设备
func GetUsingDevices(c *gin.Context) {
	equipmentType := c.Query("type")
	device := Service.GetUavByStates("using", equipmentType)

	c.JSON(200, &device)
}

// GetAllDevices 获取所有设备
func GetAllDevices(c *gin.Context) {

	typename := c.DefaultQuery("type", "")

	device := Service.GetUavsByStatesWithPage("", typename, c.Request)

	c.JSON(200, &device)
}

// GetDevices 获取所有设备(前端指定状态和类型)
func GetDevices(c *gin.Context) {
	var uavs Model.Uav
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

	flag, uav := Service.GetUavByUid(id)
	if flag {
		c.JSON(200, &uav)
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"desc": "未查找到该设备",
		})
	}

}

//AdminGetDeviceByUid 获取对应uid设备信息(管理员)
func AdminGetDeviceByUid(c *gin.Context) {
	id := c.Query("uid")

	_, uav := Service.GetUavByUid(id)
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
