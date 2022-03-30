package main

import (
	"main/Routers"
)

func main() {

	//展示界面获取可用的无人机相关设备
	r := Routers.InitRouter()

	r.Run()
}
