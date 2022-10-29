package Service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"main/Model"
	"strconv"
	"time"
)

/*
	获取及创建
*/

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

// GetUavByNames 获取对应型号及状态的设备信息
func GetUavByNames(UavName string, UavType string) []Model.Uav {
	var uav []Model.Uav
	DB := db.Model(&Model.Uav{}).Where(Model.Uav{Name: UavName, Type: UavType}).Find(&uav)

	if DB.Error != nil {
		fmt.Println("获取对应型号及状态的设备信息失败：", DB.Error.Error())
	}

	return uav
}

// InsertUva 创建新的设备
func InsertUva(uav Model.Uav) (bool, string) {
	//查询设备是否存在
	var cnt int64
	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: uav.Uid}).Count(&cnt)
	if DB.Error != nil {
		fmt.Println("创建新的设备失败1：", DB.Error.Error())
		return false, "发生未知错误1"
	}
	if cnt >= 1 {
		return false, "设备已存在"
	}

	//创建新记录
	DB = db.Create(&Model.Uav{Name: uav.Name, Type: uav.Type, Uid: uav.Uid, Location: uav.Location})

	if DB.Error != nil {
		fmt.Println("创建新的设备失败2：", DB.Error.Error())
		return false, "发生未知错误2"
	}

	return true, "创建成功"
}

// GetUavByAll 多条件查找设备信息
func GetUavByAll(uav Model.Uav) []Model.Uav {

	var uavs []Model.Uav

	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: uav.Uid, State: uav.State, Name: uav.Name, Type: uav.Type}).Find(&uavs)
	if DB.Error != nil {
		fmt.Println("多条件查找设备信息失败：", DB.Error.Error())
	}
	return uavs
}

// GetUavStateByUid 通过Uid获取设备状态
func GetUavStateByUid(u Model.Uav) string {
	var uav Model.Uav
	DB := db.Model(&Model.Uav{}).Where("uid = ?", u.Uid).First(&uav)

	if DB.Error != nil {
		fmt.Println("通过Uid获取设备状态失败：", DB.Error.Error())
	}
	return uav.State
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

// GetUavsByStatesWithPage 分页获取对应序列号组的设备组信息
func GetUavsByStatesWithPage(UavState string, UavType string, Page string, Max int) []Model.Uav {

	//转换数据格式
	pageint, err := strconv.Atoi(Page)
	if err != nil {
		fmt.Println("分页获取对应序列号组的设备组信息 数据转换失败", err.Error())
		pageint = 0
	}

	//查找数据
	var uavs []Model.Uav

	DB := db.Model(&Model.Uav{}).Where(Model.Uav{State: UavState, Type: UavType}).Offset(pageint * Max).Limit(Max).Find(&uavs)

	if DB.Error != nil {
		fmt.Println("分页获取对应序列号组的设备组信息：", DB.Error.Error())
		return []Model.Uav{}
	}

	return uavs
}

// GetUavType 获取设备类型列表
func GetUavType() (bool, []Model.UavType) {
	var types []Model.UavType
	DB := db.Model(&Model.UavType{}).Find(&types)

	if DB.Error != nil {
		fmt.Println("获取设备类型列表：", DB.Error.Error())
		return false, []Model.UavType{}
	}

	return true, types
}

// GetType 获取单独设备类型
func GetType(id int) (Model.UavType, error) {
	var uavType Model.UavType
	if err := db.Model(&Model.UavType{}).Where("id = ?", id).First(&uavType).Error; err != nil {
		return Model.UavType{}, err
	}
	return uavType, nil
}

// AddUavType 增加设备类型
func AddUavType(TypeName string, Remark string) bool {

	DB := db.Model(&Model.UavType{}).Create(&Model.UavType{TypeName: TypeName, Remark: Remark})

	if DB.Error != nil {
		fmt.Println("增加设备类型失败：", DB.Error.Error())
		return false
	}

	return true
}

// RemoveUavType 删除设备类型
func RemoveUavType(TypeName string) bool {

	DB := db.Where("type_name=?", TypeName).Delete(&Model.UavType{})

	if DB.Error != nil {
		fmt.Println("删除设备类型失败：", DB.Error.Error())
		return false
	}

	return true
}

// UpdateUavType 更新设备类型
func UpdateUavType(uavType Model.UavType) error {
	err := db.Model(&uavType).Updates(uavType).Error
	return err
}

/*
	删除
*/

// RemoveDevice 删除对应设备
func RemoveDevice(Device Model.Uav) error {
	err := db.Model(&Device).Delete(&Device).Error
	return err
}

/*
	更新信息
*/

// UpdateDevice 更新
func UpdateDevice(uav Model.Uav) error {

	if err := db.Model(&uav).Where("uid = ?", uav.Uid).Updates(uav).Error; err != nil {
		return err
	}
	if uav.Expensive == false {
		db.Model(&uav).Where("uid = ?", uav.Uid).Updates(map[string]interface{}{
			"expensive": false,
		})
	}
	return nil
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

// UpdateUavUsage 更新设备用途
func UpdateUavUsage(Uid string, Usage string) error {
	DB := db.Model(&Model.Uav{}).Where(Model.Uav{Uid: Uid}).Updates(&Model.Uav{Usage: Usage})

	if DB.Error != nil {
		fmt.Println("更新设备用途失败：", DB.Error.Error())
		return DB.Error
	}

	return nil
}

// UpdateImg 更新图片img
func UpdateImg(uid string, img string) {

	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: uid}).Update("img", img)

	if DB.Error != nil {
		fmt.Println("更新图片img失败", DB.Error.Error())
		return
	}

	return
}

// UpdateTypeImg 更新类型图片
func UpdateTypeImg(id int, img string) error {
	uavType, err := GetType(id)
	if err != nil {
		return err
	}
	if err := db.Model(&uavType).Update("img", img).Error; err != nil {
		return err
	}
	return nil
}

/*
	修改信息
*/

// UpdateDevices 强制修改设备数据
func UpdateDevices(uav Model.Uav) {
	UpdateDataInUav(uav.Uid, "type", uav.Type)
	UpdateDataInUav(uav.Uid, "name", uav.Name)
	UpdateDataInUav(uav.Uid, "location", uav.Location)
	UpdateDataInUav(uav.Uid, "remark", uav.Remark)
	//UpdateDataInUav(uav.Uid, "borrower", uav.Borrower)
	//UpdateDataInUav(uav.Uid, "phone", uav.Phone)
	//UpdateDataInUav(uav.Uid, "state", uav.State)
}

// UpdateDataInUav 修改设备单个字符串数据
func UpdateDataInUav(Uid string, HeadName string, Data string) {
	if Data != "" {
		DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: Uid}).Update(HeadName, Data)
		if DB.Error != nil {
			fmt.Println("修改设备单个字符串数据失败：", DB.Error.Error())
		}
	}
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

/*
	记录相关
*/

// UpdateRecordIdinUav 在设备中更新记录ID
func UpdateRecordIdinUav(Uid string, id uint) bool {
	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: Uid}).Updates(&Model.Uav{RecordID: id})
	if DB.Error != nil {
		fmt.Println("在设备中更新记录ID失败：", DB.Error.Error())
		return false
	}
	return true
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
