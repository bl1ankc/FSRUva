package Service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"main/Model"
)

// CreateDepartment 添加部门信息
func CreateDepartment(department Model.Department) error {
	if err := db.Model(&Model.Department{}).Create(&department).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDepartment 删除部门信息
func DeleteDepartment(department Model.Department) error {
	if err := db.Model(&department).Association("Types").Clear(); err != nil {
		fmt.Println("清理类型关联失败")
		return err
	}
	if err := db.Model(&department).Delete(&department).Error; err != nil {
		fmt.Println("删除失败")
		return err
	}
	return nil
}

// GetDepartmentList 获取部门信息列表
func GetDepartmentList() ([]Model.Department, error) {
	var department []Model.Department
	if err := db.Model(&Model.Department{}).Find(&department).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return []Model.Department{}, err
		}
	}
	return department, nil
}

// GetDepartment 获取部门信息
func GetDepartment(id uint) (Model.Department, error) {
	var department Model.Department
	if err := db.Model(&Model.Department{}).Where("id = ?", id).First(&department).Error; err != nil {
		return Model.Department{}, err
	}
	return department, nil
}

// AddTypeToDepartment 向部门添加类型
func AddTypeToDepartment(uavType Model.UavType, department Model.Department) error {
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
	})

	return nil
}

// DeleteTypeFromDepartment 删除关联
func DeleteTypeFromDepartment(uavType Model.UavType, department Model.Department) error {
	if err := db.Model(&department).Association("Types").Delete(&uavType); err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

// GetDepartmentTypes 获取部门下所有类型
func GetDepartmentTypes(department Model.Department) ([]Model.UavType, error) {
	var types []Model.UavType

	if err := db.Model(&department).Association("Types").Find(&types); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return []Model.UavType{}, err
		}
	}

	return types, nil
}
