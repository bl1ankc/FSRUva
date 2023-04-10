package Const

const (
	//OSS云端

	Endpoint   = "https://obs.cn-south-1.myhuaweicloud.com"
	AK         = "UIWD3I0XCETSEQMIH2RR"
	SK         = "yJPqICTge7BPzvRbOoLzKAvzkaztb5E82U6f2Pnw"
	RoleArn    = "acs:ram::1761650696847549:role/tmp"
	BucketName = "fsrlab-rfid"
	RegionID   = "cn-shenzhen"

	//OSS日志
	//
	//LogFullPath  = "./root/gopath/src/FSRUva/logs/OBS-SDK.log"
	//MaxLogSize   = 1024 * 1024 * 10
	//Backups      = 10
	//Level        = obs.LEVEL_INFO
	//LogToConsole = false
)

const (
	//微信小程序数据

	APPID     = ""
	APPSECRET = ""
)

var (
	//微信模板ID

	WXMESSAGE = map[string]string{
		"RemindUserReturnUav": "",
		"RemindScheduleOK":    "",
		"RemindCheckOK":       "",
		"RemindAdminCheck":    "",
	}
)
