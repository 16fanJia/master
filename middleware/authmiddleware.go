package middleware

import (
	"github.com/gin-gonic/gin"
	"stream-video/code"
	"stream-video/common"
	"stream-video/response"
	"stream-video/util"
	"strings"
)

// AuthMiddleware 判断用户是否登陆 中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头 Authorization 中读取 Token
		AuthString := c.GetHeader("Authorization")
		//如果token获取失败 判断抬头是否为 Bearer
		if AuthString == "" || !strings.HasPrefix(AuthString, "Bearer") {
			response.New(code.UserNotLoggedIn).Return(c)
			c.Abort()
			return
		}

		//token 获取成功 获取后面token
		parts := strings.SplitN(AuthString, " ", 2)
		tokenString := parts[1]

		//解析token
		token, claims, err := common.ParseToken(tokenString)
		//解析token错误 或者token 无效
		if err != nil || !token.Valid {
			util.Log.Info("token 失效 err:" + err.Error())
			response.New(code.TokenFailure).Return(c)
			c.Abort()
			return
		}

		//从缓存中获取用户信息
		name, err := common.NewUserInfo(claims.UserId).GetFromField("name")
		if err != nil {
			response.New(code.GetUserInfoFromRedisError).Return(c)
			c.Abort()
			return
		}
		c.Set("name", name)
		c.Set("userId", claims.UserId)

		c.Next()
	}
}
