package Routers

import (
	"main/Controller"
	"main/Mid"
)

func ExcelInit() {
	g := r.Group("/Excel", Mid.AuthRequired(), Mid.VerifyAdmin())
	{
		//所有设备信息导出
		g.GET("/OutPutDevices", Controller.OutPutDevices)
		//所有用户及借用信息导出
		g.GET("/OutPutUserRecords", Controller.OutPutUserBorrowing)
		//单类型设备借用记录获取
		g.GET("/OutPutDeviceRecordByType", Controller.OutPutDeviceRecordByType)
		//单个设备借用记录
		g.GET("/OutPutDevice", Controller.OutPutDeviceRecord)
	}
}
