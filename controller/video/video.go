package video

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"path/filepath"
	"strconv"
	"stream-video/allParams"
	"stream-video/code"
	"stream-video/controller/like"
	"stream-video/dbops"
	"stream-video/model"
	"stream-video/oss"
	"stream-video/response"
	"stream-video/util"
)

// UploadVideo 上传视频
func UploadVideo(c *gin.Context) {
	var params allParams.UploadVideoParams
	//绑定参数
	if err := c.ShouldBind(&params); err != nil {
		response.New(code.InvalidParam).WithError(err).Return(c)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		util.Log.Info("获取文件失败:" + err.Error())
		response.New(code.GetFileError).WithError(err).Return(c)
		return
	}

	//判断文件是否满足内置的标准格式
	allowFlag := false //标识位
	allowExits := []string{".mp4", ".flv", ".avi", ".wmv", ".mkv", ".m4v"}

	for _, v := range allowExits {
		if v == filepath.Ext(file.Filename) {
			allowFlag = true
			break
		}
	}

	if !allowFlag { //不符合标准
		response.New(code.FileFormatError).Return(c)
		return
	}

	//上传文件
	if err := oss.UploadFile(file); err != nil {
		util.Log.Info("上传文件失败: " + err.Error())
		response.New(code.FileUploadError).WithError(err).Return(c)
		return
	}

	//数据入库
	videoData := model.Video{
		UserId:   params.UserId,
		UserName: params.UserName,
		Title:    params.Title,
		Desc:     params.Dsc,
		Clicks:   0,
	}

	if err := dbops.DB.Table("video").Create(&videoData).Error; err != nil {
		util.Log.Error("数据入库失败：err:" + err.Error())
		response.New(code.DataCreateError).WithError(err).Return(c)
		return
	}

	ctx := context.Background()
	//在redis中添加video数据
	{
		_, err := dbops.RDB.ZAdd(ctx, like.KeyLikeNumberZSet, &redis.Z{
			Score:  0,
			Member: strconv.Itoa(int(videoData.ID)),
		}).Result()
		if err != nil {
			util.Log.Error("添加数据到redis 失败 err: " + err.Error())
			response.New(code.DataCreateError).WithError(err).Return(c)
			return
		}
	}

	response.New(code.Ok).Return(c)
}
