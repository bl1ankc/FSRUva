// @File: userErr.go
// @Author: b1ankc
// @Date: 2022/7/16

package Status

// 用户级别错误
const (
	UserExists         = 900 + iota // 注册时用户已存在（以账号为准）
	UserNotExists                   //用户不存在
	InvalidAccount                  // 账号密码错误
	InvalidToken                    //无效token
	InvalidParams                   //无效参数
	InvalidAdmin                    // 操作权限不足
	RecordNotFound                  //记录查询失败
	UserAuthentication              //用户身份验证错误(非本人)
)
