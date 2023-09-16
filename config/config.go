package config

import (
	"log"

	viper "github.com/spf13/viper"
)

func initConfig(fileName string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(fileName)
	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Erro while parsing configuration file", err)
	}

	return config
}
