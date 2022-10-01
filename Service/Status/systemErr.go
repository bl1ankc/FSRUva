// @File: systemErr.go
// @Author: Jason
// @Date: 2022/7/16

package Status

// 系统级别的错误
const (
	OK             = 200
	FailToBindJson = 400 + iota // 绑定json表单失败！
	FailToSave
	JWTErr
	_
	FuncFail
	ErrorData //异常数据
	MidError  //中间件异常
	UploadFail
	OBSErr // OBS异常
)
