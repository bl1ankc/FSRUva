package Model

import (
	"fmt"
	"log"
	"time"
)

// RecordBorrow 增加一条记录
func RecordBorrow(Uid string, Stuid string, Borrower string, Plan_time time.Time, Usage string) {

	DB := db.Create(&Record{Uid: Uid, StudentID: Stuid, Borrower: Borrower, Plan_time: Plan_time, Usage: Usage, Get_time: time.Unix(0, 0), Back_time: time.Unix(0, 0), GetReviewTime: time.Unix(0, 0), BackReviewTime: time.Unix(0, 0)})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
		return
	}
	UpdateRecordState(Uid, "Get under review")
	return
}

// UpdateRecordState 更新记录状态
func UpdateRecordState(Uid string, State string) {

	uav := GetUavByUid(Uid)
	//更新状态
	DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, Get_time: uav.Get_time}).Select("state").Updates(&Record{State: State})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
		return
	}

	return
}

// UpdateGetTime 更新借用时间
func UpdateGetTime(Uid string, time time.Time) {

	uav := GetUavByUid(Uid)
	//更新状态
	DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, Get_time: uav.Get_time}).Select("get_time").Updates(&Record{Get_time: time})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
		return
	}

	return
}

// GetRecordsByID 姓名查询记录
func GetRecordsByID(Stuid string) []BackRecord {
	//查找不同的借用时间
	var times []time.Time
	DB := db.Model(&Record{}).Where(&Record{StudentID: Stuid}).Distinct("Get_time").Select("Get_time").Find(&times)
	if DB.Error != nil {
		fmt.Println("GetRecordsByName1 Error")
		log.Fatal(DB.Error.Error())
	}

	var uavpacks []BackRecord

	//查询某时间下借出的设备
	for i, t := range times {

		//查找本次借用的设备
		//var uavs []BackUav
		//DB = db.Debug().Model(&Uav{}).Joins("left join records on records.uid = uavs.uid and records.get_time = ?", t).Find(&uavs)
		//DB = db.Model(&Uav{}).Where(&Uav{Borrower: Name, Get_time: t}).Find(&uavs)

		//查找设备组
		var uavs []BackUav
		DB = db.Model(&Record{}).Where(&Record{Get_time: t, StudentID: Stuid}).Select("state, uid, get_time, back_time, plan_time").Find(&uavs)
		if DB.Error != nil {
			log.Fatal(DB.Error.Error())
		}

		//填充设备信息
		for _, uav := range uavs {
			uavinfo := GetBasicUavsByUid(uav.Uid)
			uav.Name = uavinfo.Name
			uav.Type = uavinfo.Type
			uav.Remark = uavinfo.Remark
			uavs = append(uavs, uav)
		}

		//查询剩余信息
		var uavpack BackRecord
		DB = db.Model(&Record{}).Where(&Record{Get_time: t}).First(&uavpack)
		if DB.Error != nil {
			fmt.Println("GetRecordsByName Error")
			log.Fatal(DB.Error.Error())
		}
		uavpacks = append(uavpacks, uavpack)
		uavpacks[i].Uav = uavs

		//判断本次借用状态
		var states []string
		DB = db.Model(&Record{}).Where(&Record{Get_time: t, StudentID: Stuid}).Select("state").Find(&states)
		if DB.Error != nil {
			fmt.Println("GetRecordsByName Error")
			log.Fatal(DB.Error.Error())
		}
		for _, s := range states {
			uavpacks[i].State = "All returned"
			if s == "using" {
				uavpacks[i].State = "Using"
			} else if s == "damaged" {
				uavpacks[i].State = "Damaged"
				break
			} else if s == "Get under review" || s == "Back under review" {
				uavpacks[i].State = "Reviewing"
				break
			} else if s == "scheduled" {
				uavpacks[i].State = "Scheduled"
				break
			}
		}

	}

	return uavpacks
}

// GetRecordsByUid 序列号查询记录
func GetRecordsByUid(Uid string) []Record {
	var records []Record
	DB := db.Model(&Record{}).Where(&Record{Uid: Uid}).Find(&records)
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
	return records
}

// GetBackUavsByUids 获取对应序列号组的设备组返回信息
func GetBackUavsByUids(Uids []string) []BackUav {
	var uavs []BackUav
	DB := db.Model(&Uav{})

	for _, uid := range Uids {
		var uav BackUav
		DB = db.Model(&Uav{}).Where(&Uav{Uid: uid}).First(&uav)
		if DB.Error != nil {
			log.Fatal(DB.Error.Error())
		}
		uavs = append(uavs, uav)
	}

	return uavs
}

// GetAllRecords 获取所有历史记录
func GetAllRecords() [][]BackRecord {
	//获取用户列表
	Users := GetAllUsers()

	var Records [][]BackRecord
	//查询用户对应记录
	for _, User := range Users {
		Record := GetRecordsByID(User.StudentID)
		Records = append(Records, Record)
	}
	return Records
}

// GetReviewRecord 添加借用审核记录
func GetReviewRecord(Uid string, Checker string, Result string, Comment string, GetTime time.Time) {

	//匹配当前借用设备
	uav := GetUavByUid(Uid)

	DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, Get_time: time.Unix(0, 0)}).Updates(&Record{GetReviewer: Checker, Get_time: GetTime, GetReviewResult: Result, GetReviewComment: Comment})
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
}

// BackReviewRecord 添加归还审核记录
func BackReviewRecord(Uid string, Checker string, Result string, Comment string) {
	//获取时间
	Time := time.Now()

	//匹配当前借用设备
	uav := GetUavByUid(Uid)

	DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, Get_time: uav.Get_time}).Updates(&Record{BackReviewer: Checker, BackReviewTime: Time, BackReviewResult: Result, BackReviewComment: Comment})
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
}

// UpdateBackRecord 添加归还审核时间
func UpdateBackRecord(Uid string) {
	uav := GetUavByUid(Uid)
	DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, Get_time: uav.Get_time}).Updates(&Record{Back_time: time.Now()})
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
}

// UpdateImgInRecord 记录中更新图片
func UpdateImgInRecord(Uid string, col string) {
	uav := GetUavByUid(Uid)
	DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, Get_time: uav.Get_time}).Update(col, uav.Img)
	if DB.Error != nil {
		fmt.Println("UpdateImgInRecord Error!")
		log.Fatal(DB.Error.Error())
	}
}

// GetUsingUavsByStuID 通过学号查找使用中的无人机
func GetUsingUavsByStuID(Stuid string) ([]BackUav, bool) {
	var uids []string
	var uavs []BackUav

	DB := db.Model(&Record{}).Where(&Record{StudentID: Stuid, State: "using"}).Select("uid").Find(&uids)
	if DB.Error != nil {
		fmt.Println("查询错误")
		return uavs, false
	}

	uavs, flag := GetUavsByUids(uids)
	if flag {
		return uavs, true
	} else {
		return uavs, false
	}

}
