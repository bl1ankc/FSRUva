// @File: userErr.go
// @Author: Jason
// @Date: 2022/7/16

package Status

// 用户级别错误
const (
	UserExists = 900 + iota // 注册时用户已存在（以账号为准）
	UserNotExists
	InvalidAccount // 账号密码错误
	InvalidToken
	InvalidParams
	InvalidAdmin // 操作权限不足
	RecordNotFound
)
