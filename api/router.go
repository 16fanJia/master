package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		GetUserRouter(v1) //获取用户路由
	}

}
