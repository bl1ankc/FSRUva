package Model

import (
	"log"
)

// InsertUser 创建新用户
func InsertUser(Name string, Phone string, StudentID string) string {

	//判断用户是否存在
	var count int64
	DB := db.Model(&User{}).Where(&User{Name: Name}).Count(&count)
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
	if count != 0 {
		return "用户已存在"
	}

	//创建数据
	DB = db.Create(&User{Name: Name, Phone: Phone, StudentID: StudentID})
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
	return "OK"
}

// UpdatePhone 修改电话号码
func UpdatePhone(Name string, Phone string) {

	DB := db.Model(&User{}).Where(&User{Name: Name}).Updates(&User{Phone: Phone})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}

	return
}

// UpdateStudentId 修改学号
func UpdateStudentId(Name string, Id string) {

	DB := db.Model(&User{}).Where(&User{Name: Name}).Updates(&User{StudentID: Id})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}

	return
}

// UpdateUserCount 增加借用次数
func UpdateUserCount(UserName string, add int) {

	var tmp User
	DB := db.Model(&User{}).Where(&User{Name: UserName}).Select("count").First(&tmp).Update("count", tmp.Count+add)

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}

	return
}

// UpdateUserCountByUid 通过无人机序列号查询姓名增加借用次数
func UpdateUserCountByUid(Uid string, add int) {

	Uav := GetUavByUid(Uid)
	Borrower := Uav.Borrower

	UpdateUserCount(Borrower, add)

	return
}

// GetUserByName 通过名字查找用户信息
func GetUserByName(Name string) BackUser {

	var user BackUser
	DB := db.Model(&User{}).Where(&User{Name: Name}).Find(&user)
	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
	}
	return user
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
