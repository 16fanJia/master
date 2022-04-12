package config

import "github.com/spf13/viper"

// InitConfig 初始化配置文件
func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}
