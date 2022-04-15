package model

import (
	"github.com/jinzhu/gorm"
)

// Video 视频表
type Video struct {
	gorm.Model
	UserId   int    `gorm:"type:int(10)"`                 //user id  谁上传的
	UserName string `gorm:"type:varchar(50)"`             //上传者名称
	Title    string `gorm:"type:varchar(50)"`             //视频标题
	Desc     string `gorm:"varchar(100);default:'什么也没有'"` //视频简介
	Clicks   int    `gorm:"default:0"`                    //点击量
	Review   bool   `gorm:"default:false"`                //是否审查通过
}
