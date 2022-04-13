package util

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log *logrus.Logger

// InitLog 初始化日志
func InitLog() {
	Log = logrus.New()
	//定位到某个文件的日志
	Log.SetReportCaller(true)

	//指定每一天日志记录
	filename := "./log/" + time.Now().Format("2006-01-02") + "server.log" //也可将name作为参数传进来

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //如果文件不存在 则创建相应文件
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("创建日志文件失败：" + err.Error())
	}
}
