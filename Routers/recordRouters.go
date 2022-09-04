package Routers

import (
	"main/Controller"
	"main/Mid"
)

func RecordRoute() {
	g := r.Group("/Record", Mid.AuthRequired())
	{
		//查询用户借用记录
		g.GET("/GetRecordsByName", Controller.GetRecordsByUser)

		//查询设备借用记录
		g.GET("/GetRecordsByUid", Controller.GetRecordsByUva)
	}

}
