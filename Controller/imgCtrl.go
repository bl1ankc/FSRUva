package Controller

import (
	"github.com/gin-gonic/gin"
	"main/utils"
)

// GetImgUrl 获取图片临时地址
func GetImgUrl(c *gin.Context) {
	imgName := c.Query("imgName")

	url, flag := utils.GetPicUrl(imgName + ".png")

	//绑定结构体
	type T struct {
		Url string `json:"url"`
	}
	resp := &T{Url: url}

	if flag {
		c.JSON(200, gin.H{"code": 200, "desc": "获取成功", "data": resp})
	} else {
		c.JSON(200, gin.H{"code": 200, "desc": "获取失败"})
	}
}
