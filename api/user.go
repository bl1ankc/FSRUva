package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Model"
	"mime/multipart"
)

// BorrowUav 借用设备
func BorrowUav(c *gin.Context) {
	//模型定义
	var uav Model.BorrowUav

	//结构体绑定
	if err := c.ShouldBindJSON(&uav); err != nil {
		fmt.Println("绑定失败：", err.Error())
		c.JSON(400, gin.H{"code": 400, "desc": "传输数据失败"})
		return
	}

	//表单中提交不可使用的无人机
	flag := false

	//更新状态为审核中

	//再次验证是否能被借用
	if uav.GetUavStateByUid() != "free" {
		flag = true
	} else {
		Model.UpdateState(uav.Uid, "Get under review")
		Model.UpdateBorrower(uav.Uid, uav.Borrower, uav.Phone, uav.StudentID)
		Model.UpdatePlanTime(uav.Uid, uav.Plan_time)
		Model.RecordBorrow(uav.Uid, uav.StudentID, uav.Borrower, uav.Plan_time, uav.Usage) //用途
		Model.UpdateUavUsage(uav.Uid, uav.Usage)
	}

	//返回错误信息
	if flag {
		c.JSON(200, gin.H{"code": 200, "desc": "设备已被借用"})
	} else {
		c.JSON(200, gin.H{"code": 200, "desc": "预约成功"})
	}
}

// BackUav 归还设备
func BackUav(c *gin.Context) {
	//获取id
	id, flag := c.GetQuery("uid")
	fmt.Println(id)

	//获取失败
	if !flag {
		c.JSON(400, gin.H{"code": 400, "desc": "传入id失败"})
		return
	}

	//上传图片
	if UploadImg(c) == false {
		c.JSON(200, gin.H{"desc": "图片上传失败"})
		return
	}

	//更新状态为归还审核
	Model.UpdateState(id, "Back under review")
	Model.UpdateImgInRecord(id, "back_img")
	Model.UpdateRecordState(id, "Back under review")
	Model.UpdateBackRecord(id)
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
	if UploadImg(c) == false {
		return
	}

	//更新对应设备状态
	Model.UpdateState(id, "using")
	Model.UpdateImgInRecord(id, "get_img")
	Model.UpdateRecordState(id, "using")
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
	Model.UpdateState(id, "free")
	Model.UpdateRecordState(id, "cancelled")

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
	Model.UpdateState(id, "using")

}

// UploadImg 上传图片
func UploadImg(c *gin.Context) bool {
	/*
		上传图片,并保存在./img服务器文件夹中
		自动生成图片对应的uid,该uid为文件名,并且绑定对应记录
	*/

	//上传图片
	file, _ := c.FormFile("upload_img")
	filename := Model.GetUid() + ".png"

	//上传失败
	if file != nil {
		if err := c.SaveUploadedFile(file, "./img/"+filename); err != nil {
			//c.JSON(500, gin.H{"code": 500, "desc": "保存图片失败"})
			return false
		}
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
	if !Model.UploadImgToOSS("img/"+filename, src) {
		fmt.Println("OSS上传失败")
	}

	//上传成功
	Model.UpdateImg(c.Query("uid"), filename)
	fmt.Println("OSS上传成功")
	//c.JSON(200, gin.H{"code": 200, "desc": "上传图片成功", "src": "/img/" + filename})
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

	uavs, flag := Model.GetUsingUavsByStuID(stuid, page, PAGEMAX)
	if flag {
		c.JSON(200, &uavs)
	} else {
		c.JSON(502, gin.H{"code": 502, "message": "查询失败"})
	}

}

// GetOwnHistory 查看历史借用中的设备
func GetOwnHistory(c *gin.Context) {
	stuid, flag := c.GetQuery("stuid")

	page := c.DefaultQuery("page", "0")
	//获取失败
	if !flag {
		c.JSON(400, gin.H{"code": 400, "desc": "传入stuid失败"})
		return
	}

	uavs, flag := Model.GetHistoryUavsByStuID(stuid, page, PAGEMAX)
	if flag {
		c.JSON(200, &uavs)
	} else {
		c.JSON(502, gin.H{"code": 502, "message": "查询失败"})
	}

}
