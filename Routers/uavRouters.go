package Routers

import "main/Controller"

func UavRoute() {
	g := r.Group("/Uav")
	{
		//可借用无人机展示
		g.GET("/GetUav", Controller.GetNotUsedDrones)

		//获取所有设备信息
		g.GET("/GetAllDevices", Controller.GetAllDevices)

		//可借用电池展示
		//g.GET("/GetBattery", Controller.GetNotUsedBattery)

		//可借用遥控展示
		//g.GET("/GetControl", Controller.GetNotUsedControl)

		//单独获取设备数据
		g.GET("/GetDevice", Controller.GetDeviceByUid)

		//获取设备组信息
		g.POST("/GetDeviceByUids", Controller.GetDeviceByUids)

		//获取设备类型列表
		g.GET("/GetUavTypeList", Controller.GetUavTypeList)

		//获取设备类型
		g.GET("/GetUavType", Controller.GetUavType)
	}
}
