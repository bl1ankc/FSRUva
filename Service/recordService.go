package Service

import (
	"errors"
	"fmt"
	"main/Model"
	"strconv"
	"time"
)

// RecordBorrow 增加一条记录
func RecordBorrow(uav Model.Uav) error {
	var id uint

	//uav := GetUavByUid(Uid)

	DB := db.Create(&Model.Record{Name: uav.Uid, State: "Get under review", Uid: uav.Uid, StudentID: uav.StudentID, Borrower: uav.Borrower, PlanTime: uav.PlanTime, Usage: uav.Usage, GetTime: time.Now(), BackTime: time.Unix(0, 0), GetReviewTime: time.Unix(0, 0), BackReviewTime: time.Unix(0, 0)}).Select("id").Find(&id)
	DB = db.Model(&uav).Updates(&Model.Uav{RecordID: id})
	if DB.Error != nil {
		fmt.Println("增加一条记录失败：", DB.Error.Error())
		return DB.Error
	}

	return nil
}

// GetRecordByUid 序列号单一查找
func GetRecordByUid(uid string) (Model.Record, error) {
	var data Model.Record
	if err := db.Model(&Model.Record{}).Where("uid = ?", uid).First(&data).Error; err != nil {
		return Model.Record{}, err
	}
	return data, nil
}

// UpdateRecord 更新记录
func UpdateRecord(record Model.Record) error {
	if DB := db.Model(&record).Updates(&record); DB.Error != nil {
		return DB.Error
	}
	return nil
}

// GetRecordById 通过记录ID获取记录实例
func GetRecordById(recordid uint) (bool, Model.Record) {
	var record Model.Record

	DB := db.Model(&Model.Record{}).Where("id", recordid).First(&record)
	if DB.Error != nil {
		fmt.Println("通过记录ID查找记录信息失败", DB.Error.Error())
		return false, record
	}

	return true, record
}

// UpdateRecordState 更新记录状态
func UpdateRecordState(Uid string, State string) error {

	//uav := GetUavByUid(Uid)
	//更新状态
	//DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, GetTime: uav.GetTime}).Updates(&Record{State: State})
	id, flag := GetRecordIdinUav(Uid)
	if !flag {
		return errors.New("NotFind")
	}

	DB := db.Model(&Model.Record{}).Where("id = ?", id).Updates(&Model.Record{State: State})
	if DB.Error != nil {
		fmt.Println("更新记录状态失败：", DB.Error.Error())
		return DB.Error
	}

	return nil
}

// UpdateGetTimeinRecords 记录中更新借用时间
func UpdateGetTimeinRecords(Uid string) bool {

	//uav := GetUavByUid(Uid)
	//更新状态
	//DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, GetTime: uav.GetTime}).Select("GetTime").Updates(&Record{GetTime: time})

	id, flag := GetRecordIdinUav(Uid)
	if !flag {
		return false
	}

	DB := db.Model(&Model.Record{}).Where("id", id).Updates(&Model.Record{GetTime: time.Now()})
	if DB.Error != nil {
		fmt.Println("记录中更新借用时间：", DB.Error.Error())
		return false
	}

	return true
}

