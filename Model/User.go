package Model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Name      string `json:"name"`                         //姓名
	Phone     string `json:"phone"`                        //电话
	StudentID string `json:"stuid"gorm:"unique"`           //学号
	Pwd       string `json:"pwd"`                          //密码
	IsAdmin   bool   `json:"isadmin" gorm:"default:false"` //判断管理员
	AdminType int    `json:"adminType"gorm:"default:0"`
	//Count     int    `json:"count" gorm:"default:0"`
	NickName  string   `json:"nickName"`  //昵称
	AvatarUrl string   `json:"avatarUrl"` //头像
	Openid    string   `json:"openid"`    //微信openid
	Unionid   string   `json:"unionid"`   //微信unionid
	Records   []Record `json:"records"`
}

// BackUser 返回用户模型
type BackUser struct {
	gorm.Model
	Name      string `json:"name"`
	Phone     string `json:"phone" `
	StudentID string `json:"stuid"`
	IsAdmin   bool   `json:"IsAdmin"`
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	AdminType int    `json:"adminType"`
}
