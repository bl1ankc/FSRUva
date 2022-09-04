package Routers

import (
	"main/Controller"
	"main/Mid"
)

func UserRoute() {

	//登录
	r.POST("/login", Mid.Login)
	//上传用户信息
	r.POST("/UploadUser", Controller.UploadUser)
	//获取用户手机号
	//r.POST("/GetPhoneNumber", Controller.GetPhoneNumber)

	g := r.Group("/User", Mid.AuthRequired())
	{
		//预约设备
		g.POST("/BorrowUav", Controller.BorrowUav)

		//取走设备
		g.POST("/GetUav", Controller.GetUav)

		//归还设备
		g.POST("/BackUav", Controller.BackUav)

		//取消借用
		g.POST("/CancelBorrow", Controller.CancelBorrow)

		//取消归还
		g.POST("/CancelBack", Controller.CancelBack)

		//更新用户电话
		g.POST("/UpdateUserPhone", Controller.UpdateUserPhone)

		//更新用户密码
		g.POST("/UpdateUserPwd", Controller.UpdateUserPwd)

		//获取用户信息
		g.GET("/GetUser", Controller.GetUser)

		//查看个人正在借用中的设备
		g.GET("/GetOwnUsing", Controller.GetOwnUsing)

		//查看个人历史借用的设备
		g.GET("/GetOwnHistory", Controller.GetOwnHistory)
	}

}
