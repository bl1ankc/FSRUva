package Model

import (
	"gorm.io/gorm"
	"time"
)

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
