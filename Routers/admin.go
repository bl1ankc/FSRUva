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
}
