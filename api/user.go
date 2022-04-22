package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// BorrowUav 借用设备
func BorrowUav(c *gin.Context) {
	//模型定义
	var uavs []Model.BorrowUav

	//结构体绑定
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//表单中提交不可使用的无人机
	flag := false
	var erruav []Model.Uav
	var Uids []string

	//更新状态为审核中
	for _, uav := range uavs {
		//再次验证是否能被借用
		if uav.GetUavStateByUid() != "free" {
			flag = true
			Uids = append(Uids, uav.Uid)
			continue
		}
		Model.UpdateState(uav.Uid, "Get under review")
		Model.UpdateBorrower(uav.Uid, uav.Borrower, uav.Phone)
		Model.UpdatePlanTime(uav.Uid, uav.Plan_time)
		Model.RecordBorrow(uav.Uid, uav.Borrower, uav.Get_time, uav.Plan_time, uav.Usage) //用途
	}
	erruav = Model.GetUavsByUids(Uids) //返回设备此时的状态信息
	//返回错误信息
	if flag {
		c.JSON(200, erruav)
	} else {
		c.JSON(200, "OK")
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
}

// CancelBorrow 取消借用
func CancelBorrow(c *gin.Context) {
	//模型定义
	var uavs []Model.Uav

	//结构体绑定
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态为审核中
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "free")
		Model.UpdateRecordState(uav.Uid, "cancelled")
	}

}

// CancelBack 取消归还
func CancelBack(c *gin.Context) {
	//模型
	var uavs []Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uavs); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态为归还审核
	for _, uav := range uavs {
		Model.UpdateState(uav.Uid, "using")
	}
}

// UploadImg 上传图片
func UploadImg(c *gin.Context) bool {
	/*
		上传图片,并保存在./img服务器文件夹中
		自动生成图片对应的uid,该uid为文件名,并且绑定对应记录
	*/

	//上传图片
	file, _ := c.FormFile("upload_img")
	filename := Model.GetUid()

	//上传失败
	if file != nil {
		if err := c.SaveUploadedFile(file, "./img"); err != nil {
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
