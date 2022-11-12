package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/Const"
	"main/Model"
	"main/Service"
	"main/Service/Status"
	"main/utils"
	"strconv"
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

/*
	类型操作
*/

// GetUavTypeList 获取设备类型列表
func GetUavTypeList(c *gin.Context) {

	flag, types := Service.GetUavType()

	if flag {
		c.JSON(200, gin.H{"code": 200, "desc": "查询成功", "data": types})
	} else {
		c.JSON(200, gin.H{"code": 200, "desc": "查询失败"})
	}

}

// GetUavType 名称获取单个设备类型
func GetUavType(c *gin.Context) {
	var code int

	typeName := c.Query("typeName")
	uavType, err := Service.GetTypeByName(typeName)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "获取设备类型失败，服务器报错"))
		return
	}
	uavType.Img, _ = utils.GetPicUrl(uavType.Img)
	code = Status.OK
	c.JSON(code, R(code, uavType, "获取成功"))
	return
}

// AddUavType 添加设备类型
func AddUavType(c *gin.Context) {
	//定义结构体
	var typename Model.UavType
	var code int

	//绑定数据
	remark := c.PostForm("remark")
	typeName := c.PostForm("typename")
	typename.TypeName = typeName
	typename.Remark = remark

	//添加数据
	if err, typeID := Service.AddUavType(typename.TypeName, typename.Remark); err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "添加设备类型失败"))
		return
	} else {
		file, _ := c.FormFile("upload_img")
		if file != nil && UploadImg(c, "UavType", typeID) == false {
			code = Status.OBSErr
			c.JSON(code, R(code, nil, "上传图片出错"))
			return
		}
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "添加类型成功"))
	return

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

// UpdateUavType 更新设备类型
func UpdateUavType(c *gin.Context) {
	var code int
	var uavType Model.UavType

	typeName := c.PostForm("typeName")
	remark := c.PostForm("remark")
	v, _ := strconv.Atoi(c.PostForm("id"))
	id := uint(v)

	uavType.ID = id
	uavType.TypeName = typeName
	uavType.Remark = remark

	if err := Service.UpdateUavType(uavType); err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "更新失败"))
		return
	}

	file, _ := c.FormFile("upload_img")
	if file != nil && UploadImg(c, "UavType", id) == false {
		code = Status.OBSErr
		c.JSON(code, R(code, nil, "上传图片出错"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "更新成功"))
	return

}

// UploadUavTypeImg 添加设备类型图片
func UploadUavTypeImg(c *gin.Context) {
	var code int

	file, _ := c.FormFile("upload_img")
	if file != nil && UploadImg(c, "UavType", "") == false {
		code = Status.OBSErr
		c.JSON(code, R(code, nil, "上传图片出错"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "上传图片成功"))
	return

}
