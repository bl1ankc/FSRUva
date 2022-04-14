package Model

import (
	"fmt"
	"log"
	"time"
)

// RecordBorrow 增加一条记录
func RecordBorrow(Uid string, Name string, Get_time time.Time, Plan_time time.Time) {

	DB := db.Create(&Record{Uid: Uid, Borrower: Name, Get_time: Get_time, Plan_time: Plan_time})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
		return
	}
	UpdateRecordState(Uid, Name, Get_time, "Get under review")
	return
}

// UpdateRecordState 更新记录状态
func UpdateRecordState(Uid string, Name string, Get_time time.Time, State string) {
	DB := db.Model(&Record{}).Where(&Record{Uid: Uid, Borrower: Name, Get_time: Get_time}).Select("state").Updates(&Record{State: State})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
		return
	}

	return
}

// GetRecordsByName 姓名查询记录
func GetRecordsByName(Name string) []BackRecord {
	//查找不同的借用时间
	var times []time.Time
	DB := db.Model(&Record{}).Where(&Record{Borrower: Name}).Distinct("Get_time").Select("Get_time").Find(&times)
	if DB.Error != nil {
		fmt.Println("GetRecordsByName1 Error")
		log.Fatal(DB.Error.Error())
	}

	var uavpacks []BackRecord
	i := 0
	//查询某时间下借出的设备
	for _, t := range times {
		var uavs []BackUav
		DB = db.Model(&Uav{}).Where(&Uav{Borrower: Name, Get_time: t}).Find(&uavs)
		if DB.Error != nil {
			fmt.Println("GetRecordsByName2 Error")
			log.Fatal(DB.Error.Error())
		}
		var uavpack BackRecord
		DB = db.Model(&Record{}).Where(&Record{Get_time: t}).First(&uavpack)
		if DB.Error != nil {
			fmt.Println("GetRecordsByName Error")
			log.Fatal(DB.Error.Error())
		}
		uavpacks = append(uavpacks, uavpack)
		uavpacks[i].Uav = uavs
		i++
	}

	return uavpacks
}
