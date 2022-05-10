package cron

import "main/Model"

// RemindUserReturnUav 提醒用户归还无人机(未完成)
func RemindUserReturnUav() {
	//查找即将要归还的无人机
	uavs, _ := Model.SearchStuInOneDay()

	//定义返回消息
	type Data struct {
	}
	data := Data{}

	for _, uav := range uavs {
		user := Model.GetUserInfoByStudentId(uav.StudentID)
		Model.SendMessage(user.Openid, "", "", data)
	}
}

//预约成功
//审核成功
//审核通知
