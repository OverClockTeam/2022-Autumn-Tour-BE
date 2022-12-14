package errmsg

const(
	SUCCEED = 100
	ERROR   = 200

	//code = 1000... 用户模块的错误
	ERROR_USERNAME_USED = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXSIT = 1003
	ERROR_TOKEN_EXIST = 1004
	ERROT_TOKEN_RUNTIME = 1005
	ERROR_TOKEN_WRONG = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	//code = 2000...文章模块的错误
	ERROR_TITLE_USED = 2001
	//code = 3000...分类模块错误
	ERROR_CATENAME_USED = 3001
)

var codeMsg = map[int]string{
	SUCCEED:                 "OK",
	ERROR :                  "FAIL",
	ERROR_USERNAME_USED :    "用户名已存在！",
	ERROR_PASSWORD_WRONG :   "密码错误",
	ERROR_USER_NOT_EXSIT :   "用户不存在",
	ERROR_TOKEN_EXIST :      "TOKEN不存在",
	ERROT_TOKEN_RUNTIME :    "TOKEN已过期",
	ERROR_TOKEN_WRONG :      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG : "TOKEN格式错误",
	ERROR_CATENAME_USED : "分类已存在",
}

func GetErrMsg(code int)string{
	return codeMsg[code]
}
