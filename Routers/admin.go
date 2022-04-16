package Routers

import "main/api"

func AdminRoute() {
	g := r.Group("/Admin")

	//借用审核设备展示
	g.GET("/GetUnderReviewUav", api.GetUnderReview)

	//归还审核设备展示
	g.GET("/BackUnderReviewUav", api.BackUnderReview)

	//提交新的设备
	g.POST("/UploadUav", api.UploadNewUav)

	//审核通过借用设备
	g.POST("/GetPassedUav", api.GetPassedUav)

	//审核不通过借用设备
	g.POST("/GetFailUav", api.GetFailUav)

	//审核通过归还设备
	g.POST("/BackPassedUav", api.BackPassedUav)

	//获取所有用户
	g.GET("/GetAllUsers", api.GetAllUsers)

	//获取所有历史记录
	g.GET("/GetAllRecords", api.GetAllRecords)

	//获取设备信息
	g.POST("/GetDevices", api.GetDevices)

	//强制修改设备信息
	g.POST("/ForceUpdateDevices", api.ForceUpdateDevices)

	//修改设备备注信息
	g.POST("/UpdateUavRemark", api.UpdateUavRemark)
}
