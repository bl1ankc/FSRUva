package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Model"
	"main/Service"
	"main/Service/Status"
	"main/utils"
)

// GetNotUsedDrones 获取空闲的设备 @2023/3/14
func GetNotUsedDrones(c *gin.Context) {
	var code int
	var response struct {
		Devices []Model.Uav
		Total   int64
	}

	typename := c.Query("type")

	//获取设备信息
	if uav, total, err := Service.GetUavsByStatesWithPage("free", typename, c.Request); err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "获取数据失败"))
		return
	} else {
		response.Devices = uav
		response.Total = total
	}

	//JSON格式返回
	code = Status.OK
	c.JSON(code, R(code, response, "获取数据成功"))
	return
}

// GetUsingDevices 获取使用中的所有设备
func GetUsingDevices(c *gin.Context) {
	equipmentType := c.Query("type")
	device := Service.GetUavByStates("using", equipmentType)

	c.JSON(200, &device)
}

// GetAllDevices 获取所有设备 @2023/3/14
func GetAllDevices(c *gin.Context) {
	var code int
	var response struct {
		Devices []Model.Uav
		Total   int64
	}

	typename := c.DefaultQuery("type", "")

	if device, total, err := Service.GetUavsByStatesWithPage("", typename, c.Request); err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "获取数据失败"))
		return
	} else {
		response.Devices = device
		response.Total = total
	}
	code = Status.OK
	c.JSON(code, R(code, response, "获取数据成功"))
	return
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
