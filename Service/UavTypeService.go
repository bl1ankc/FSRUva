package Service

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"main/Model"
)

// AddUavType 增加设备类型
func AddUavType(uavType Model.UavType, department Model.Department) (error, uint) {
	uavType.DepartmentName = department.DepartmentName
	return db.Transaction(func(tx *gorm.DB) error {
		if err := db.Model(&department).Association("Types").Append(&uavType); err != nil {
			log.Printf(err.Error())
			return err
		}
		if err := db.Model(&uavType).Updates(uavType).Error; err != nil {
			return err
		}
		return nil
	}), uavType.ID
}

// GetUavType 获取设备类型列表
func GetUavType() (bool, []Model.UavType) {
	var types []Model.UavType
	DB := db.Model(&Model.UavType{}).Find(&types)

	if DB.Error != nil {
		fmt.Println("获取设备类型列表：", DB.Error.Error())
		return false, []Model.UavType{}
	}

	return true, types
}

// GetType 获取单独设备类型
func GetType(id uint) (Model.UavType, error) {
	var uavType Model.UavType
	if err := db.Model(&Model.UavType{}).Where("id = ?", id).First(&uavType).Error; err != nil {
		return Model.UavType{}, err
	}
	return uavType, nil
}

// GetTypeByName 获取单独设备类型
func GetTypeByName(typeName string) (Model.UavType, error) {
	var uavType Model.UavType
	if err := db.Model(&Model.UavType{}).Where("type_Name = ?", typeName).First(&uavType).Error; err != nil {
		return Model.UavType{}, err
	}
	return uavType, nil
}

// RemoveUavType 删除设备类型
func RemoveUavType(TypeName string) bool {

	DB := db.Where("type_name=?", TypeName).Delete(&Model.UavType{})

	if DB.Error != nil {
		fmt.Println("删除设备类型失败：", DB.Error.Error())
		return false
	}

	return true
}

// UpdateUavType 更新设备类型
func UpdateUavType(uavType Model.UavType) error {
	return db.Transaction(func(tx *gorm.DB) error {
		//获取实例
		var instance Model.UavType
		if err := db.Model(&Model.UavType{}).Where("id = ?", uavType.ID).First(&instance).Error; err != nil {
			return err
		}
		//替换关联
		if uavType.DepartmentID != 0 && uavType.DepartmentID != instance.DepartmentID {
			//删除就部门关联
			if err := db.Model(&Model.Department{}).Where("id = ?", instance.DepartmentID).Association("Types").Delete(&instance); err != nil {
				return err
			}
			//获取新部门
			var department Model.Department
			if err := db.Model(&Model.Department{}).Where("id = ?", uavType.DepartmentID).First(&department).Error; err != nil {
				return err
			}
			//添加新部门关联
			if err := db.Model(&department).Association("Types").Append(&instance); err != nil {
				return err
			}
			fmt.Println(department.DepartmentName, "----------------")
			//更新部门名称
			if err := db.Model(&Model.UavType{}).Where("id = ?", uavType.ID).Updates(&Model.UavType{DepartmentName: department.DepartmentName}).Error; err != nil {
				return err
			}
			return nil
		}
		//相同直接更新
		return db.Model(&Model.UavType{}).Where("id = ?", uavType.ID).Updates(&uavType).Error
	})
}

// UpdateTypeImg 更新类型图片
func UpdateTypeImg(id uint, img string) error {
	uavType, err := GetType(id)
	if err != nil {
		return err
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
		//更新图片
		if err := db.Model(&uavType).Update("img", img).Error; err != nil {
			return err
		}

		//更新设备图片
		if err := db.Model(&Model.Uav{}).Where("type = ?", uavType.TypeName).Select("img").Updates(&Model.Uav{Img: uavType.Img}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
