package Routers

import (
	"main/Controller"
	"main/Mid"
)

func DepartmentInit() {
	g := r.Group("/Department", Mid.AuthRequired())
	{
		//获取部门列表
		g.GET("/GetDepartmentList", Controller.DepartmentList)

		//获取部门下类型列表
		g.GET("/GetDepartmentTypes", Controller.DepartmentTypes)

		a := g.Group("/Admin", Mid.VerifyAdmin())
		{
			//添加部门
			a.POST("/AddDepartment", Controller.AddDepartment)

			//删除部门
			a.DELETE("/RemoveDepartment", Controller.RemoveDepartment)

			//添加类型到部门
			a.GET("/AddTypeToDepartment", Controller.TypeToDepartment)
		}
	}
}
