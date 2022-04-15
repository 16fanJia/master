package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"mime/multipart"
	"path"
	"stream-video/util"
)

var OssClient *oss.Client

func InitOss() {
	var err error
	OssClient, err = oss.New(
		viper.GetString("oss.endPoint"),
		viper.GetString("oss.accessID"),
		viper.GetString("oss.accessKey"),
	)
	if err != nil {
		util.Log.Panic("oss connection failed err:" + err.Error())
		return
	}
}

// UploadFile oss上传文件
func UploadFile(file *multipart.FileHeader) error {
	ossBucket, err := OssClient.Bucket(viper.GetString("OSS.bucket"))
	if err != nil {
		return err
	}
	//上传文件流

	fd, errOpen := file.Open()
	if errOpen != nil {
		return errOpen
	}

	prefix := "video/"
	ossName := path.Join(prefix, file.Filename)

	if err := ossBucket.PutObject(ossName, fd); err != nil {
		return err
	}

	return nil
}

func DeleteFile(fileName string) error {
	ossBucket, err := OssClient.Bucket(viper.GetString("OSS.bucket"))
	if err != nil {
		return err
	}
	prefix := "video/"
	ossName := path.Join(prefix, fileName)

	//判断oss 是否存在
	ok, _ := ossBucket.IsObjectExist(ossName)
	if ok {
		if err := ossBucket.DeleteObject(ossName); err != nil {
			return err
		}
	}
	return nil
}
