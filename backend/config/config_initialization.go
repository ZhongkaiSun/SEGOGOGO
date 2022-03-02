package config

import (
	_ "fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig(mode string) {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")

	if mode == "main" {
		viper.AddConfigPath(workDir + "/config")
	} else if mode == "test" {
		viper.AddConfigPath("../config")
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
