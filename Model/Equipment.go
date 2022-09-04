package Model

import (
	"gorm.io/gorm"
	"time"
)

// 设备状态类型：1.空闲 free  2.预定审核 Get under review  3.已预定 scheduled  4.使用 using  5.归还审核 Back under review
// 记录状态类型：1.使用中 using 2.拒绝借用 refuse 3.已归还returned  4.损坏damaged  5.取消cancelled 6.预定审核 Get under review 7.已预定 scheduled  8.归还审核 Back under review

// Uav 设备模型
type Uav struct {
	gorm.Model
	Name  string `json:"name"`                      //设备名称
	State string `gorm:"default:free" json:"state"` //设备状态
	Type  string `json:"type"`                      //设备类型  遥控器Control 电池Battery 无人机Drone
	Uid   string `json:"uid"`                       //设备序号

	StudentID string `json:"stuid"`    //借用人学号
	Borrower  string `json:"borrower"` //借用人姓名
	Phone     string `json:"phone"`    //借用人电话

	RecordID uint //记录ID

	Get_time  time.Time `json:"get_time"`  //借出时间
	Plan_time time.Time `json:"plan_time"` //预计归还时间
	Back_time time.Time `json:"back_time"` //实际归还时间

	Img   string `json:"Data"`  //当前图片索引
	Usage string `json:"usage"` //当前借用用途

	Location string `json:"location"` //设备存放位置
	Remark   string `json:"remark"`   //设备备注信息
}

// BackUav 返回设备模型
type BackUav struct {
	Name      string    `json:"name"`      //设备名称
	State     string    `json:"state"`     //设备状态
	Type      string    `json:"type"`      //设备类型
	Uid       string    `json:"uid"`       //设备序号
	Get_time  time.Time `json:"get_time"`  //借出时间
	Plan_time time.Time `json:"plan_time"` //预计归还时间
	Back_time time.Time `json:"back_time"` //实际归还时间
	Borrower  string    `json:"borrower"`  //借用人

	Location string `json:"location"` //设备存放位置
	Remark   string `json:"remark"`   //设备备注信息

}

// SearchUav 查询设备模型
type SearchUav struct {
	Name  string `json:"name" form:"name" binding:"-"`   //设备名称
	State string `json:"state" form:"state" binding:"-"` //设备状态
	Type  string `json:"type" form:"type" binding:"-"`   //设备类型
	Uid   string `json:"uid" form:"uid" binding:"-"`     //设备序号
}

// BorrowUav 借用设备模型
type BorrowUav struct {
	Uid       string    `json:"uid"`       //设备序号
	StudentID string    `json:"stuid"`     //学号
	Borrower  string    `json:"borrower"`  //借用人姓名
	Phone     string    `json:"phone"`     //借用人电话
	Plan_time time.Time `json:"plan_time"` //预计归还时间
	Usage     string    `json:"usage"`     //用途
}

// CheckUav 审核设备模型
type CheckUav struct {
	Uid     string `json:"uid"`     //设备序号
	Checker string `json:"checker"` //审核人姓名
	Comment string `json:"comment"` //备注原因
}

// UsingUav 使用中的无人机
type UsingUav struct {
	Uid       string `json:"uid"`
	Name      string `json:"name"`
	State     string `json:"state"`
	Get_Time  string `json:"get_time"`  //借用时间
	Plan_Time string `json:"plan_time"` //预计归还时间
	LastDays  int    `json:"lastDays"`  //剩余时间
}

type UavType struct {
	gorm.Model
	TypeName string `json:"typeName"` //设备类型名
	Remark   string `json:"remark"`   //备注
}
