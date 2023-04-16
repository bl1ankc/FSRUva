package Trans

import (
	"gorm.io/gorm"
	"main/Model"
	"time"
)

// ForcedGet 强制取走
func ForcedGet(uav *Model.Uav) error {
	return db.Transaction(func(tx *gorm.DB) error {
		now := time.Now().Local()
		//设备更新
		if err := db.Model(&Model.Uav{}).
			Where("uid = ?", uav.Uid).
			Updates(&Model.Uav{
				State:   "using",
				GetTime: now,
			}).Error; err != nil {
			return err
		}
		//记录更新
		if err := db.Model(&Model.Record{}).
			Where("id = ?", uav.RecordID).
			Updates(&Model.Record{
				State:         "using",
				GetTime:       now,
				GetReviewTime: now,
			}).Error; err != nil {
			return err
		}

		return nil
	})
}

// ForcedBack 强制归还
func ForcedBack(uav *Model.Uav) error {
	return db.Transaction(func(tx *gorm.DB) error {
		now := time.Now().Local()
		//Device Update
		if err := db.Model(&Model.Uav{}).
			Where("id = ?", uav.ID).
			Updates(&Model.Uav{
				State:    "free",
				BackTime: now,
			}).Error; err != nil {
			return err
		}
		//Record Update
		if err := db.Model(&Model.Record{}).
			Where("id = ?", uav.RecordID).
			Updates(&Model.Record{
				State:            "returned",
				BackTime:         now,
				BackReviewResult: "passed",
				BackReviewTime:   now,
			}).Error; err != nil {
			return err
		}
		return nil
	})
}
