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
