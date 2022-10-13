package Service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main/Model"
)

var db *gorm.DB //数据库指针

//初始化数据库
func init() {
	var err error

	dsn := "root:Qq3318055.@tcp(127.0.0.1:3306)/xdstest?parseTime=true&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	//自动迁移
	db.AutoMigrate(&Model.Uav{}, &Model.Record{}, &Model.User{}, &Model.UavType{}, &Model.Department{})

}
