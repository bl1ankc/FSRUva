package Model

import (
	"gorm.io/gorm"
	"time"
)

//该文件存储所有数据库模型

// Uav 设备模型
type Uav struct {
	gorm.Model
	Name  string `json:"name"`                      //设备名称
	State string `gorm:"default:free" json:"state"` //设备状态
	Type  string `json:"type"`                      //设备类型
	Uid   string `json:"uid"`                       //设备序号

	Borrower  string    `json:"borrower"`  //借用人姓名
	Phone     string    `json:"phone"`     //借用人电话
	Get_time  time.Time `json:"get_time"`  //借出时间
	Plan_time time.Time `json:"plan_time"` //预计归还时间
	Back_time time.Time `json:"back_time"` //实际归还时间
}

// User 用户模型
type User struct {
	gorm.Model
	Name      string `json:"name"`
	Phone     string `json:"tel"`
	pwd       string
	StudentID string
}

// Record 历史记录模型
type Record struct {
	gorm.Model
	State     string    `json:"state"`     //状态
	Uid       string    `json:"uid"`       //设备序号
	Borrower  string    `json:"name"`      //借用人姓名
	Phone     string    `json:"tel"`       //借用人电话
	Get_time  time.Time `json:"get_Time"`  //借出时间
	Plan_time time.Time `json:"plan_Time"` //预计借用时长
	Real_time time.Time `json:"real_Time"` //实际借用时长
	Back_time time.Time `json:"back_Time"` //归还时间
}

// BackRecord 查询历史记录返回模型
type BackRecord struct {
	Borrower  string    `json:"borrower"`      //借用人姓名
	State     string    `json:"state" `        //状态
	Get_time  time.Time `json:"get_Time"`      //借出时间
	Plan_time time.Time `json:"plan_Time"`     //预计借用时长
	Real_time time.Time `json:"real_Time"`     //实际借用时长
	Back_time time.Time `json:"back_Time"`     //归还时间
	Uav       []BackUav `json:"uavs" gorm:"-"` //设备
}

// BackUav 设备模型
type BackUav struct {
	Name      string    `json:"name"`      //设备名称
	State     string    `json:"state"`     //设备状态
	Type      string    `json:"type"`      //设备类型
	Uid       string    `json:"uid"`       //设备序号
	Get_time  time.Time `json:"get_time"`  //借出时间
	Plan_time time.Time `json:"plan_time"` //预计归还时间
	Back_time time.Time `json:"back_time"` //实际归还时间
}
