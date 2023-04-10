package Service

import (
	"main/Model"
	"time"
)

type DeviceInfo struct {
	ID         uint
	Name       string
	State      string
	Department string `gorm:"default(-)"`
	Type       string
	Uid        string
	Location   string
	Remark     string
	Expensive  bool
}

// GetDeviceData 所有设备信息
func GetDeviceData() ([]DeviceInfo, error) {

	var re []DeviceInfo
	if err := db.Model(&Model.Uav{}).Omit("Department").Find(&re).Error; err != nil {
		return []DeviceInfo{}, err
	}
	return re, nil
}

type ExcelRecord struct {
	Borrower  string    //借用人
	StudentID string    //学号
	ID        uint      //记录id
	Name      string    //设备名称
	State     string    //准备状态
	Type      string    //设备类型
	Usage     string    //用途
	GetTime   time.Time //借出时间
	BackTime  time.Time //实际归还时间
}

type UserRecordInfo struct {
	Record ExcelRecord
}

// GetUserRecords 所有用户的所有借用信息
func GetUserRecords() ([]UserRecordInfo, error) {
	var re []UserRecordInfo

	//获取用户信息
	var users []Model.User
	if err := db.Model(&Model.User{}).Find(&users).Error; err != nil {
		return []UserRecordInfo{}, err
	}
	for _, v := range users {
		//该用户下所有记录
		var records []ExcelRecord
		if err := db.Model(&v).Order("State").Association("Records").Find(&records); err != nil {
			return []UserRecordInfo{}, err
		}

		for _, v2 := range records {
			re = append(re, UserRecordInfo{
				Record: v2,
			})
		}
	}
	return re, nil
}

type DeviceRecordInfo struct {
	//TypeName string
	Record ExcelRecord
}

// GetDeviceRecordByType 获取一个类型的设备记录
func GetDeviceRecordByType(typeName string) ([]DeviceRecordInfo, error) {
	var re []DeviceRecordInfo

	//获取
	var device []Model.Uav
	if err := db.Model(&Model.Uav{}).Where("type = ?", typeName).Find(&device).Error; err != nil {
		return []DeviceRecordInfo{}, err
	}

	var records []ExcelRecord
	if err := db.Model(&device).Order("State").Association("Records").Find(&records); err != nil {
		return []DeviceRecordInfo{}, err
	}

	for _, v := range records {
		re = append(re, DeviceRecordInfo{
			//TypeName: typeName,
			Record: v,
		})
	}

	return re, nil
}

// GetDeviceRecord 单个设备记录
func GetDeviceRecord(uav Model.Uav) ([]DeviceRecordInfo, error) {
	var re []DeviceRecordInfo
	var records []ExcelRecord
	if err := db.Model(&uav).Order("get_time").Association("Records").Find(&records); err != nil {
		return []DeviceRecordInfo{}, err
	}

	for _, v := range records {
		re = append(re, DeviceRecordInfo{Record: v})
	}

	return re, nil
}
