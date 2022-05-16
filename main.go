package main

import (
	"main/Routers"
)

func main() {
	//初始化定时任务
	//c := cron.InitCron()

	//c.Start()

	//defer c.Stop()

	//展示界面获取可用的无人机相关设备
	r := Routers.InitRouter()

	r.Static("/img", "./img")

	r.Run()
}
