package Trans

import (
	"gorm.io/gorm"
	"main/Model"
	"time"
)

var get = []string{"state", "get_reviewer", "get_time", "get_review_result", "get_review_comment", "get_review_time"}
var back = []string{"state", "back_reviewer", "back_time", "back_review_result", "back_review_comment", "back_review_time"}
var GetPass = []string{"scheduled", "scheduled", "passed"}
var BackPass = []string{"free", "returned", "passed"}
var GetRefuse = []string{"free", "refuse", "fail"}
var BackRefuse = []string{"using", "using", "fail"}

//返回对应状态更新数据与更新列
func getState(str string) ([]string, []string) {
	switch str {
	case "GetPass":
		return GetPass, get
	case "BackPass":
		return BackPass, back
	case "GetRefuse":
		return GetRefuse, get
	case "BackRefuse":
		return BackRefuse, back
	default:
		return []string{}, []string{}
	}

}

func Review(uav *Model.CheckUav, Judge string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		//数据获取定义
		var now = time.Now().Local()
		state, col := getState(Judge)
		//实例获取
		var instance Model.Uav
		if err := db.Model(&Model.Uav{}).Where("uid = ?", uav.Uid).First(&instance).Error; err != nil {
			return err
		}
		////归还时间是否更新
		//backTime := func() time.Time {
		//	if state[1] == "returned" {
		//		return now
		//	} else {
		//		return instance.BackTime
		//	}
		//}
		//设备更新
		if err := db.Model(&instance).Updates(&Model.Uav{State: state[0]}).Error; err != nil {
			return err
		}
		//记录更新
		if err := db.Model(&Model.Record{}).Where("id = ?", instance.RecordID).
			Select(col).
			Updates(&Model.Record{
				State:        state[1],
				GetReviewer:  uav.Checker,
				BackReviewer: uav.Checker,

				GetReviewResult:  state[2],
				BackReviewResult: state[2],

				GetReviewComment:  uav.Comment,
				BackReviewComment: uav.Comment,

				GetReviewTime:  now,
				BackReviewTime: now}).Error; err != nil {
			return err
		}

		return nil
	})
}
