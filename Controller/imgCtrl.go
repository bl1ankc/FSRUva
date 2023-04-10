package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Service"
	"main/utils"
	"mime/multipart"
	"strconv"
)

// UploadImg 上传图片
func UploadImg(c *gin.Context, imgType string, id interface{}) bool {
	/*
		上传图片,保存路径在obs中对应bucket.endpoint/object路径
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
		fmt.Println("文件流转换失败")
		return false
	}

	//关闭失败
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			fmt.Println("文件关闭失败")
		}
	}(src)

	//云端上传
	filePath, err := utils.UploadFileToObs(filename, src)
	if err != nil {
		fmt.Println("文件上传obs失败")
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
		if err = Service.UpdateUavImg(uid, filePath); err != nil {
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
		if err = Service.UpdateTypeImg(typeID, filePath); err != nil {
			return false
		}
	}
	fmt.Println("Obs上传成功")
	return true
}
