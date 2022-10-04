package Routers

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() *gin.Engine {
	r = gin.Default()

	UavRoute()

	AdminRoute()

	UserRoute()

	RecordRoute()

	DepartmentInit()

	return r
}
