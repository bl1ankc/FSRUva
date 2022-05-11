package cron

import (
	"main/Model"
)

type Value struct {
	Value string `json:"value"`
}

// RemindUserReturnUav 提醒用户归还无人机(未完成)
func RemindUserReturnUav() {
	//查找即将要归还的无人机
	uavs, _ := Model.SearchStuInOneDay()

	//定义返回消息
	Message := "您有要归还的设备，请及时归还"

	type T struct {
		StartTime Value `json:"date01"`
		EndTime   Value `json:"date02"`
		Comment   Value `json:"thing01"`
	}

	for _, uav := range uavs {
		data := T{
			StartTime: Value{Value: uav.Get_time.Format("2006-01-02 15:04")},
			EndTime:   Value{Value: uav.Plan_time.Format("2006-01-02 15:04")},
			Comment:   Value{Value: Message},
		}
		user := Model.GetUserInfoByStudentId(uav.StudentID)
		Model.SendMessage(user.Openid, "RemindUserReturnUav", "/pages/returnby_rfid/returnby_rfid", data)
	}
}

// RemindScheduleOK 预约成功
func RemindScheduleOK(Uid string) {
	uav := Model.GetUavByUid(Uid)

	//定义返回消息
	//Message := "您已预约成功！"

	type T struct {
		Name      Value `json:"thing01"`
		StartTime Value `json:"date01"`
		EndTime   Value `json:"date02"`
		Usage     Value `json:"thing02"`
		UavName   Value `json:"thing03"`
	}

	data := T{
		Name:      Value{Value: uav.Borrower},
		StartTime: Value{Value: uav.Get_time.Format("2006-01-02 15:04")},
		EndTime:   Value{Value: uav.Plan_time.Format("2006-01-02 15:04")},
		Usage:     Value{Value: uav.Usage},
		UavName:   Value{Value: uav.Name},
	}

	user := Model.GetUserInfoByStudentId(uav.StudentID)

	Model.SendMessage(user.Openid, "RemindScheduleOK", "", data)
}

// RemindCheckOK 审核成功
func RemindCheckOK(uid string, op string) {
	uav := Model.GetUavByUid(uid)

	_, record := Model.GetRecordById(uav.RecordID)

	//定义返回消息
	//Message := "您已预约成功！"
	var page string
	type T struct {
		Name      Value `json:"thing01"`
		CheckTime Value `json:"date01"`
		Result    Value `json:"thing02"`
		Checker   Value `json:"thing03"`
	}

	data := T{
		Name: Value{Value: uav.Borrower},
	}

	if op == "get" {
		data.CheckTime = Value{Value: record.GetReviewTime.Format("2006-01-02 15:04")}
		data.Result = Value{Value: record.GetReviewResult}
		data.Checker = Value{Value: record.GetReviewer}
		page = "/pages/takeuav/takeuav?uid=" + uav.Uid
	} else if op == "back" {
		data.CheckTime = Value{Value: record.BackReviewTime.Format("2006-01-02 15:04")}
		data.Result = Value{Value: record.BackReviewResult}
		data.Checker = Value{Value: record.BackReviewer}
		page = ""
	}

	user := Model.GetUserInfoByStudentId(record.StudentID)

	Model.SendMessage(user.Openid, "RemindCheckOK", page, data)
}

// RemindAdminCheck 审核通知
func RemindAdminCheck(uid string, op string) {

	uav := Model.GetUavByUid(uid)
	//定义返回消息
	//Message := "您已预约成功！"

	type T struct {
		Name      Value `json:"thing01"`
		StartTime Value `json:"date01"`
		EndTime   Value `json:"date02"`
		Comment   Value `json:"thing02"`
	}

	data := T{
		Name:      Value{Value: uav.Borrower},
		StartTime: Value{Value: uav.Get_time.Format("2006-01-02 15:04")},
		EndTime:   Value{Value: uav.Plan_time.Format("2006-01-02 15:04")},
	}
	if op == "get" {
		data.Comment = Value{Value: "借用审核：" + uav.Name}
	} else if op == "back" {
		data.Comment = Value{Value: "归还审核：" + uav.Name}
	}

	user := Model.GetUserInfoByStudentId(uav.StudentID)

	Model.SendMessage(user.Openid, "RemindAdminCheck", "pages/showunderreview/showunderreview", data)
}
