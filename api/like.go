package api

import (
	"github.com/gin-gonic/gin"
	"stream-video/controller/like"
	"stream-video/middleware"
)

// RegisterLikeRouter 注册点赞路由
func RegisterLikeRouter(r *gin.RouterGroup) {
	r.Use(middleware.AuthMiddleware())
	r.POST("/like", like.VideoLike) //点赞路由
}
