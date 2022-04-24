package Model

import (
	"gorm.io/gorm"
	"time"
)

//该文件存储所有数据库模型

// 设备状态类型：1.空闲 free  2.预定审核 Get under review  3.已预定 scheduled  4.使用 using  5.归还审核 Back under review
// 记录状态类型：1.使用中 using 2.拒绝借用 refuse 3.已归还returned  4.损坏damaged  5.取消cancelled 6.预定审核 Get under review 7.已预定 scheduled  8.归还审核 Back under review

// Uav 设备模型
type Uav struct {
	gorm.Model
	Name  string `json:"name"`                      //设备名称
	State string `gorm:"default:free" json:"state"` //设备状态
	Type  string `json:"type"`                      //设备类型  遥控器Control 电池Battery 无人机Drone
	Uid   string `json:"uid"`                       //设备序号

	StudentID string    `json:"stuid"`     //借用人学号
	Borrower  string    `json:"borrower"`  //借用人姓名
	Phone     string    `json:"phone"`     //借用人电话
	Get_time  time.Time `json:"get_time"`  //借出时间
	Plan_time time.Time `json:"plan_time"` //预计归还时间
	Back_time time.Time `json:"back_time"` //实际归还时间

	Img    string `json:"img"`    //当前图片索引
	Usage  string `json:"usage"`  //当前借用用途
	Remark string `json:"remark"` //设备备注信息
}

// User 用户模型
type User struct {
	gorm.Model
	StudentID string `json:"stuid"`                        //学号
	Pwd       string `json:"pwd"`                          //密码
	Name      string `json:"name"`                         //姓名
	Phone     string `json:"phone"`                        //电话
	IsAdmin   bool   `json:"isadmin" gorm:"default:false"` //判断管理员
	//Count     int    `json:"count" gorm:"default:0"`
}

// Record 历史记录模型
type Record struct {
	gorm.Model
	State     string    `json:"state"`     //状态	使用中using 拒绝借用refuse 已归还returned  损坏damaged  取消cancelled
	Uid       string    `json:"uid"`       //设备序号
	StudentID string    `json:"stuid"`     //学号
	Borrower  string    `json:"name"`      //借用人姓名
	Phone     string    `json:"phone"`     //借用人电话
	Get_time  time.Time `json:"get_time"`  //借出时间
	Plan_time time.Time `json:"plan_time"` //预计归还时间
	Back_time time.Time `json:"back_Time"` //实际归还时间
	Usage     string    `json:"usage"`     //用途

	GetReviewer      string    `json:"getreviewer"`       //借用审核人
	GetReviewTime    time.Time `json:"getreview_time"`    //借用审核时间
	GetReviewResult  string    `json:"getreview_result"`  //借用审核结果  通过passed 失败fail
	GetReviewComment string    `json:"getreview_comment"` //借用审核原因
	GetImg           string    `json:"getimg"`            //借用图片记录

	BackReviewer      string    `json:"backreviewer"`       //归还审核人
	BackReviewTime    time.Time `json:"backreview_time"`    //归还审核时间
	BackReviewResult  string    `json:"backreview_result"`  //归还审核结果  通过passed 失败fail
	BackReviewComment string    `json:"backreview_comment"` //归还审核原因
	BackImg           string    `json:"backimg"`            //归还图片记录

}

// BackRecord 查询历史记录返回模型
type BackRecord struct {
	Borrower  string    `json:"borrower"` //借用人姓名
	StudentID string    `json:"stuid"`    //学号
	State     string    `json:"gbstate" ` //全局状态	已全部归还 All returned 损坏Damaged 使用中Using 审核中Reviewing 已预定Scheduled
	Get_time  time.Time `json:"get_time"` //借出时间
	Usage     string    `json:"usage"`    //用途

	GetReviewer      string    `json:"getreviewer"`       //借用审核人
	GetReviewTime    time.Time `json:"getreview_time"`    //借用审核时间
	GetReviewResult  string    `json:"getreview_result"`  //借用审核结果  通过passed 失败fail
	GetReviewComment string    `json:"getreview_comment"` //借用审核原因
	GetImg           string    `json:"getimg"`            //借用图片记录

	BackReviewer      string    `json:"backreviewer"`       //归还审核人
	BackReviewTime    time.Time `json:"backreview_time"`    //归还审核时间
	BackReviewResult  string    `json:"backreview_result"`  //归还审核结果  通过passed 失败fail
	BackReviewComment string    `json:"backreview_comment"` //归还审核原因
	BackImg           string    `json:"backImg"`            //归还图片记录

	Uav []BackUav `json:"uavs" gorm:"-"` //设备组
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

	Remark string `json:"remark"` //设备备注信息

}

// BackUser 返回用户模型
type BackUser struct {
	Name      string `json:"name"`
	Phone     string `json:"phone" `
	StudentID string `json:"stuid"`
	IsAdmin   bool   `json:"isadmin"`
}

// SearchUav 查询设备模型
type SearchUav struct {
	Name  string `json:"name" form:"name" binding:"-"`   //设备名称
	State string `json:"state" form:"state" binding:"-"` //设备状态
	Type  string `json:"type" form:"type" binding:"-"`   //设备类型
	Uid   string `json:"uid" form:"uid" binding:"-"`     //设备序号
}

// ChangeUav 修改设备模型
type ChangeUav struct {
	Name     string `json:"name"`     //设备名称
	State    string `json:"state"`    //设备状态
	Type     string `json:"type"`     //设备类型
	Uid      string `json:"uid"`      //设备序号（识别设备）
	Borrower string `json:"borrower"` //借用人姓名
	Phone    string `json:"phone"`    //借用人电话
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

// RemarkUav 修改设备备注模型
type RemarkUav struct {
	Uid    string `json:"uid"`    //设备序号
	Remark string `json:"remark"` //备注
}

// BasicUav 查询记录模型
type BasicUav struct {
	Name   string `json:"name"`   //设备名称
	Type   string `json:"type"`   //设备类型
	Uid    string `json:"uid"`    //设备序号
	Remark string `json:"remark"` //设备备注
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
