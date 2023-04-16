package Service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"main/Model"
	"main/Service/Scopes"
	"net/http"
	"time"
)

// InsertUva 创建
func InsertUva(uav Model.Uav) (bool, string) {
	var cnt int64
	//查询设备是否存在
	if err := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: uav.Uid}).Count(&cnt).Error; err != nil {
		fmt.Println("创建新的设备失败1：", err.Error())
		return false, "发生未知错误1"
	}
	if cnt >= 1 {
		return false, "设备已存在"
	}

	//创建新记录
	if err := db.Select("name", "type", "uid", "location", "expensive").Create(&uav).Error; err != nil {
		fmt.Println("创建新的设备失败2：", err.Error())
		return false, "发生未知错误2"
	}

	return true, "创建成功"
}

// RemoveDevice 删除
func RemoveDevice(Device Model.Uav) error {
	err := db.Model(&Device).Delete(&Device).Error
	return err
}

// UpdateDevice 更新
func UpdateDevice(uav Model.Uav) error {

	if err := db.Model(&Model.Uav{}).Where("uid = ?", uav.Uid).Updates(uav).Error; err != nil {
		return err
	}
	if uav.Expensive == false {
		db.Model(&uav).Where("uid = ?", uav.Uid).Updates(map[string]interface{}{
			"expensive": false,
		})
	}
	return nil
}

// GetUavByUid 获取对应序列号的设备信息
func GetUavByUid(Uid string) (bool, Model.Uav) {
	var uav Model.Uav
	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: Uid}).First(&uav)

	if errors.Is(DB.Error, gorm.ErrRecordNotFound) {
		fmt.Println("获取信息失败：", DB.Error.Error())
		return false, Model.Uav{}
	}

	return true, uav
}

// GetUavsByUids 获取对应序列号组的设备组信息
func GetUavsByUids(Uids []string) []Model.Uav {
	var uavs []Model.Uav

	for _, uid := range Uids {
		var uav Model.Uav
		DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: uid}).First(&uav)
		if DB.Error != nil {
			fmt.Println("获取对应序列号组的设备组信息失败：", DB.Error.Error())
		} else {
			uavs = append(uavs, uav)
		}

	}

	return uavs
}

// GetUavByStates 获取对应状态及类型的设备信息
func GetUavByStates(UavState string, UavType string) []Model.Uav {
	var uav []Model.Uav
	DB := db.Model(&Model.Uav{}).Where(Model.Uav{State: UavState, Type: UavType}).Find(&uav)

	if DB.Error != nil {
		fmt.Println("获取对应状态及类型的设备信息失败：", DB.Error.Error())
	}

	return uav
}

// GetUavByAll 多条件查找设备信息
func GetUavByAll(uav Model.Uav, filed []string) []Model.Uav {

	var uavs []Model.Uav

	DB := db.Model(&Model.Uav{}).Select(filed).Where(&Model.Uav{State: uav.State, Type: uav.Type}).Find(&uavs)
	if DB.Error != nil {
		fmt.Println("多条件查找设备信息失败：", DB.Error.Error())
	}
	return uavs
}

// SearchStuInOneDay 获取即将归还的设备
func SearchStuInOneDay() ([]Model.Uav, bool) {
	var uavs []Model.Uav
	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{State: "using"}).Where("PlanTime > ?", time.Now().AddDate(0, 0, -2)).Find(&uavs)
	if DB.Error != nil {
		fmt.Println("获取即将要归还的无人机失败：", DB.Error.Error())
		return uavs, false
	}
	return uavs, true
}

// GetUavsByStatesWithPage 分页获取对应序列号组的设备组信息 @2023/3/14
func GetUavsByStatesWithPage(UavState string, UavType string, r *http.Request) ([]Model.Uav, int64, error) {
	//查找数据
	var uavs []Model.Uav
	var total int64
	if err := db.Model(&Model.Uav{}).Where(Model.Uav{State: UavState, Type: UavType}).Scopes(Scopes.Paginate(r)).Find(&uavs).Error; err != nil {
		return []Model.Uav{}, 0, err
	}
	if err := db.Model(&Model.Uav{}).Where(Model.Uav{State: UavState, Type: UavType}).Count(&total).Error; err != nil {
		return []Model.Uav{}, 0, err
	}
	return uavs, total, nil
}

// UpdateState 更新状态
func UpdateState(UavUid string, UavState string) error {
	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: UavUid}).Select("state").Updates(Model.Uav{State: UavState})

	if DB.Error != nil {
		fmt.Println("更新状态失败：", DB.Error.Error())
		return DB.Error
	}

	return nil
}

// UpdateBorrowTime 更新借出时间
func UpdateBorrowTime(UavUid string, BorrowTime time.Time) error {
	BorrowTime = BorrowTime.Add(-time.Hour * 8) //时区矫正
	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: UavUid}).Updates(Model.Uav{GetTime: BorrowTime})

	if DB.Error != nil {
		fmt.Println("更新借出时间失败：", DB.Error.Error())
		return DB.Error
	}

	return nil
}

// UpdateBackTime 更新归还时间
func UpdateBackTime(UavUid string) {
	var uav Model.Uav

	//获取对应设备结构体信息
	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: UavUid}).First(&uav)

	db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: UavUid}).First(&uav).Updates(Model.Uav{BackTime: time.Now()})

	if DB.Error != nil {
		fmt.Println("更新归还时间失败：", DB.Error.Error())
		return
	}

	return
}

// UpdateUavRemark 更新设备备注信息
func UpdateUavRemark(Uid string, Remark string) {

	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: Uid}).Updates(&Model.Uav{Remark: Remark})
	if DB.Error != nil {
		fmt.Println("更新设备备注信息失败：", DB.Error.Error())
	}

}

// UpdateUavImg 更新图片img
func UpdateUavImg(uid string, img string) error {

	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: uid}).Update("cur_img", img)

	if DB.Error != nil {
		fmt.Println("更新图片img失败", DB.Error.Error())
		return DB.Error
	}

	return nil
}

// GetUavNameByUid 通过序列号查找设备名
func GetUavNameByUid(Uid string) (string, bool) {
	var name string
	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: Uid}).Select("name").First(&name)
	if DB.Error != nil {
		fmt.Println("通过序列号查找设备名失败：", DB.Error.Error())
		return "", false
	}
	return name, true
}

// GetRecordIdinUav 在设备中获取记录ID
func GetRecordIdinUav(Uid string) (uint, bool) {
	var id uint
	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: Uid}).Select("record_id").Find(&id)
	if DB.Error != nil {
		fmt.Println("在设备中更新记录ID失败：", DB.Error.Error())
		return 0, false
	}
	return id, true
}
