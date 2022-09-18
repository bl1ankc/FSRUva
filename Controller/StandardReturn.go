package Controller

// R 响应数据规范：{code, data, msg}
func R(code int, data interface{}, msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"data": data,
		"msg":  msg,
	}
}
