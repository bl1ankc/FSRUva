package Routers

import (
	"github.com/gin-gonic/gin"
)

var r *gin.RouterGroup

func InitRouter() *gin.Engine {
	def := gin.Default()

	r = def.Group("/api")

	UavRoute()

	AdminRoute()

	UserRoute()

	RecordRoute()

	DepartmentInit()

	ExcelInit()

	return def
}
