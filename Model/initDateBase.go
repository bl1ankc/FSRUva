package Model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB //数据库指针

//初始化数据库
func init() {
	var err error

	dsn := "root:123456@tcp(127.0.0.1:3306)/FSRUav?parseTime=true"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	//自动迁移
	db.AutoMigrate(&Uav{}, &Record{}, &User{})
}
