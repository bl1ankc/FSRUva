package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Model"
)

//获取未被使用的无人机设备
func GetUnuseUva(c *gin.Context) {
	//获取设备信息
	testuva, err := Model.GetUvasByState("free", "uva")

	if err != nil {
		fmt.Println("dberror")
		return
	} else {
		fmt.Println("dbok")
	}

	//JSON格式返回
	c.JSON(200, &testuva)
}
