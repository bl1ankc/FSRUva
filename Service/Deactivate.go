package Service

//// GetUavStateByUid 通过Uid获取设备状态
//func GetUavStateByUid(u Model.Uav) string {
//	var uav Model.Uav
//	DB := db.Model(&Model.Uav{}).Where("uid = ?", u.Uid).First(&uav)
//
//	if DB.Error != nil {
//		fmt.Println("通过Uid获取设备状态失败：", DB.Error.Error())
//	}
//	return uav.State
//}

//// GetUavByNames 获取对应型号及状态的设备信息
//func GetUavByNames(UavName string, UavType string) []Model.Uav {
//	var uav []Model.Uav
//	DB := db.Model(&Model.Uav{}).Where(Model.Uav{Name: UavName, Type: UavType}).Find(&uav)
//
//	if DB.Error != nil {
//		fmt.Println("获取对应型号及状态的设备信息失败：", DB.Error.Error())
//	}
//
//	return uav
//}

//// UpdateUavUsage 更新设备用途
//func UpdateUavUsage(Uid string, Usage string) error {
//	DB := db.Model(&Model.Uav{}).Where(Model.Uav{Uid: Uid}).Updates(&Model.Uav{Usage: Usage})
//
//	if DB.Error != nil {
//		fmt.Println("更新设备用途失败：", DB.Error.Error())
//		return DB.Error
//	}
//
//	return nil
//}

//// UpdateDevices 强制修改设备数据
//func UpdateDevices(uav Model.Uav) {
//	UpdateDataInUav(uav.Uid, "type", uav.Type)
//	UpdateDataInUav(uav.Uid, "name", uav.Name)
//	UpdateDataInUav(uav.Uid, "location", uav.Location)
//	UpdateDataInUav(uav.Uid, "remark", uav.Remark)
//	//UpdateDataInUav(uav.Uid, "borrower", uav.Borrower)
//	//UpdateDataInUav(uav.Uid, "phone", uav.Phone)
//	//UpdateDataInUav(uav.Uid, "state", uav.State)
//}
//
//// UpdateDataInUav 修改设备单个字符串数据
//func UpdateDataInUav(Uid string, HeadName string, Data string) {
//	if Data != "" {
//		DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: Uid}).Update(HeadName, Data)
//		if DB.Error != nil {
//			fmt.Println("修改设备单个字符串数据失败：", DB.Error.Error())
//		}
//	}
//}

//// UpdateRecordIdinUav 在设备中更新记录ID
//func UpdateRecordIdinUav(Uid string, id uint) bool {
//	DB := db.Model(&Model.Uav{}).Where(&Model.Uav{Uid: Uid}).Updates(&Model.Uav{RecordID: id})
//	if DB.Error != nil {
//		fmt.Println("在设备中更新记录ID失败：", DB.Error.Error())
//		return false
//	}
//	return true
//}

// GetRecordByUid 序列号单一查找
//func GetRecordByUid(uid string) (Model.Record, error) {
//	var data Model.Record
//	if err := db.Model(&Model.Record{}).Where("uid = ?", uid).First(&data).Error; err != nil {
//		return Model.Record{}, err
//	}
//	return data, nil
//}
//
//// UpdateRecord 更新记录
//func UpdateRecord(record Model.Record) error {
//	if DB := db.Model(&record).Updates(&record); DB.Error != nil {
//		return DB.Error
//	}
//	return nil
//}

//// UpdateGetTimeinRecords 记录中更新借用时间
//func UpdateGetTimeinRecords(Uid string) bool {
//
//	//uav := GetUavByUid(Uid)
//	//更新状态
//	//DB := db.Model(&Record{}).Where(&Record{Uid: Uid, StudentID: uav.StudentID, GetTime: uav.GetTime}).Select("GetTime").Updates(&Record{GetTime: time})
//
//	id, flag := GetRecordIdinUav(Uid)
//	if !flag {
//		return false
//	}
//
//	DB := db.Model(&Model.Record{}).Where("id", id).Updates(&Model.Record{GetTime: time.Now()})
//	if DB.Error != nil {
//		fmt.Println("记录中更新借用时间：", DB.Error.Error())
//		return false
//	}
//
//	return true
//}

//// UpdateStudentId 修改学号
//func UpdateStudentId(Name string, Id string) {
//
//	DB := db.Model(&User{}).Where(&User{Name: Name}).Updates(&User{StudentID: Id})
//
//	if DB.Error != nil {
//		log.Fatal(DB.Error.Error())
//	}
//
//	return
//}

//// UpdateUserCount 增加借用次数
//func UpdateUserCount(UserName string, add int) {
//
//	var tmp User
//	DB := db.Model(&User{}).Where(&User{Name: UserName}).Select("count").First(&tmp).Update("count", tmp.Count+add)
//
//	if DB.Error != nil {
//		log.Fatal(DB.Error.Error())
//	}
//
//	return
//}

//// UpdateUserCountByUid 通过无人机序列号查询姓名增加借用次数
//func UpdateUserCountByUid(Uid string, add int) {
//
//	Uav := GetUavByUid(Uid)
//	Borrower := Uav.Borrower
//
//	UpdateUserCount(Borrower, add)
//
//	return
//}
