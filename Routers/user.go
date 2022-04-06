package Routers

import "main/api"

func UserRoute() {
	g := r.Group("/User")

	//借用设备
	g.POST("/BorrowUav", api.BorrowUav)

	//归还设备
	g.POST("/BackUav", api.BackUav)
}
