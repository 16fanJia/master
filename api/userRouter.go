package api

import (
	"github.com/gin-gonic/gin"
	"stream-video/controller/user"
)

func RegisterUserRouter(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		//用户注册路由
		userGroup.POST("/register", user.RegisterUser)
		//用户登陆
		userGroup.POST("/login", user.LoginUser)
	}
}
