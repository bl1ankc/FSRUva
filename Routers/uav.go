package Routers

import (
	"main/api"
)

func UavRoute() {
	g := r.Group("/Uav")

	g.GET("/GetUav", api.GetNotUsedDrones)

	g.GET("/GetBattery", api.GetNotUsedBattery)

	g.GET("/GetControl", api.GetNotUsedControl)

	g.POST("/UploadUav", api.UploadNewUav)
}
