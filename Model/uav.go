package Model

import (
	"fmt"
	"log"
	"time"
)

// GetUavByUid 获取对应序列号的设备信息
func GetUavByUid(Uid string) Uav {
	var uav Uav
	DB := db.Model(&Uav{}).Where(&Uav{Uid: Uid}).First(&uav)

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}

	return uav
}

// GetUavsByUids 获取对应序列号组的设备组信息
func GetUavsByUids(Uids []string) []Uav {
	var uavs []Uav
	DB := db.Model(&Uav{})

	for _, uid := range Uids {
		var uav Uav
		DB = db.Model(&Uav{}).Where(&Uav{Uid: uid}).First(&uav)
		if DB.Error != nil {
			log.Fatal(DB.Error.Error())
		}
		uavs = append(uavs, uav)
	}

	return uavs
}

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

// UpdateBorrowTime 更新借出时间
func UpdateBorrowTime(UavUid string) {
	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).Updates(Uav{Get_time: time.Now().Local()})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
		return
	}

	return
}

// UpdatePlanTime 更新预计归还时间
func UpdatePlanTime(UavUid string, UavPlanTime time.Time) {
	DB := db.Model(&Uav{}).Where(&Uav{Uid: UavUid}).Updates(Uav{Plan_time: UavPlanTime})

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

// UpdateUavRemark 更新设备备注信息
func UpdateUavRemark(Uid string, Remark string) {

	DB := db.Model(&Uav{}).Where(&Uav{Uid: Uid}).Updates(&Uav{Remark: Remark})
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}

}

// GetUavByAll 多条件查找设备信息
func GetUavByAll(uav SearchUav) []BackUav {

	var uavs []BackUav

	DB := db.Model(&Uav{}).Where(&Uav{Uid: uav.Uid, State: uav.State, Name: uav.Name, Type: uav.Type, Borrower: uav.Borrower}).Find(&uavs)
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
	return uavs
}

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
			log.Fatal(DB.Error.Error())
		}
	}
}

// GetUavStateByUid 通过Uid获取设备状态
func (u *Uav) GetUavStateByUid() string {
	var uav Uav
	DB := db.Model(&Uav{}).Where("uid = ?", u.Uid).First(&uav)

	if DB.Error != nil {
		fmt.Println("GetUvaByUid Error")
		log.Fatal(DB.Error.Error())
	}

	return uav.Uid
}
