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
	UserNotLoggedIn            Code = 4013
	TokenFailure               Code = 4014
	GetUserInfoFromRedisError  Code = 4015
	GetFileError               Code = 4016
	FileFormatError            Code = 4017
	FileUploadError            Code = 4018
	DataCreateError            Code = 4019
	DataDoesNotExist           Code = 4020
	UserIdNotExist             Code = 4021
	ErrAlreadyLike             Code = 4022
	LikeError                  Code = 4023
	RedisError                 Code = 4024
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
	UserNotLoggedIn:            "用户未登陆",
	TokenFailure:               "token失效",
	GetUserInfoFromRedisError:  "获取用户信息失败",
	GetFileError:               "获取文件失败",
	FileFormatError:            "文件格式错误",
	FileUploadError:            "文件上传失败",
	DataCreateError:            "数据创建入库失败",
	DataDoesNotExist:           "数据不存在",
	UserIdNotExist:             "用户ID不存在",
	ErrAlreadyLike:             "重复点赞",
	LikeError:                  "点赞失败",
	RedisError:                 "redis 错误",
}

func (c Code) GetMessage() string {
	return codeMsg[c]
}
