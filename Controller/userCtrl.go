package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Const"
	"main/Model"
	"main/Service"
	"main/Service/Trans"
	"main/Service/Status"
)

// BorrowUav 借用设备
func BorrowUav(c *gin.Context) {
	//模型定义
	var data Model.Uav
	var code int
	//结构体绑定
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("绑定失败：", err.Error())
		code = Status.FailToBindJson
		c.JSON(code, R(code, nil, "传输数据失败"))
		return
	}
	//uav示例获取
	exist, uav := Service.GetUavByUid(data.Uid)
	if exist == false {
		c.JSON(200, R(200, nil, "该设备不存在"))
		return
	} else {
		data.ID = uav.ID
	}

	//user实例获取
	userID, _ := c.Get("userID")
	user, err := Service.GetUser(userID.(uint))
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "用户获取失败"))
		return
	}

	//表单中提交不可使用的无人机
	flag := false
	//再次验证是否能被借用
	if uav.State != "free" {
		flag = true
	} else {
		//借用事务
		if err := Trans.Borrow(data, user); err != nil {
			code = Status.FuncFail
			c.JSON(code, R(code, nil, "生成记录失败"))
			return
		}
	}

	//返回错误信息
	code = Status.OK
	if flag {
		c.JSON(code, R(code, nil, "设备已被借用"))
	} else {
		c.JSON(code, R(code, nil, "预约成功"))
	}
	return
}

// BackUav 归还设备
func BackUav(c *gin.Context) {
	//define
	var code int

	//获取id
	id, flag := c.GetQuery("uid")
	if !flag {
		code = Status.FailToGetQuery
		c.JSON(code, R(code, nil, "传入id失败"))
		return
	}

	//获取设备信息
	exist, uav := Service.GetUavByUid(id)
	if exist == false {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "未找到对应设备"))
		return
	}

	//user实例
	userID := c.MustGet("userID")
	user, err := Service.GetUser(userID.(uint))
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "用户获取失败"))
		return
	}

	//身份验证
	if uav.StudentID != user.StudentID {
		code = Status.UserAuthentication
		c.JSON(code, R(code, nil, "不可归还别人借用的设备"))
		return
	}

	//再次验证是否可以归还
	if uav.State != "using" {
		code = Status.ErrorControl
		c.JSON(code, R(code, nil, "设备已归还或者处于其他状态"))
		return
	}

	//上传图片
	if UploadImg(c, "Uav", "") == false {
		code = Status.UploadFail
		c.JSON(code, R(code, nil, "上传图片失败"))
		return
	}

	//归还
	if err := Trans.Back(&uav); err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "归还失败,函数处理异常"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "归还成功"))
	return
}

// GetUav 取走设备
func GetUav(c *gin.Context) {
	var code int

	//获取id
	id, flag := c.GetQuery("uid")

	//获取失败
	if !flag {
		c.JSON(400, gin.H{"code": 400, "desc": "传入id失败"})
		return
	}

	//获取实例
	exist, uav := Service.GetUavByUid(id)
	if !exist {
		code = Status.ErrorData
		c.JSON(code, R(code, nil, "查找对应设备失败,检查uid是否正确"))
		return
	}

	//上传图片
	if UploadImg(c, "Uav", "") == false {
		code = Status.ErrorControl
		c.JSON(code, R(code, nil, "错误操作,未上传取走图片"))
		return
	}

	//更新信息
	if err := Trans.Get(&uav); err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "BorrowGet事务处理错误"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "取走成功"))
	return
}

// CancelBorrow 取消借用
func CancelBorrow(c *gin.Context) {
	//获取id
	id, flag := c.GetQuery("uid")

	//获取失败
	if !flag {
		c.JSON(400, gin.H{"code": 400, "desc": "传入id失败"})
		return
	}

	//更新状态为审核中
	err := Service.UpdateState(id, "free")
	err = Service.UpdateRecordState(id, "cancelled")
	if err != nil {
		c.JSON(503, gin.H{"code": 503, "desc": "函数操作失败"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "desc": "取消成功"})
	return
}

// CancelBack 取消归还
func CancelBack(c *gin.Context) {
	//获取id
	id, flag := c.GetQuery("uid")

	//获取失败
	if !flag {
		c.JSON(400, gin.H{"code": 400, "desc": "传入id失败"})
		return
	}

	//更新状态为归还审核
	err := Service.UpdateState(id, "using")
	if err != nil {
		c.JSON(503, gin.H{"code": 503, "desc": "函数操作失败"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "desc": "取消成功"})
	return
}

// GetOwnUsing 查看个人正在借用中的设备
func GetOwnUsing(c *gin.Context) {
	stuid, flag := c.GetQuery("stuid")

	page := c.DefaultQuery("page", "0")

	//获取失败
	if !flag {
		c.JSON(400, gin.H{"code": 400, "desc": "传入stuid失败"})
		return
	}

	uavs, flag := Service.GetUsingUavsByStuID(stuid, page, Const.PAGEMAX)
	if flag {
		c.JSON(200, &uavs)
	} else {
		c.JSON(502, gin.H{"code": 502, "message": "查询失败"})
	}

}

// GetOwnHistory 查看历史借用中的设备
func GetOwnHistory(c *gin.Context) {
	var code int
	stuid, flag := c.GetQuery("stuid")

	page := c.DefaultQuery("page", "0")
	//获取失败
	if !flag {
		c.JSON(400, gin.H{"code": 400, "desc": "传入stuid失败"})
		return
	}

	uavs, flag := Service.GetHistoryUavsByStuID(stuid, page, Const.PAGEMAX)
	if !flag {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "查询失败"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, uavs, "查询成功"))
	return

}
