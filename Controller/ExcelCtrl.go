package Controller

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"main/Const"
	"main/Service"
	"main/Service/Status"
	"main/utils"
)

// OutPutDevices @2023/3/23
func OutPutDevices(c *gin.Context) {
	var code int

	//获取数据信息
	data, err := Service.GetDeviceData()
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "获取设备数据失败"))
		return
	}

	//补充部门信息
	for index, device := range data {
		uavType, err := Service.GetTypeByName(device.Type)
		if err != nil {
			continue
			//code = Status.ErrorData
			//c.JSON(code, R(code, nil, "设备类型获取失败"))
			//return
		}

		department, err := Service.GetDepartment(uavType.DepartmentID)
		if err != nil {
			continue
			//code = Status.FuncFail
			//c.JSON(code, R(code, nil, "获取部门实例失败"))
			//return
		}
		data[index].Department = department.DepartmentName
	}

	//类型泛化
	re, err := utils.ToDoubleInterfaceSlice(data)
	if err == errors.New(Const.ErrorDataType) {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "错误数据类型"))
	}

	//文件头
	headers := []string{"ID", "设备名称", "设备状态", "设备部门", "设备类型", "设备序列号", "设备存放位置", "设备备注", "是否贵重"}

	file, err := utils.ExportExcel("设备信息", headers, re)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "打印失败"))
		return
	}

	//输出到缓冲区中
	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())
	//返回到前端
	utils.ResponseXls(c, content, "DeviceInformation")
}

// OutPutUserBorrowing @2023/3/24
func OutPutUserBorrowing(c *gin.Context) {
	var code int

	//获取数据
	data, err := Service.GetUserRecords()
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "获取用户及记录数据失败"))
		return
	}

	//类型泛化
	re, err := utils.ToDoubleInterfaceSlice(data)
	if err == errors.New(Const.ErrorDataType) {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "错误数据类型"))
	}

	//文件头
	headers := []string{"借用人", "学号", "记录ID", "设备名称", "设备状态", "设备类型", "用途", "借出时间", "归还时间"}

	file, err := utils.ExportExcel("借用记录", headers, re)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "打印失败"))
		return
	}

	//输出到缓冲区中
	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())
	//返回到前端
	utils.ResponseXls(c, content, "UserBorrowingRecord")
}

// OutPutDeviceRecordByType @2023/3/24
func OutPutDeviceRecordByType(c *gin.Context) {
	var code int
	typeName := c.Query("typeName")

	//获取数据
	data, err := Service.GetDeviceRecordByType(typeName)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "获取用户及记录数据失败"))
		return
	}

	//类型泛化
	re, err := utils.ToDoubleInterfaceSlice(data)
	if err == errors.New(Const.ErrorDataType) {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "错误数据类型"))
	}

	//文件头
	headers := []string{"借用人", "学号", "记录ID", "设备名称", "设备状态", "设备类型", "用途", "借出时间", "归还时间"}

	file, err := utils.ExportExcel(typeName, headers, re)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "打印失败"))
		return
	}

	//输出到缓冲区中
	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())
	//返回到前端
	utils.ResponseXls(c, content, "OneTypeDeviceRecord")
}

// OutPutDeviceRecord @2023/3/24
func OutPutDeviceRecord(c *gin.Context) {
	var code int
	uid := c.Query("uid")

	//实例获取
	exist, uav := Service.GetUavByUid(uid)
	if !exist {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "未找到设备,uid错误"))
		return
	}

	//数据获取
	data, err := Service.GetDeviceRecord(uav)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "获取记录失败"))
		return
	}

	//类型泛化
	re, err := utils.ToDoubleInterfaceSlice(data)
	if err == errors.New(Const.ErrorDataType) {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "错误数据类型"))
	}

	//文件头
	headers := []string{"借用人", "学号", "记录ID", "设备名称", "设备状态", "设备类型", "用途", "借出时间", "归还时间"}

	file, err := utils.ExportExcel("设备记录", headers, re)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "打印失败"))
		return
	}

	//输出到缓冲区中
	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())
	//返回到前端
	utils.ResponseXls(c, content, "DeviceRecord")

}
