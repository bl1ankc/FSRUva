package Model

import (
	"fmt"
	"gorm.io/gorm"
	"main/utils"
	"time"
)

// Record 历史记录模型
type Record struct {
	gorm.Model
	UavID     uint
	Name      string    `json:"name"`     //设备名称
	State     string    `json:"state"`    //状态	使用中using 拒绝借用refuse 已归还returned  损坏damaged  取消cancelled
	Uid       string    `json:"uid"`      //设备序号
	Type      string    `json:"type"`     //设备类型
	StudentID string    `json:"stuid"`    //学号
	Borrower  string    `json:"borrower"` //借用人姓名
	Phone     string    `json:"phone"`    //借用人电话
	GetTime   time.Time `json:"getTime"`  //借出时间
	PlanTime  time.Time `json:"planTime"` //预计归还时间
	BackTime  time.Time `json:"backTime"` //实际归还时间
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
	BackImg           string    `json:"backimg"`            //归还图片记录

	TmpImg string `json:"tmpImg"` //临时图片（类型图片）
}

// BackRecord 查询历史记录返回模型
type BackRecord struct {
	Borrower  string    `json:"borrower"` //借用人姓名
	StudentID string    `json:"stuid"`    //学号
	State     string    `json:"gbstate" ` //全局状态	已全部归还 All returned 损坏Damaged 使用中Using 审核中Reviewing 已预定Scheduled
	GetTime   time.Time `json:"GetTime"`  //借出时间
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

	Uav []Uav `json:"uavs" gorm:"-"` //设备组
}

func (r *Record) AfterFind(tx *gorm.DB) (err error) {
	var uavType UavType
	var uav Uav
	if r.Type == "" {
		if err = tx.Model(&Uav{}).Where(Uav{Uid: r.Uid}).First(&uav).Error; err != nil {
			return err
		}
		r.Type = uav.Type
		r.TmpImg = uav.TmpImg
		fmt.Println("查询失败" + r.TmpImg + " another" + r.Type)
		return nil
	}
	if err = tx.Model(&UavType{}).Where(UavType{TypeName: r.Type}).First(&uavType).Error; err != nil {
		return err
	}
	tmpImg, _ := utils.GetPicUrl(uavType.Img)
	r.TmpImg = tmpImg
	fmt.Println("sssssssssssssssss" + r.Type + r.TmpImg)
	return nil
}
