package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Model"
)

// UploadNewUav 上传新设备
func UploadNewUav(c *gin.Context) {
	//模型定义
	var uav Model.Uav

	//结构体绑定
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}

	//数据插入
	Model.InsertUva(uav.Name, uav.Type, uav.Uid)

}

// GetReviewUav 借用审核设备
func GetReviewUav(c *gin.Context) {
	//模型定义
	var uav Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态
	Model.UpdateState(uav.Uid, "borrowing")
}

// BackReviewUav 借用审核设备
func BackReviewUav(c *gin.Context) {
	//模型定义
	var uav Model.Uav

	//绑定结构体
	if err := c.BindJSON(&uav); err != nil {
		log.Fatal(err.Error())
		return
	}

	//更新状态
	Model.UpdateState(uav.Uid, "free")
}
