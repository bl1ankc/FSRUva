package Routers

import (
	"main/api"
)

func UserRoute() {
	g := r.Group("/user")

	g.GET("/GetUva", api.GetUnuseUva)

}
