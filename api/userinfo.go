package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Model"
)

/*
	注册
*/

// UploadUser 上传用户信息ljy
func UploadUser(c *gin.Context) {
	//模型定义
	var user Model.User

	//绑定结构体
	if err := c.BindJSON(&user); err != nil {
		fmt.Println("绑定失败")
		c.JSON(400, gin.H{"code": 400, "desc": "传输数据失败"})
		return
	}
	//数据插入
	flag, response := Model.InsertUser(user.Name, user.Phone, user.StudentID, user.Pwd)
	if flag {
		c.JSON(200, gin.H{"code": 200, "desc": response})
	} else {
		c.JSON(200, gin.H{"code": 200, "desc": response})
	}

}

/*
	更改信息 获取信息
*/

// UpdateUserPhone 更新电话
func UpdateUserPhone(c *gin.Context) {
	var user Model.User

	//结构体绑定
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//数据插入

	if Model.UpdatePhone(user.StudentID, user.Phone) {
		c.JSON(200, gin.H{"code": 200, "message": "电话更改成功"})
	} else {
		c.JSON(502, gin.H{"code": 502, "message": "电话更改失败"})
	}
}

// UpdateUserPwd 更改密码
func UpdateUserPwd(c *gin.Context) {
	type UpdatePwd struct {
		StudentID string `json:"stuid"`  //学号
		OldPwd    string `json:"oldpwd"` //旧密码
		NewPwd    string `json:"newpwd"` //新密码
	}
	var user UpdatePwd

	//结构体绑定
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//数据更新
	resopnse, success := Model.UpdatePwd(user.StudentID, user.OldPwd, user.NewPwd)
	if success {
		c.JSON(200, gin.H{"code": 200, "message": resopnse})
	} else {
		c.JSON(502, gin.H{"code": 502, "message": resopnse})
	}
}

// GetUser 获取单个用户信息
func GetUser(c *gin.Context) {
	//数据绑定
	id := c.Query("stuid")

	//数据获取
	response := Model.GetUserByID(id)
	c.JSON(200, &response)
}

// SetAdmin 设置管理员
func SetAdmin(c *gin.Context) {
	//数据绑定
	id := c.Query("stuid")

	//数据更改
	Model.UpdateAdmin(id, true)
}

// DelAdmin 取消管理员
func DelAdmin(c *gin.Context) {
	//数据绑定
	id := c.Query("stuid")

	//数据更改
	Model.UpdateAdmin(id, false)
}
