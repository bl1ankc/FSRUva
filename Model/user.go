package Model

import (
	"fmt"
	"log"
)

// InsertUser 创建新用户
func InsertUser(Name string, Phone string, StudentID string, Pwd string) string {

	//判断用户是否存在
	var count int64
	DB := db.Model(&User{}).Where(&User{StudentID: StudentID}).Count(&count)
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
	if count != 0 {
		return "用户已存在"
	}

	//创建数据
	DB = db.Create(&User{Name: Name, Phone: Phone, StudentID: StudentID, Pwd: Pwd})
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
	return "OK"
}

// UpdatePhone 修改电话号码
func UpdatePhone(Stuid string, Phone string) bool {

	DB := db.Model(&User{}).Where(&User{StudentID: Stuid}).Updates(&User{Phone: Phone})

	if DB.Error != nil {
		fmt.Println(Stuid, "电话更改失败")
		return false
	}

	return true
}

// UpdatePwd 修改密码
func UpdatePwd(Stuid string, OldPwd string, NewPwd string) (string, bool) {
	//验证旧密码
	var old string
	DB := db.Model(&User{}).Where(&User{StudentID: Stuid}).Select("pwd").First(&old)

	if DB.Error != nil {
		fmt.Println(Stuid, "密码更改失败")
		return "密码更改失败", false
	}

	if old == OldPwd {
		//更改新密码
		DB = db.Model(&User{}).Where(&User{StudentID: Stuid}).Updates(&User{Pwd: NewPwd})

		if DB.Error != nil {
			fmt.Println(Stuid, "密码更改失败")
			return "密码更改失败", false
		}

		return "密码更改成功", true
	} else {
		return "密码错误", false
	}

}

// UpdateAdmin 修改管理员标识
func UpdateAdmin(Stuid string, Isadmin bool) {
	//用户变量
	var user User

	//更新用户管理权限
	db.Model(&User{}).Where(&User{StudentID: Stuid}).First(&user).Update("is_admin", Isadmin)
}

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

// GetUserByName 通过名字查找用户信息
func GetUserByName(Name string) (BackUser, error) {

	var user BackUser
	DB := db.Model(&User{}).Where(&User{Name: Name}).Find(&user)
	if DB.Error != nil {
		return user, db.Error
	}
	return user, nil
}

// GetUserByIDToLogin 通过名字查找用户登录所需信息
func GetUserByIDToLogin(stuid string) (User, error) {

	var user User
	DB := db.Model(&User{}).Where(&User{StudentID: stuid}).Find(&user)
	if DB.Error != nil {
		return user, db.Error
	}
	return user, nil
}

// GetAllUsers 获取所有用户
func GetAllUsers() []BackUser {

	var user []BackUser
	DB := db.Model(&User{}).Find(&user)
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
	return user
}

// GetUserByID 通过学号查找用户信息
func GetUserByID(Stuid string) BackUser {

	var user BackUser
	DB := db.Model(&User{}).Where(&User{StudentID: Stuid}).Find(&user)
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
	return user
}
