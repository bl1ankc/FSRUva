package Routers

import (
	"main/Mid"
	"main/utils"
)

func RFIDRoute() {
	g := r.Group("/RFID", Mid.AuthRequired())
	{
		//
		g.GET("/getId", utils.GetID)
	}

}
