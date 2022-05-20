package Routers

import "main/api"

func AdminRoute() {
	g := r.Group("/Admin", api.AuthRequired(), api.VerifyAdmin())
	{
		//审核设备展示
		g.GET("/GetUnderReviewUav", api.GetReview)

		//获取所有设备信息
		g.GET("/GetAllDevices", api.GetAllDevices)

		//提交新的设备
		g.POST("/UploadUav", api.UploadNewUav)

		//审核通过借用设备
		g.POST("/GetPassedUav", api.GetPassedUav)

		//审核不通过借用设备
		g.POST("/GetFailUav", api.GetFailUav)

		//审核通过归还设备
		g.POST("/BackPassedUav", api.BackPassedUav)

		//审核不通过归还设备
		g.POST("/BackFailUav", api.BackFailUav)

		//获取所有用户
		g.GET("/GetAllUsers", api.GetAllUsers)

		//获取所有历史记录
		g.GET("/GetAllRecords", api.GetAllRecords)

		//导出设备
		g.POST("/GetDevices", api.GetDevices)

		//通过uid获取设备信息（管理员）
		g.GET("/GetDevice", api.AdminGetDeviceByUid)

		//强制修改设备信息
		g.POST("/ForceUpdateDevices", api.ForceUpdateDevices)

		//修改设备备注信息
		g.POST("/UpdateUavRemark", api.UpdateUavRemark)

		//管理员设置
		g.GET("/SetAdmin", api.SetAdmin)

		//取消管理员
		g.GET("/DelAdmin", api.DelAdmin)

		//添加设备类型
		g.POST("/UavType", api.AddUavType)

		//删除设备类型
		g.DELETE("/UavType", api.RemoveUavType)

		//获取图片临时访问地址
		g.GET("/GetImgUrl", api.GetImgUrl)
	}

}
