package Controller

import (
	"github.com/gin-gonic/gin"
	"main/Model"
	"main/Service"
	"main/Service/Status"
	"strconv"
)

// AddDepartment 添加新部门
func AddDepartment(c *gin.Context) {
	var code int
	var department Model.Department

	if err := c.ShouldBindJSON(&department); err != nil {
		code = Status.FailToBindJson
		c.JSON(code, R(code, nil, "绑定结构体失败，检查数据是否传输正确"))
		return
	}

	if err := Service.CreateDepartment(department); err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "创建部门失败，检查服务器错误"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "添加部门成功"))
	return
}

// RemoveDepartment 删除部门
func RemoveDepartment(c *gin.Context) {
	var code int

	id, err := strconv.Atoi(c.Query("departmentID"))
	if err != nil {
		code = Status.FailToGetQuery
		c.JSON(code, R(code, nil, "获取参数失败，检查传入参数是否正确"))
		return
	}

	department, err := Service.GetDepartment(uint(id))
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "查询失败，检查传入参数是否正确，若正确联系后端人员"))
		return
	}

	err = Service.DeleteDepartment(department)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "删除失败，检查函数及数据库"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "删除成功"))
	return
}

// DepartmentList 获取部门列表
func DepartmentList(c *gin.Context) {
	var code int

	data, err := Service.GetDepartmentList()
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "获取列表失败，检查程序或服务器"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, data, "获取列表成功"))
	return
}

// DepartmentTypes 获取部门下类型
func DepartmentTypes(c *gin.Context) {
	var code int

	id, err := strconv.Atoi(c.Query("departmentID"))
	if err != nil {
		code = Status.FailToGetQuery
		c.JSON(code, R(code, nil, "获取参数失败，检查传入参数是否正确"))
		return
	}

	department, err := Service.GetDepartment(uint(id))
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "查询失败，检查传入参数是否正确，若正确联系后端人员"))
		return
	}

	types, err := Service.GetDepartmentTypes(department)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "查询失败，检查传入参数是否正确，若正确联系后端人员"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, types, "获取类型列表成功"))
	return
}

// TypeToDepartment 添加类型
func TypeToDepartment(c *gin.Context) {
	var code int
	var departmentID, typeID int
	var err error

	departmentID, err = strconv.Atoi(c.Query("departmentID"))
	typeID, err = strconv.Atoi(c.Query("typeID"))

	department, err := Service.GetDepartment(uint(departmentID))
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "查询失败，检查传入参数是否正确，若正确联系后端人员"))
		return
	}

	uavType, err := Service.GetType(typeID)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "查询失败，检查传入参数是否正确，若正确联系后端人员"))
		return
	}

	err = Service.AddTypeToDepartment(uavType, department)
	if err != nil {
		code = Status.FuncFail
		c.JSON(code, R(code, nil, "关联添加失败"))
		return
	}

	code = Status.OK
	c.JSON(code, R(code, nil, "添加设备类型到部门成功"))
	return
}
