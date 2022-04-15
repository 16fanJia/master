package main

import (
	"github.com/gin-gonic/gin"
	"stream-video/api"
	"stream-video/config"
	"stream-video/dbops"
	"stream-video/oss"
	"stream-video/util"
)

//
func main() {
	util.InitLog()      //初始化日志文件
	config.InitConfig() //初始还配置文件
	dbops.InitDB()      // 初始化gorm db
	dbops.InitRedis()   //初始化redis
	oss.InitOss()       //初始化Oss

	r := gin.Default()

	api.RegisterApi(r)

	if err := r.Run(":8080"); err != nil {
		return
	}
}
