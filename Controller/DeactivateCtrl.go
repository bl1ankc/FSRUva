package Controller

////UpdateUavRemark 修改设备备注信息
//func UpdateUavRemark(c *gin.Context) {
//	var remark Model.Uav
//	//结构体绑定
//	if err := c.BindJSON(&remark); err != nil {
//		fmt.Println("修改设备备注信息数据绑定失败：", err.Error())
//		c.JSON(400, gin.H{"msg": "参数格式错误"})
//	}
//
//	Service.UpdateUavRemark(remark.Uid, remark.Remark)
//	c.JSON(200, gin.H{"desc": "修改成功"})
//}

// ForceUpdateDevices 强制修改设备信息
//func ForceUpdateDevices(c *gin.Context) {
//	var uav Model.Uav
//	//结构体绑定
//	if err := c.BindJSON(&uav); err != nil {
//		fmt.Println("强制修改设备信息数据绑定失败：", err.Error())
//		c.JSON(400, gin.H{"msg": "参数格式错误"})
//		return
//	}
//	Service.UpdateDevice(uav)
//	_, device := Service.GetUavByUid(uav.Uid)
//
//	c.JSON(200, &device)
//}

// GetDevices 获取所有设备(前端指定状态和类型)
//func GetDevices(c *gin.Context) {
//	var uavs Model.Uav
//	//结构体绑定
//	if err := c.BindJSON(&uavs); err != nil {
//		log.Fatal(err.Error())
//		return
//	}
//	device := Service.GetUavByAll(uavs)
//
//	c.JSON(200, &device)
//}

//// SampleOut xlsx格式下
//func SampleOut(c *gin.Context) {
//	var code int
//
//	//获取数据信息
//	data, err := Service.GetDeviceData()
//	if err != nil {
//		code = Status.FuncFail
//		c.JSON(code, R(code, nil, "获取设备数据失败"))
//		return
//	}
//
//	//补充部门信息
//	for index, device := range data {
//		uavType, err := Service.GetTypeByName(device.Type)
//		if err != nil {
//			continue
//			//code = Status.ErrorData
//			//c.JSON(code, R(code, nil, "设备类型获取失败"))
//			//return
//		}
//
//		department, err := Service.GetDepartment(uavType.DepartmentID)
//		if err != nil {
//			continue
//			//code = Status.FuncFail
//			//c.JSON(code, R(code, nil, "获取部门实例失败"))
//			//return
//		}
//		data[index].Department = department.DepartmentName
//	}
//	//类型泛化
//	re := utils.ToInterfaceSlice(data)
//	headers := []string{"ID", "设备名称", "设备状态", "设备部门", "设备类型", "设备序列号", "设备存放位置", "设备备注", "是否贵重"}
//
//	content := utils.ToExcel(headers, re, "设备")
//	utils.ResponseXls(c, content, "DeviceInfo")
////}
