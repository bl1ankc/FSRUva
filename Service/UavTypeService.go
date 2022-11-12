package Service

import (
	"fmt"
	"main/Model"
)

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

// AddUavType 增加设备类型
func AddUavType(typeName string, remark string) (error, uint) {

	uavType := Model.UavType{TypeName: typeName, Remark: remark}
	DB := db.Model(&Model.UavType{}).Create(&uavType)

	if DB.Error != nil {
		fmt.Println("增加设备类型失败：", DB.Error.Error())
		return DB.Error, 0
	}

	return nil, uavType.ID
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
	err := db.Model(&uavType).Updates(uavType).Error
	return err
}

// UpdateTypeImg 更新类型图片
func UpdateTypeImg(id uint, img string) error {
	uavType, err := GetType(id)
	if err != nil {
		return err
	}
	if err := db.Model(&uavType).Update("img", img).Error; err != nil {
		return err
	}
	return nil
}
