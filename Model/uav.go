package Model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/*
	获取及创建
*/

// GetUavByUid 获取对应序列号的设备信息
func GetUavByUid(Uid string) Uav {
	var uav Uav
	DB := db.Model(&Uav{}).Where(&Uav{Uid: Uid}).First(&uav)

	if errors.Is(DB.Error, gorm.ErrRecordNotFound) {
		fmt.Println("获取信息失败：", DB.Error.Error())
		return Uav{}
	}

	return uav
}

// GetUavsByUids 获取对应序列号组的设备组信息
func GetUavsByUids(Uids []string) ([]BackUav, bool) {
	var uavs []BackUav

	for _, uid := range Uids {
		var uav BackUav
		DB := db.Model(&Uav{}).Where(&Uav{Uid: uid}).First(&uav)
		if DB.Error != nil {
			fmt.Println("获取对应序列号组的设备组信息失败：", DB.Error.Error())
			return uavs, false
		}
		uavs = append(uavs, uav)
	}

	return uavs, true
}

// GetUavByStates 获取对应状态及类型的设备信息
func GetUavByStates(UavState string, UavType string) []Uav {
	var uav []Uav
	DB := db.Model(&Uav{}).Where(Uav{State: UavState, Type: UavType}).Find(&uav)

	if DB.Error != nil {
		fmt.Println("获取对应状态及类型的设备信息失败：", DB.Error.Error())
	}

	return uav
}

// GetUavByNames 获取对应型号及状态的设备信息
func GetUavByNames(UavName string, UavType string) []Uav {
	var uav []Uav
	DB := db.Model(&Uav{}).Where(Uav{Name: UavName, Type: UavType}).Find(&uav)

	if DB.Error != nil {
		fmt.Println("获取对应型号及状态的设备信息失败：", DB.Error.Error())
	}

	return uav
}

// InsertUva 创建新的设备
func InsertUva(UavName string, UavType string, UavUid string) (bool, string) {
	//查询设备是否存在
	var cnt int64
	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).Count(&cnt)
	if DB.Error != nil {
		fmt.Println("创建新的设备失败1：", DB.Error.Error())
		return false, "发生未知错误1"
	}
	if cnt >= 1 {
		return false, "设备已存在"
	}

	//创建新记录
	DB = db.Create(&Uav{Name: UavName, Type: UavType, Uid: UavUid, Get_time: time.Unix(0, 0), Plan_time: time.Unix(0, 0), Back_time: time.Unix(0, 0)})

	if DB.Error != nil {
		fmt.Println("创建新的设备失败2：", DB.Error.Error())
		return false, "发生未知错误2"
	}

	return true, "创建成功"
}

// GetUavByAll 多条件查找设备信息
func GetUavByAll(uav SearchUav) []Uav {

	var uavs []Uav

	DB := db.Model(&Uav{}).Where(&Uav{Uid: uav.Uid, State: uav.State, Name: uav.Name, Type: uav.Type}).Find(&uavs)
	if DB.Error != nil {
		fmt.Println("多条件查找设备信息失败：", DB.Error.Error())
	}
	return uavs
}

// GetUavStateByUid 通过Uid获取设备状态
func (u *BorrowUav) GetUavStateByUid() string {
	var uav Uav
	DB := db.Model(&Uav{}).Where("uid = ?", u.Uid).First(&uav)

	if DB.Error != nil {
		fmt.Println("通过Uid获取设备状态失败：", DB.Error.Error())
	}
	return uav.State
}

// GetBasicUavsByUid 通过Uid获取设备基础信息
func GetBasicUavsByUid(Uid string) BasicUav {
	var uav BasicUav
	DB := db.Model(&Uav{}).Where(&Uav{Uid: Uid}).First(&uav)

	if DB.Error != nil {
		fmt.Println("通过Uid获取设备基础信息失败：", DB.Error.Error())
	}

	return uav
}

/*
	更新信息
*/

// UpdateState 更新状态
func UpdateState(UavUid string, UavState string) {
	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).Select("state").Updates(Uav{State: UavState})

	if DB.Error != nil {
		fmt.Println("更新状态失败：", DB.Error.Error())
		return
	}

	return
}

// UpdateBorrower 更新借用人信息
func UpdateBorrower(UavUid string, UavBorrower string, UavPhone string, UavStuID string) {
	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).Updates(Uav{Borrower: UavBorrower, Phone: UavPhone, StudentID: UavStuID})

	if DB.Error != nil {
		fmt.Println("更新借用人信息失败：", DB.Error.Error())
		return
	}

	return
}

