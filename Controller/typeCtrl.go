package Controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main/Model"
	"main/Service"
	"main/Service/Status"
	"main/utils"
	"strconv"
)

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
	var uavType Model.UavType
	var code int

	//绑定数据
	uavType.Remark = c.PostForm("remark")
	uavType.TypeName = c.PostForm("typename")
	DepartmentID, _ := strconv.Atoi(c.PostForm("departmentID"))

	//获取部门实例
	department, err := Service.GetDepartment(uint(DepartmentID))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "查询不到对应部门信息"))
		return
	} else if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "数据库操作错误"))
		return
	}

	//添加数据
	if err, typeID := Service.AddUavType(uavType, department); err != nil {
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
