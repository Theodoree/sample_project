package errs

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_PASSWORD:                 "密码错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_PASSWORD_OR_USER:         "账户或密码错误",
	ERROR_MYSQL_INSERT_FAIL:        "插入失败",
	ERROR_MYSQL_SELECT_FAIL:        "查询失败",
	ERROR_MYSQL_UPDATE_FAIL:        "更新失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
