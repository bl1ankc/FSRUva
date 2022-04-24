package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// BorrowUav 借用设备
func BorrowUav(c *gin.Context) {
	//模型定义
	var uav Model.BorrowUav

	//结构体绑定
	if err := c.ShouldBindJSON(&uav); err != nil {
		fmt.Println("绑定失败")
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
		Model.UpdateBorrower(uav.Uid, uav.Borrower, uav.Phone)
		Model.RecordBorrow(uav.Uid, uav.StudentID, uav.Borrower, uav.Plan_time, uav.Usage) //用途
		Model.UpdateUavUsage(uav.Uid, uav.Usage)
	}

	//返回错误信息
	if flag {
		c.JSON(200, gin.H{"code": 200, "desc": "设备已被借用"})
	} else {
		c.JSON(200, gin.H{"code": 200, "desc": "借用成功"})
	}
}

// BackUav 归还设备
func BackUav(c *gin.Context) {
	//获取id
	id := c.Query("uid")

	//上传图片
	if UploadImg(c) == false {
		return
	}

	//更新状态为归还审核
	Model.UpdateState(id, "Back under review")
	Model.UpdateImgInRecord(id, "back_img")
}

// GetUav 取走设备
func GetUav(c *gin.Context) {
	//获取id
	id := c.Query("uid")

	//上传图片
	if UploadImg(c) == false {
		return
	}

	//更新对应设备状态
	Model.UpdateState(id, "using")
	Model.UpdateImgInRecord(id, "get_img")
}

// CancelBorrow 取消借用
func CancelBorrow(c *gin.Context) {
	//模型定义
	var uav Model.Uav

	//结构体绑定
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态为审核中
	Model.UpdateState(uav.Uid, "free")
	Model.UpdateRecordState(uav.Uid, "cancelled")

}

// CancelBack 取消归还
func CancelBack(c *gin.Context) {
	//模型
	var uav Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态为归还审核
	Model.UpdateState(uav.Uid, "using")

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
			c.JSON(500, gin.H{"code": 500, "desc": "保存图片失败"})
			return false
		}
	} else {
		c.JSON(400, gin.H{"code": 400, "desc": "未上传图片"})
		return false
	}

	//上传成功
	Model.UpdateImg(c.Query("uid"), filename)
	c.JSON(200, gin.H{"code": 200, "desc": "上传图片成功"})
	return true
}
