package dbops

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"stream-video/util"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	dbname := viper.GetString("database.dbname")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		util.Log.Panic("failed to connect database err:" + err.Error())
	}
}
