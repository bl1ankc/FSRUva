package Model

import (
	"fmt"
	"log"
	"time"
)

// GetUavByStates 获取对应状态及类型的设备信息
func GetUavByStates(UavState string, UavType string) []Uav {
	var uav []Uav
	DB := db.Model(&Uav{}).Where(Uav{State: UavState, Type: UavType}).Find(&uav)

	if DB.Error != nil {
		fmt.Println("GetUvasByState Error")
		log.Fatal(DB.Error.Error())
	}

	return uav
}

// GetUavByNames 获取对应型号及状态的设备信息
func GetUavByNames(UavName string, UavType string) []Uav {
	var uav []Uav
	DB := db.Model(&Uav{}).Where(Uav{Name: UavName, Type: UavType}).Find(&uav)

	if DB.Error != nil {
		fmt.Println("GetUvasByState Error")
		log.Fatal(DB.Error.Error())
	}

	return uav
}

// InsertUva 创建新的设备
func InsertUva(UavName string, UavType string) {
	//创建新记录
	DB := db.Create(&Uav{Name: UavName, Type: UavType, Uid: GetUid()})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
		return
	}

	return
}

// UpdateState 更新状态
func UpdateState(UavUid string, UavState string) {
	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).Select("state").Updates(Uav{State: UavState})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
		return
	}

	return
}

// UpdateBorrower 更新借用人信息
func UpdateBorrower(UavUid string, UavBorrower string, UavPhone string) {
	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).Select("borrower", "phone").Updates(Uav{Borrower: UavBorrower, Phone: UavPhone})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
		return
	}

	return
}

// UpdateBorrowTime 更新借用时间
func UpdateBorrowTime(UavUid string, UavPlantime time.Time) {
	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).Updates(Uav{Get_time: time.Now().Local(), Plan_time: UavPlantime})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
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
		log.Fatal(DB.Error.Error())
		return
	}

	return
}
