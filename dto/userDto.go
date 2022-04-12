package dto

// LoginParam 用户登录参数
type LoginParam struct {
	Email    string `form:"email" json:"email" binding:"required,printascii"`                    // 用户邮箱，不能包含特殊字符，用于登录
	Password string `form:"password" json:"password" binding:"required,printascii,min=6,max=16"` // 用户密码，最短长度为6，最大长度为16，不能包含特殊字符
}

// RegisterParam 用户注册参数
type RegisterParam struct {
	Name     string `form:"name" json:"name" binding:"required,printascii,min=3,max=16"`
	Email    string `form:"email" json:"email" binding:"required,printascii"`
	Password string `form:"password" json:"password" binding:"required,printascii,min=6,max=16"`
	Gender   int    `form:"column:gender;default:0" json:"gender"` //性别
	//Code     string
}
