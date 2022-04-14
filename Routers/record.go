package Routers

import (
	"main/api"
)

func RecordRoute() {
	g := r.Group("/Record")

	//查询用户借用记录
	g.GET("/GetRecordsByName", api.GetRecordsByUser)

	//查询设备借用记录
	g.GET("/GetRecordsByUid", api.GetRecordsByUva)

}
