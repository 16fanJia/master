package api

import (
	"github.com/gin-gonic/gin"
	"stream-video/controller/video"
	"stream-video/middleware"
)

// RegisterVideoRouter 注册video 视频路由
func RegisterVideoRouter(r *gin.RouterGroup) {
	videoRouter := r.Group("/video", middleware.AuthMiddleware())
	{
		videoRouter.POST("/uploadVideo", video.UploadVideo)                 //上传视频
		videoRouter.GET("/getVideo/:id", middleware.Hits(), video.GetVideo) //观看视频
		videoRouter.GET("/videoList")                                       //获取全部视频列表
		videoRouter.GET("/videoFromUser")                                   //获取用户的视屏列表
		videoRouter.DELETE("/deleteVideo")                                  //用户删除自己已上传的视屏
		videoRouter.PUT("/updateVideo")                                     //更新视频基本信息
	}

}
