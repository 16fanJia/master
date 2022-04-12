package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// User 用户表
type User struct {
	gorm.Model
	Name     string    `gorm:"column:name;type:varchar(20);not null"`
	Email    string    `gorm:"column:email;type:varchar(20);not null;index"`
	Password string    `gorm:"column:password;not null"`
	Gender   int       `gorm:"column:gender;default:0"`
	Birthday time.Time `gorm:"column:birthday"`
}