// UpdateBorrowTime 更新借出时间
func UpdateBorrowTime(UavUid string, BorrowTime time.Time) {

	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).Updates(Uav{Get_time: BorrowTime})

	if DB.Error != nil {
		fmt.Println("更新借出时间失败：", DB.Error.Error())
		return
	}

	return
}

// UpdatePlanTime 更新预计归还时间
func UpdatePlanTime(UavUid string, UavPlanTime time.Time) {
	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).Updates(Uav{Plan_time: UavPlanTime})

	if DB.Error != nil {
		fmt.Println("更新预计归还时间失败：", DB.Error.Error())
		return
	}

	return
}

// UpdateBackTime 更新归还时间
func UpdateBackTime(UavUid string) {
	var uav Uav

	//获取对应设备结构体信息
	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).First(&uav)

	db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).First(&uav).Updates(Uav{Back_time: time.Now()})

	if DB.Error != nil {
		fmt.Println("更新归还时间失败：", DB.Error.Error())
		return
	}

	return
}

// UpdateUavRemark 更新设备备注信息
func UpdateUavRemark(Uid string, Remark string) {

	DB := db.Model(&Uav{}).Where(&Uav{Uid: Uid}).Updates(&Uav{Remark: Remark})
	if DB.Error != nil {
		fmt.Println("更新设备备注信息失败：", DB.Error.Error())
	}

}

// UpdateUavUsage 更新设备用途
func UpdateUavUsage(Uid string, Usage string) {
	DB := db.Model(&Uav{}).Where(Uav{Uid: Uid}).Updates(&Uav{Usage: Usage})

	if DB.Error != nil {
		fmt.Println("更新设备用途失败：", DB.Error.Error())
	}
}

// UpdateImg 更新图片img
func UpdateImg(uid string, img string) {

	DB := db.Model(&Uav{}).Where(&Uav{Uid: uid}).Update("img", img)

	if DB.Error != nil {
		fmt.Println("更新图片img失败", DB.Error.Error())
		return
	}

	return
}

/*
	修改信息
*/

// UpdateDevices 强制修改设备数据
func UpdateDevices(uav ChangeUav) {
	UpdateDataInUav(uav.Uid, "type", uav.Type)
	UpdateDataInUav(uav.Uid, "name", uav.Name)
	UpdateDataInUav(uav.Uid, "borrower", uav.Borrower)
	UpdateDataInUav(uav.Uid, "phone", uav.Phone)
	UpdateDataInUav(uav.Uid, "state", uav.State)
}

// UpdateDataInUav 修改设备单个字符串数据
func UpdateDataInUav(Uid string, HeadName string, Data string) {
	if Data != "" {
		DB := db.Model(&Uav{}).Where(&Uav{Uid: Uid}).Update(HeadName, Data)
		if DB.Error != nil {
			fmt.Println("修改设备单个字符串数据失败：", DB.Error.Error())
		}
	}
}

// GetUavNameByUid 通过序列号查找设备名
func GetUavNameByUid(Uid string) (string, bool) {
	var name string
	DB := db.Model(&Uav{}).Where(&Uav{Uid: Uid}).Select("name").First(&name)
	if DB.Error != nil {
		fmt.Println("通过序列号查找设备名失败：", DB.Error.Error())
		return "", false
	}
	return name, true
}

// UpdateRecordIdinUav 在设备中更新记录ID
func UpdateRecordIdinUav(Uid string, id uint) bool {
	DB := db.Model(&Uav{}).Where(&Uav{Uid: Uid}).Updates(&Uav{RecordID: id})
	if DB.Error != nil {
		fmt.Println("在设备中更新记录ID失败：", DB.Error.Error())
		return false
	}
	return true
}

// GetRecordIdinUav 在设备中获取记录ID
func GetRecordIdinUav(Uid string) (uint, bool) {
	var id uint
	DB := db.Model(&Uav{}).Where(&Uav{Uid: Uid}).Select("record_id").Find(&id)
	if DB.Error != nil {
		fmt.Println("在设备中更新记录ID失败：", DB.Error.Error())
		return 0, false
	}
	return id, true
}

// SearchStuInOneDay 获取即将归还的设备
func SearchStuInOneDay() ([]Uav, bool) {
	var uavs []Uav
	DB := db.Model(&Uav{}).Where(&Uav{State: "using"}).Where("plan_time > ?", time.Now().AddDate(0, 0, -2)).Find(&uavs)
	if DB.Error != nil {
		fmt.Println("获取即将要归还的无人机失败：", DB.Error.Error())
		return uavs, false
	}
	return uavs, true
}
