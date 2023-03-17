// @File: systemErr.go
// @Author: b1ankc
// @Date: 2022/7/16

package Status

// 系统级别的错误
const (
	OK             = 200
	FailToBindJson = 400 + iota // 绑定json表单失败！
	FailToSave                  //存储失败
	JWTErr                      //jwt验证失败
	FailToGetQuery              //获取Query参数失败
	FuncFail                    //函数异常
	ErrorData                   //异常数据
	MidError                    //中间件异常
	UploadFail                  //提交错误
	OBSErr                      // OBS异常
)
