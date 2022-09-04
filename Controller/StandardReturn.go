package Controller

// R 响应数据规范：{code, data, msg}
func R(code int, data any, msg string) map[string]any {
	return map[string]any{
		"code": code,
		"data": data,
		"msg":  msg,
	}
}