// GetRecordsByID 学号查询记录
func GetRecordsByID(Stuid string) []Model.BackRecord {
	//查找不同的借用时间
	var times []time.Time
	DB := db.Model(&Model.Record{}).Where(&Model.Record{StudentID: Stuid}).Distinct("Get_Time").Select("Get_Time").Order("Get_Time Desc").Find(&times)
	if DB.Error != nil {
		fmt.Println("学号查询记录失败1：", DB.Error.Error())
	}

	var uavpacks []Model.BackRecord

	//查询某时间下借出的设备
	for i, t := range times {

		//查找本次借用的设备
		//var uavs []Uav
		//DB = db.Debug().Model(&Uav{}).Joins("left join records on records.uid = uavs.uid and records.GetTime = ?", t).Find(&uavs)
		//DB = db.Model(&Uav{}).Where(&Uav{Borrower: Name, GetTime: t}).Find(&uavs)

		//查找设备组
		var uavs []Model.Uav
		DB = db.Model(&Model.Record{}).Where(&Model.Record{GetTime: t, StudentID: Stuid}).Select("state, uid, Get_Time, BackTime, PlanTime").Find(&uavs)
		if DB.Error != nil {
			fmt.Println("学号查询记录失败2：", DB.Error.Error())
		}

		//填充设备信息
		for _, uav := range uavs {
			exist, uavinfo := GetUavByUid(uav.Uid)
			if exist == true {
				uav.Name = uavinfo.Name
				uav.Type = uavinfo.Type
				uav.Location = uavinfo.Location
				uav.Remark = uavinfo.Remark
				uavs = append(uavs, uav)
			}
		}

		//查询剩余信息
		var uavpack Model.BackRecord
		DB = db.Model(&Model.Record{}).Where(&Model.Record{GetTime: t}).First(&uavpack)
		if DB.Error != nil {
			fmt.Println("学号查询记录失败3：", DB.Error.Error())
		}
		uavpacks = append(uavpacks, uavpack)
		uavpacks[i].Uav = uavs

		//判断本次借用状态
		var states []string
		DB = db.Model(&Model.Record{}).Where(&Model.Record{GetTime: t, StudentID: Stuid}).Select("state").Find(&states)
		if DB.Error != nil {
			fmt.Println("学号查询记录失败4：", DB.Error.Error())
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
func GetRecordsByUid(Uid string, Page string, Max int) []Model.Record {

	//转换数据格式
	pageint, err := strconv.Atoi(Page)
	if err != nil {
		fmt.Println("序列号查询记录 数据转换失败", err.Error())
		pageint = 0
	}

	var records []Model.Record
	DB := db.Model(&Model.Record{}).Where(&Model.Record{Uid: Uid}).Order("Get_Time Desc").Offset(pageint * Max).Limit(Max).Find(&records)
	if DB.Error != nil {
		fmt.Println("序列号查询记录失败：", DB.Error.Error())
	}
	return records
}

// GetAllRecords 获取所有历史记录
func GetAllRecords() [][]Model.BackRecord {
	//获取用户列表
	Users := GetAllUsers()

	var Records [][]Model.BackRecord
	//查询用户对应记录
	for _, User := range Users {
		Record := GetRecordsByID(User.StudentID)
		Records = append(Records, Record)
	}
	return Records
}

// GetReviewRecord 添加借用审核记录
func GetReviewRecord(Uid string, Checker string, Result string, Comment string, GetTime time.Time) error {

	//匹配当前借用设备
	//uav := GetUavByUid(Uid)

	//DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, GetTime: time.Unix(0, 0)}).Updates(&Record{GetReviewer: Checker, GetTime: GetTime, GetReviewResult: Result, GetReviewComment: Comment})
	//if DB.Error != nil {
	//	fmt.Println("添加借用审核记录失败：", DB.Error.Error())
	//}

	id, flag := GetRecordIdinUav(Uid)
	if !flag {
		return errors.New("NotFind")
	}
	fmt.Println("//////////+" + Comment)
	DB := db.Model(&Model.Record{}).Where("id", id).Updates(&Model.Record{GetReviewer: Checker, GetTime: GetTime, GetReviewResult: Result, GetReviewComment: Comment, GetReviewTime: time.Now()})
	if DB.Error != nil {
		fmt.Println("添加借用审核记录失败：", DB.Error.Error())
		return DB.Error
	}

	return nil
}

// BackReviewRecord 添加归还审核记录
func BackReviewRecord(Uid string, Checker string, Result string, Comment string) bool {
	//获取时间
	//Time := time.Now()

	//匹配当前借用设备
	//uav := GetUavByUid(Uid)

	//DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, GetTime: uav.GetTime}).Updates(&Record{BackReviewer: Checker, BackReviewTime: Time, BackReviewResult: Result, BackReviewComment: Comment})
	//if DB.Error != nil {
	//	fmt.Println("添加归还审核记录失败：", DB.Error.Error())
	//}
	id, flag := GetRecordIdinUav(Uid)
	if !flag {
		return false
	}

	DB := db.Model(&Model.Record{}).Where("id", id).Updates(&Model.Record{BackReviewer: Checker, BackReviewTime: time.Now(), BackReviewResult: Result, BackReviewComment: Comment})
	if DB.Error != nil {
		fmt.Println("添加归还审核记录失败：", DB.Error.Error())
		return false
	}

	return true
}

// UpdateBackRecord 添加归还时间
func UpdateBackRecord(Uid string) bool {
	//uav := GetUavByUid(Uid)
	//DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, GetTime: uav.GetTime}).Updates(&Record{BackTime: time.Now()})
	//if DB.Error != nil {
	//	fmt.Println("添加归还时间失败：", DB.Error.Error())
	//}
	id, flag := GetRecordIdinUav(Uid)
	if !flag {
		return false
	}

	DB := db.Model(&Model.Record{}).Where("id", id).Updates(&Model.Record{BackTime: time.Now()})
	if DB.Error != nil {
		fmt.Println("添加归还时间失败：", DB.Error.Error())
		return false
	}

	return true
}

// UpdateImgInRecord 记录中更新图片
func UpdateImgInRecord(Uid string, col string) bool {
	exist, uav := GetUavByUid(Uid)
	if exist == false {
		fmt.Println("未找到对应设备")
		return exist
	}
	DB := db.Model(&Model.Record{}).Where("id", uav.RecordID).Update(col, uav.Img)
	if DB.Error != nil {
		fmt.Println("更新照片失败：", DB.Error.Error())
		return false
	}
	return true
}

// GetUsingUavsByStuID 通过学号查找使用中的无人机
func GetUsingUavsByStuID(Stuid string, Page string, Max int) ([]Model.UsingUav, bool) {

	//转换数据格式
	pageint, err := strconv.Atoi(Page)
	if err != nil {
		fmt.Println("通过学号查找使用中的无人机 数据转换失败", err.Error())
		pageint = 0
	}

	var uavs []Model.UsingUav
	var uav Model.UsingUav

	var tempuav []Model.Record
	DB := db.Model(&Model.Record{}).Where(&Model.Record{StudentID: Stuid, State: "using"}).Or(&Model.Record{StudentID: Stuid, State: "Get under review"}).Or(&Model.Record{StudentID: Stuid, State: "scheduled"}).Order("Get_Time Desc").Offset(pageint * Max).Limit(Max).Find(&tempuav)
	if DB.Error != nil {
		fmt.Println("通过学号查找使用中的无人机失败：", DB.Error.Error())
		return uavs, false
	}
	flag := true
	for _, tempuav := range tempuav {
		uav.Uid = tempuav.Uid
		uav.Name, flag = GetUavNameByUid(uav.Uid)
		if flag == false {
			return uavs, false
		}
		uav.State = tempuav.State
		year, month, day := tempuav.GetTime.Date()
		uav.GetTime = strconv.Itoa(year) + "." + strconv.Itoa(int(month)) + "." + strconv.Itoa(day)
		year, month, day = tempuav.PlanTime.Date()
		uav.PlanTime = strconv.Itoa(year) + "." + strconv.Itoa(int(month)) + "." + strconv.Itoa(day)
		uav.LastDays = int(tempuav.PlanTime.Sub(time.Now()).Hours()) / 24
		uav.GetComment = tempuav.GetReviewComment
		uav.BackComment = tempuav.BackReviewComment
		uav.TmpImg = tempuav.TmpImg
		uavs = append(uavs, uav)
	}
	return uavs, true
}

// GetHistoryUavsByStuID 通过学号查找历史借用的无人机
func GetHistoryUavsByStuID(Stuid string, Page string, Max int) ([]Model.UsingUav, bool) {

	//转换数据格式
	pageint, err := strconv.Atoi(Page)
	if err != nil {
		fmt.Println("通过学号查找历史借用的无人机 数据转换失败", err.Error())
		pageint = 0
	}

	var uavs []Model.UsingUav
	var uav Model.UsingUav

	var tempuav []Model.Record
	DB := db.Model(&Model.Record{}).Where(&Model.Record{StudentID: Stuid, State: "returned"}).Or(&Model.Record{StudentID: Stuid, State: "refuse"}).Or(&Model.Record{StudentID: Stuid, State: "cancelled"}).Or(&Model.Record{StudentID: Stuid, State: "Back under review"}).Order("get_time desc").Offset(pageint * Max).Limit(Max).Find(&tempuav)
	if DB.Error != nil {
		fmt.Println("通过学号查找历史借用的无人机失败", DB.Error.Error())
		return uavs, false
	}
	flag := true
	for _, tempuav := range tempuav {
		uav.Uid = tempuav.Uid
		uav.Name, flag = GetUavNameByUid(uav.Uid)
		if flag == false {
			return uavs, false
		}
		year, month, day := tempuav.GetTime.Date()
		uav.GetTime = strconv.Itoa(year) + "." + strconv.Itoa(int(month)) + "." + strconv.Itoa(day)
		year, month, day = tempuav.PlanTime.Date()
		uav.PlanTime = strconv.Itoa(year) + "." + strconv.Itoa(int(month)) + "." + strconv.Itoa(day)
		uav.LastDays = 0
		uav.State = tempuav.State
		uav.GetComment = tempuav.GetReviewComment
		uav.BackComment = tempuav.BackReviewComment
		uav.TmpImg = tempuav.TmpImg
		uavs = append(uavs, uav)
	}
	return uavs, true
}
