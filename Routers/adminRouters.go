package Routers

import (
	"main/Controller"
	"main/Mid"
)

func AdminRoute() {
	g := r.Group("/Admin", Mid.AuthRequired(), Mid.VerifyAdmin())
	{
		//审核设备展示
		g.GET("/GetUnderReviewUav", Controller.GetReview)

		//提交新的设备
		g.POST("/UploadUav", Controller.UploadNewUav)

		//更新设备	@2022/09/04 b1ank
		g.POST("/UpdateDevice", Controller.UpdateDevice)

		//删除设备	@2022/09/04 b1ank
		g.POST("/DeleteDevice", Controller.DeleteDevice)

		//审核通过借用设备
		g.POST("/GetPassedUav", Controller.GetPassedUav)

		//审核不通过借用设备
		g.POST("/GetFailUav", Controller.GetFailUav)

		//审核通过归还设备
		g.POST("/BackPassedUav", Controller.BackPassedUav)

		//审核不通过归还设备
		g.POST("/BackFailUav", Controller.BackFailUav)

		//获取所有用户
		g.GET("/GetAllUsers", Controller.GetAllUsers)

		//获取所有历史记录
		g.GET("/GetAllRecords", Controller.GetAllRecords)

		//导出设备
		g.POST("/GetDevices", Controller.GetDevices)

		//通过uid获取设备信息（管理员）
		g.GET("/GetDevice", Controller.AdminGetDeviceByUid)

		//强制修改设备信息
		g.POST("/ForceUpdateDevices", Controller.ForceUpdateDevices)

		//修改设备备注信息
		g.POST("/UpdateUavRemark", Controller.UpdateUavRemark)

		//管理员设置
		g.GET("/SetAdmin", Controller.SetAdmin)

		//取消管理员
		g.GET("/DelAdmin", Controller.DelAdmin)

		//管理员类型设置 @2022/10/01 b1ank
		g.GET("/ChangeAdminType", Controller.ChangeAdminType)

		//添加设备类型
		g.POST("/AddUavType", Controller.AddUavType)

		//更新设备类型图片
		g.POST("/UploadUavTypeImg", Controller.UploadUavTypeImg)

		//删除设备类型
		g.DELETE("/RemoveUavType", Controller.RemoveUavType)

		//更新设备类型
		g.POST("/UpdateUavType", Controller.UpdateUavType)

		//获取图片临时访问地址
		g.GET("/GetImgUrl", Controller.GetImgUrl)
	}

}
