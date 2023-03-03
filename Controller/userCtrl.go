package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Const"
	"main/Model"
	"main/Service"
	"main/Service/Status"
	"main/utils"
	"mime/multipart"
	"strconv"
	"time"
)

// BorrowUav 借用设备
func BorrowUav(c *gin.Context) {
	//模型定义
	var uav Model.Uav
	var err error
	//结构体绑定
	if err = c.ShouldBindJSON(&uav); err != nil {
		fmt.Println("绑定失败：", err.Error())
		c.JSON(400, gin.H{"code": 400, "desc": "传输数据失败"})
		return
	}

	exist, tmp := Service.GetUavByUid(uav.Uid)
	expensive := tmp.Expensive
	if exist == false {
		c.JSON(200, R(200, nil, "该设备不存在"))
		return
	}

	uav.Expensive = tmp.Expensive
	//表单中提交不可使用的无人机
	flag := false

	//再次验证是否能被借用
	if Service.GetUavStateByUid(uav) != "free" {
		flag = true
	} else {
		Service.RecordBorrow(uav) //用途
		if expensive != true {    //非贵重直接跳到预约成功
			uav.State = "scheduled"
			uav.GetTime = time.Now().Local()
			Service.UpdateRecordState(uav.Uid, "scheduled")
		} else {
			uav.State = "Get under review"
		}
		fmt.Println(uav)
		fmt.Println(tmp)
		//更新设备信息
		if err = Service.UpdateDevice(uav); err != nil {
			c.JSON(401, R(401, nil, "更新函数错误"))
			return
		}
	}

	//返回错误信息
	if flag {
		c.JSON(200, gin.H{"code": 200, "desc": "设备已被借用"})
	} else {
		c.JSON(200, gin.H{"code": 200, "desc": "预约成功"})
	}
	return
}

// BackUav 归还设备
func BackUav(c *gin.Context) {
	//获取id
	id, flag := c.GetQuery("uid")

	//获取失败
	if !flag {
		c.JSON(400, gin.H{"code": 400, "desc": "传入id失败"})
		return
	}

	//获取设备信息
	exist, uav := Service.GetUavByUid(id)
	if exist == false {
		c.JSON(400, gin.H{"code": 400, "desc": "未找到对应设备"})
		return
	}
	////获取记录实例
	//exist, record := Service.GetRecordById(uav.RecordID)
	//if exist == false {
	//	c.JSON(400, gin.H{"code": 400, "desc": "未找到对应记录"})
	//	return
	//}

	if uav.StudentID != c.MustGet("studentid") {
		c.JSON(403, gin.H{"code": 403, "desc": "不可归还别人借用的设备"})
		return
	}
	//再次验证是否可以归还
	if uav.State != "using" {
		c.JSON(403, gin.H{"code": 403, "desc": "设备不处于使用中状态，不可归还"})
		return
	}

	//上传图片
	if UploadImg(c, "Uav", "") == false {
		c.JSON(200, gin.H{"code": 200, "desc": "图片上传失败"})
		return
	}

	//更新状态为归还审核
	err := Service.UpdateState(id, "Back under review")
	exist = Service.UpdateImgInRecord(id, "back_img")
	if err != nil || exist == false {
		c.JSON(503, gin.H{"code": 503, "desc": "函数操作错误"})
		return
	}
	err = Service.UpdateRecordState(id, "Back under review")
	exist = Service.UpdateBackRecord(id)
	if err != nil || exist == false {
		c.JSON(503, gin.H{"code": 503, "desc": "函数操作错误"})
		return
	}
	c.JSON(200, gin.H{"desc": "归还成功"})
}

// GetUav 取走设备
func GetUav(c *gin.Context) {
	//获取id
	id, flag := c.GetQuery("uid")

	//获取失败
	if !flag {
		c.JSON(400, gin.H{"code": 400, "desc": "传入id失败"})
		return
	}

	//上传图片
	if UploadImg(c, "Uav", "") == false {
		return
	}

	//更新对应设备状态
	err := Service.UpdateState(id, "using")
	err = Service.UpdateBorrowTime(id, time.Now().Local())
	err = Service.GetReviewRecord(id, "", "", "", time.Now().Local())
	exist := Service.UpdateImgInRecord(id, "get_img")
	err = Service.UpdateRecordState(id, "using")

	if err != nil || exist != true {
		c.JSON(503, gin.H{"code": 503, "desc": "函数操作失败"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "desc": "取走成功"})
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

// UploadImg 上传图片
func UploadImg(c *gin.Context, imgType string, id interface{}) bool {
	/*
		上传图片,并保存在./img服务器文件夹中
		自动生成图片对应的uid,该uid为文件名,并且绑定对应记录
	*/

	//上传图片
	file, _ := c.FormFile("upload_img")
	filename := utils.GetUid() + ".png"

	//上传失败
	if file != nil {
		/* 保存到本地
		if err := c.SaveUploadedFile(file, "./Data/"+filename); err != nil {
			//c.JSON(500, gin.H{"code": 500, "desc": "保存图片失败"})
			return false
		}
		*/
	} else {
		//c.JSON(400, gin.H{"code": 400, "desc": "未上传图片"})
		return false
	}

	//转换失败
	src, err := file.Open()
	if err != nil {
		fmt.Println("OSS文件转换失败")
		return false
	}

	//关闭失败
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			fmt.Println("OSS文件关闭失败")
		}
	}(src)

	//云端上传
	if !utils.UploadImgToOSS("Data/"+filename, src) {
		fmt.Println("OSS上传失败")
		return false
	}

	//上传成功
	if imgType == "Uav" {
		uid := c.Query("uid") //前端携带
		if uid == "" {        //前端未携带，检查是否有参数
			if value, ok := id.(string); ok {
				uid = value
			} else {
				return false
			}
		}
		if err = Service.UpdateUavImg(c.Query("uid"), filename); err != nil {
			return false
		}

	} else if imgType == "UavType" {
		v, err := strconv.Atoi(c.Query("typeID"))
		typeID := uint(v)
		if err != nil {
			fmt.Println("无参")
			if value, ok := id.(uint); ok {
				typeID = value
			} else {
				return false
			}
		}
		if err = Service.UpdateTypeImg(typeID, filename); err != nil {
			return false
		}
	}
	fmt.Println("OSS上传成功")
	//c.JSON(200, gin.H{"code": 200, "desc": "上传图片成功", "src": "/Data/" + filename})
	return true
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
