package Transaction

import (
	"fmt"
	"gorm.io/gorm"
	"main/Model"
	"main/Service"
	"time"
)

var db = Service.GetDB()

func Borrow(uav *Model.Uav) error {

	//事务处理
	if err := db.Transaction(func(tx *gorm.DB) error {
		if uav.Expensive == true {
			uav.State = "Get under review"
		} else {
			uav.State = "scheduled"
		}

		if err := tx.Create(&Model.Record{Name: uav.Uid, State: uav.State, Uid: uav.Uid, StudentID: uav.StudentID, Borrower: uav.Borrower, PlanTime: uav.PlanTime, Usage: uav.Usage, GetTime: time.Now(), BackTime: time.Unix(0, 0), GetReviewTime: time.Unix(0, 0), BackReviewTime: time.Unix(0, 0)}).Select("id").Find(&uav.RecordID).Error; err != nil {
			fmt.Println("增加一条记录失败：", err)
			return err
		}

		if err := tx.Model(&Model.Uav{}).Where("uid = ?", uav.Uid).Updates(uav).Error; err != nil {
			fmt.Println("更新设备失败：", err)
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
