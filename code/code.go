package code

type Code int

const (
	InternalServerError Code = 500

	Ok                         Code = 0
	ErrorNotAuthUser           Code = 4001
	RequestError               Code = 4002
	UserAlreadyExists          Code = 4003
	DatabaseQueryError         Code = 4004
	UserNameAlreadyExists      Code = 4005
	InvalidParam               Code = 4006
	InvalidEmail               Code = 4007
	UserPasswordError          Code = 4008
	ReleaseTokenError          Code = 4009
	AddTokenToRedisError       Code = 4010
	UpdateTokenExpirationError Code = 4011
	AddCacheToRidesError       Code = 4012
)

var codeMsg = map[Code]string{
	Ok:                         "OK",
	ErrorNotAuthUser:           "用户验证失败",
	RequestError:               "请求错误",
	UserAlreadyExists:          "用户已存在",
	DatabaseQueryError:         "数据库查询错误",
	InternalServerError:        "服务器错误",
	UserNameAlreadyExists:      "用户名已存在",
	InvalidParam:               "请求参数无效",
	InvalidEmail:               "此邮箱未注册",
	UserPasswordError:          "用户密码错误",
	ReleaseTokenError:          "发放token失败",
	AddTokenToRedisError:       "缓存token异常",
	UpdateTokenExpirationError: "更新token过期时间异常",
	AddCacheToRidesError:       "缓存用户信息异常",
}

func (c Code) GetMessage() string {
	return codeMsg[c]
}
