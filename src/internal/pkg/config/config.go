package config

import "github.com/spf13/viper"

func Init() {
	viper.AddConfigPath("config/")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
