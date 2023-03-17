package Service

import (
	"fmt"
	"main/Model"
)

// UpdateAdmin 修改管理员标识
func UpdateAdmin(Stuid string, Isadmin bool) bool {
	//用户变量
	var user Model.User

	//更新用户管理权限
	DB := db.Model(&Model.User{}).Where(&Model.User{StudentID: Stuid}).First(&user).Update("is_admin", Isadmin)
	if DB.Error != nil {
		fmt.Println("修改管理员标识失败：", DB.Error.Error())
		return false
	}
	return true
}

func UpdateAdminType(stuID string, Type int) error {

	result := db.Model(&Model.User{}).Where(&Model.User{StudentID: stuID}).Updates(map[string]interface{}{
		"AdminType": Type,
	})
	if result.Error != nil {
		fmt.Println("修改管理员类型失败:", result.Error.Error())
	}
	return result.Error
}

//// VerifyTeacher 验证教师
//func VerifyTeacher(stuID string) bool {
//	user := GetUserByID(stuID)
//
//	user.
//}
