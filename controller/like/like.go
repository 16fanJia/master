package like

import (
	"github.com/gin-gonic/gin"
	"stream-video/allParams"
	"stream-video/code"
	"stream-video/response"
	"stream-video/util"
)

// VideoLike 点赞逻辑
func VideoLike(c *gin.Context) {
	//绑定参数
	var params allParams.ParamLikeData
	if err := c.ShouldBind(&params); err != nil {
		response.New(code.InvalidParam).WithError(err).Return(c)
		return
	}

	//获取用户ID
	id, exists := c.Get("userId")
	if !exists { //userId 不存在 错误
		util.Log.Info("UserID 不存在！！")
		response.New(code.UserIdNotExist).Return(c)
		return
	}

	//判断用户是否对此视频点过赞
	if ok := IsVideoLike(id.(int), params.VideoId); ok { //已经点过赞
		//走取消赞的逻辑
		if ok := RemVideoLike(id.(int), params.VideoId); !ok {
			response.New(code.LikeError).Return(c)
			return
		}
	}

	//为视频点赞
	if ok := DoVideoLike(id.(int), params.VideoId); !ok { //点赞失败
		response.New(code.LikeError).Return(c)
		return
	}

	response.New(code.Ok).Return(c)
}
