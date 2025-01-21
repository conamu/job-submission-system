package config

import "github.com/spf13/viper"

func Init() {
	viper.SetConfigFile("config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
