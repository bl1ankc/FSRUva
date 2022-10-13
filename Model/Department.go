package Model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	DepartmentName string `json:"departmentName"`
	Types          []UavType
}
