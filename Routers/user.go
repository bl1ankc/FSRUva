package Routers

import "main/api"

func UserRoute() {
	g := r.Group("/User")

	//预约设备
	g.POST("/BorrowUav", api.BorrowUav)

	//取走设备
	g.POST("/GetUav", api.GetUav)

	//归还设备
	g.POST("/BackUav", api.BackUav)

	//取消借用
	g.POST("/CancelBorrow", api.CancelBorrow)

	//取消归还
	g.POST("/CancelBack", api.CancelBack)

	//上传用户信息
	g.POST("/UploadUser", api.UploadUser)

	//更新用户电话
	g.POST("/UpdateUserPhone", api.UpdateUserPhone)

	//更新用户密码
	g.POST("/UpdateUserPwd", api.UpdateUserPwd)

	//获取用户信息
	g.GET("/GetUser", api.GetUser)

	//查看个人正在借用中的设备
	g.GET("/GetOwnUsing", api.GetUsingDevices)
}
