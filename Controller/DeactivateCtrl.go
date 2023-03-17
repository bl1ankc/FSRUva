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
