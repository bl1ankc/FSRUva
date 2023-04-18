package Trans

import (
	"fmt"
	"gorm.io/gorm"
	"main/Model"
	"main/Service"
	"time"
)

var db = Service.GetDB()

// Borrow 借用
func Borrow(uav Model.Uav, user Model.User) error {
	//事务处理
	return db.Transaction(func(tx *gorm.DB) error {
		//贵重审核判断
		if uav.Expensive == true {
			uav.State = "Get under review"
		} else {
			uav.State = "scheduled"
		}

		//生成记录
		record := Model.Record{Name: uav.Name, State: uav.State, Uid: uav.Uid, StudentID: uav.StudentID, Borrower: uav.Borrower, PlanTime: uav.PlanTime, Usage: uav.Usage, GetTime: time.Unix(0, 0), BackTime: time.Unix(0, 0), GetReviewTime: time.Unix(0, 0), BackReviewTime: time.Unix(0, 0)}

		//更新设备借用信息
		if err := tx.Model(&Model.Uav{}).Where("uid = ?", uav.Uid).Updates(uav).Error; err != nil {
			fmt.Println("更新设备失败：", err)
			return err
		}

		//添加设备关联记录
		if err := tx.Model(&uav).Where("id = ?", uav.ID).Association("Records").Append(&record); err != nil {
			fmt.Println("增加一条记录失败：", err)
			return err
		}
		fmt.Println("记录id----------------", record.ID)
		fmt.Println("设备id----------------", uav.ID)
		//更新自定义id索引
		if err := tx.Model(&Model.Uav{}).Where("id = ?", uav.ID).Updates(&Model.Uav{RecordID: record.ID}).Error; err != nil {
			fmt.Println("更新设备记录id失败", err)
			return err
		}

		//添加用户关联记录
		if err := tx.Model(&user).Association("Records").Append(&record); err != nil {
			fmt.Println("添加用户关联失败：", err)
			return err
		}

		return nil
	})
}

// Back 归还
func Back(uav *Model.Uav) error {
	return db.Transaction(func(tx *gorm.DB) error {
		//更新状态
		if err := tx.Model(&uav).Select("state").Updates(Model.Uav{State: "Back under review", BackTime: time.Now().Local()}).Error; err != nil {
			fmt.Println("更新借用状态失败")
			return err
		}

		//更新记录图片
		if err := tx.Model(&Model.Record{}).Where("id = ?", uav.RecordID).Select("back_img").Updates(Model.Record{BackImg: uav.CurImg}).Error; err != nil {
			fmt.Println("更新归还图片失败")
			return err
		}

		//更新记录
		if err := tx.Model(&Model.Record{}).Where("id = ?", uav.RecordID).Updates(&Model.Record{State: "Back under review", BackTime: time.Now().Local()}).Error; err != nil {
			fmt.Println("更新记录借用状态失败")
			return err
		}

		return nil
	})
}

// Get 取走
func Get(uav *Model.Uav) error {
	return db.Transaction(func(tx *gorm.DB) error {
		//更新设备状态,取走时间
		if err := tx.Model(&uav).Select("state", "get_time").Updates(&Model.Uav{State: "using", GetTime: time.Now().Local()}).Error; err != nil {
			return err
		}

		//更新记录图片,取走时间
		if err := tx.Model(&Model.Record{}).Where("id = ?", uav.RecordID).Updates(&Model.Record{State: "using", GetImg: uav.CurImg, GetTime: time.Now().Local()}).Error; err != nil {
			return err
		}

		return nil
	})
}
