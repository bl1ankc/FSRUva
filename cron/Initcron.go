package cron

import (
	"fmt"
	cron2 "github.com/robfig/cron/v3"
	"main/Model"
)

var c *cron2.Cron

var cronid []cron2.EntryID

func InitCron() *cron2.Cron {
	// Cron v3默认支持精确到分钟的cron表达式
	// cron.WithSeconds()表示指定支持精确到秒的表达式
	c = cron2.New(cron2.WithSeconds())

	//每2个小时获取wxcode
	id, err := c.AddFunc("0 0 */2 * * ?", Model.GetWXAccessToken)
	if err != nil {
		fmt.Println("定时任务1出错", err.Error())
	}
	cronid = append(cronid, id)

	//每天7点，18点 提醒用户归还
	id, err = c.AddFunc("0 0 7,18 * * ?", RemindUserReturnUav)
	if err != nil {
		fmt.Println("定时任务2出错", err.Error())
	}
	cronid = append(cronid, id)

	return c
}
