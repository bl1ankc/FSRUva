package Model

import (
	"fmt"
	"log"
)

//该文件储存关于设备的处理方法

//获取对应状态的无人机数据
func GetUvasByState(UvaState string, UvaType string) (uva []Uva, err error) {
	DB := db.Model(&Uva{}).Where("State = ? and Type = ?", UvaState, UvaType).Find(&uva)
	if DB.Error != nil {
		fmt.Println("GetUvasByState Error")
		log.Fatal(DB.Error)
	}
	return
}

//
func Insert() {
	DB := db.Create(&Uva{Name: "Firstuva", Uid: "123456", Type: "Uva", State: "free"})
	if DB.Error != nil {
		fmt.Println("Insert error")
		log.Fatal(DB.Error)
	}
	return
}
