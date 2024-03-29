package Service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"main/Model"
)

// GetUser 获取用户实例
func GetUser(id uint) (Model.User, error) {
	var user Model.User
	if err := db.Model(&Model.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// InsertUser 创建新用户
func InsertUser(Name string, Phone string, StudentID string, Pwd string) (bool, string) {

	//判断用户是否存在
	var count int64
	DB := db.Model(&Model.User{}).Where(&Model.User{StudentID: StudentID}).Count(&count)
	if DB.Error != nil {
		fmt.Println("创建新用户失败1：", DB.Error.Error())
	}
	if count != 0 {
		return false, "用户已存在"
	}

	h := md5.New()
	h.Write([]byte(Pwd))
	ciphertext := hex.EncodeToString(h.Sum(nil))

	//创建数据
	DB = db.Create(&Model.User{Name: Name, Phone: Phone, StudentID: StudentID, Pwd: ciphertext})
	if DB.Error != nil {
		fmt.Println("创建新用户失败2：", DB.Error.Error())
		return false, "注册失败，请联系管理员"
	}
	return true, "注册成功"
}

// DeleteUser 删除用户数据
func DeleteUser(user Model.User) error {
	return db.Model(&Model.User{}).Where("id = ?", user.ID).Delete(&user).Error
}

// UpdatePhone 修改电话号码
func UpdatePhone(Stuid string, Phone string) bool {

	DB := db.Model(&Model.User{}).Where(&Model.User{StudentID: Stuid}).Updates(&Model.User{Phone: Phone})

	if DB.Error != nil {
		fmt.Println(Stuid, "电话更改失败：", DB.Error.Error())
		return false
	}

	return true
}

// UpdatePwd 修改密码
func UpdatePwd(Stuid string, OldPwd string, NewPwd string) (string, bool) {
	//密码加密
	h1 := md5.New()
	h1.Write([]byte(OldPwd))
	ciphertext1 := hex.EncodeToString(h1.Sum(nil))

	//验证旧密码
	var old string
	DB := db.Model(&Model.User{}).Where(&Model.User{StudentID: Stuid}).Select("pwd").First(&old)

	if DB.Error != nil {
		fmt.Println(Stuid, "密码更改失败1：", DB.Error.Error())
		return "密码更改失败", false
	}

	if old == ciphertext1 {
		h2 := md5.New()
		h2.Write([]byte(NewPwd))
		ciphertext2 := hex.EncodeToString(h2.Sum(nil))

		//更改新密码
		DB = db.Model(&Model.User{}).Where(&Model.User{StudentID: Stuid}).Updates(&Model.User{Pwd: ciphertext2})

		if DB.Error != nil {
			fmt.Println(Stuid, "密码更改失败2：", DB.Error.Error())
			return "密码更改失败", false
		}

		return "密码更改成功", true
	} else {
		return "密码错误", false
	}

}

// GetUserByName 通过名字查找用户信息
func GetUserByName(Name string) (Model.BackUser, error) {

	var user Model.BackUser
	DB := db.Model(&Model.User{}).Where(&Model.User{Name: Name}).Find(&user)
	if DB.Error != nil {
		return user, db.Error
	}
	return user, nil
}

// GetUserByIDToLogin 通过名字查找用户登录所需信息
func GetUserByIDToLogin(stuid string) (Model.User, error) {

	var user Model.User
	DB := db.Model(&Model.User{}).Where(&Model.User{StudentID: stuid}).Find(&user)
	if DB.Error != nil {
		return user, db.Error
	}
	return user, nil
}

// GetAllUsers 获取所有用户
func GetAllUsers() []Model.BackUser {

	var user []Model.BackUser
	DB := db.Model(&Model.User{}).Find(&user)
	if DB.Error != nil {
		fmt.Println("获取所有用户失败：", DB.Error.Error())
	}
	return user
}

// GetUserByID 通过学号查找登录用户信息
func GetUserByID(Stuid string) Model.BackUser {

	var user Model.BackUser
	DB := db.Model(&Model.User{}).Where(&Model.User{StudentID: Stuid}).First(&user)
	if DB.Error != nil {
		fmt.Println("通过学号查找用户信息失败：", DB.Error.Error())
	}
	return user
}

// UpdateUserInfo 更新用户昵称头像
func UpdateUserInfo(Stuid string, Nickname string, AvatarUrl string, Openid string, Unionid string) bool {
	DB := db.Model(&Model.User{}).Where(&Model.User{StudentID: Stuid}).Updates(&Model.User{NickName: Nickname, AvatarUrl: AvatarUrl, Openid: Openid, Unionid: Unionid})
	if DB.Error != nil {
		fmt.Println("更新用户昵称头像失败：", DB.Error.Error())
		return false
	}
	return true
}
