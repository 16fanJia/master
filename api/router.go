package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		RegisterUserRouter(v1)  //注册用户路由
		RegisterVideoRouter(v1) //注册video路由
		RegisterLikeRouter(v1)  //注册点赞路由
	}

}
