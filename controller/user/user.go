package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"stream-video/code"
	"stream-video/common"
	"stream-video/dbops"
	"stream-video/dto"
	"stream-video/model"
	"stream-video/response"
)

// RegisterUser 用户注册逻辑函数
func RegisterUser(c *gin.Context) {
	//获取参数
	var params dto.RegisterParam
	if err := c.ShouldBind(&params); err != nil {
		response.New(code.InvalidParam).WithError(err).Return(c)
		return
	}

	//查询新注册用户邮箱是否存在
	var user model.User
	result := dbops.DB.Table("user").Where("email = ?", params.Email).First(&user)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) { //如果数据库错误不是 记录不存在 则返回错误
			response.New(code.DatabaseQueryError).WithError(result.Error).Return(c)
			return
		}
	} else { //无错误 则证明用户已存在
		response.New(code.UserAlreadyExists).Return(c)
		return
	}

	//用户名唯一限制
	errName := dbops.DB.Table("user").Where("name = ?", params.Name).First(&user).Error
	if errName != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) { //如果数据库错误不是 记录不存在 则返回错误
			response.New(code.DatabaseQueryError).WithError(result.Error).Return(c)
			return
		}
	} else { //查询成功
		response.New(code.UserNameAlreadyExists).Return(c)
		return
	}

	//密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.New(code.InternalServerError).WithError(err).Return(c)
		return
	}
	//创建新用户
	newUser := model.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: string(hashedPassword),
		Gender:   params.Gender,
	}

	dbops.DB.Create(&newUser)
	response.New(code.Ok).Return(c)
}

// LoginUser 用户登陆
func LoginUser(c *gin.Context) {
	//参数绑定
	var params dto.LoginParam
	if err := c.ShouldBind(&params); err != nil {
		response.New(code.InvalidParam).WithError(err).Return(c)
		return
	}

	//登陆邮箱验证
	var user model.User
	result := dbops.DB.Table("user").Where("email = ?", params.Email).First(&user)
	//如果邮箱存在 则影响行数不为0
	affected := result.RowsAffected
	if affected == 0 {
		response.New(code.InvalidEmail).Return(c)
		return
	}

	//核对密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		response.New(code.UserPasswordError).WithError(err).Return(c)
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.New(code.ReleaseTokenError).WithError(err).Return(c)
		return
	}
	//判断redis中 是否含有用户token
	ok := common.GetRedisToken(user.ID)
	if !ok { //未获取到token
		//将token 放入缓存
		err := common.AddRedisToken(token, user.ID)
		if err != nil {
			response.New(code.AddTokenToRedisError).WithError(err).Return(c)
			return
		}
	}

	//将用户信息 存入缓存
	userInfo := common.NewUserInfo(user.ID)
	if err := userInfo.AddUserInfoToRedis(user); err != nil {
		response.New(code.AddCacheToRidesError).WithError(err).Return(c)
		return
	}

	response.New(code.Ok).WithToken(token).Return(c)
}
