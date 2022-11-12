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

	GetTime  time.Time `json:"getTime"`  //借出时间
	PlanTime time.Time `json:"planTime"` //预计归还时间
	BackTime time.Time `json:"backTime"` //实际归还时间

	Img   string `json:"Data"`  //当前图片索引
	Usage string `json:"usage"` //当前借用用途

	Location  string `json:"location"`                //设备存放位置
	Remark    string `json:"remark"`                  //设备备注信息
	matter    bool   `json:"matter"`                  //
	Expensive bool   `json:"expensive" gorm :"false"` //是否贵重
	TmpImg    string `json:"tmpImg"`                  //临时图片
}

// CheckUav 审核设备模型
type CheckUav struct {
	Uid     string `json:"uid"`     //设备序号
	Checker string `json:"checker"` //审核人姓名
	Comment string `json:"comment"` //备注原因
	Type    string `json:"type"`    //设备类型
}

// UsingUav 使用中的无人机
type UsingUav struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	State    string `json:"state"`
	GetTime  string `json:"GetTime"`  //借用时间
	PlanTime string `json:"PlanTime"` //预计归还时间
	LastDays int    `json:"lastDays"` //剩余时间
}

type UavType struct {
	gorm.Model
	DepartmentID uint   `gorm:"default:NULL"`
	TypeName     string `json:"typeName"` //设备类型名
	Remark       string `json:"remark"`   //备注
	Img          string `json:"img"`
}

//func (uav Uav) AfterFind(tx *gorm.DB) (err error) {
//
//	tmpType, _ := Service.GetTypeByName(uav.Type)
//	uav.TmpImg, _ = utils.GetPicUrl(tmpType.TypeName)
//
//	return nil
//}
