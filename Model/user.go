package Model

import (
	"log"
)

// InsertUser 创建新用户ljy
func InsertUser(Name string, Phone string, StudentID string) {

	DB := db.Create(&User{Name: Name, Phone: Phone, StudentID: StudentID})

	if DB.Error != nil {
		log.Fatal(DB.Error.Error())
		return
	}

	return
}
