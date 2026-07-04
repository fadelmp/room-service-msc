package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("❌ Cannot read config.yaml: %v", err)
	}
}

func Get(key string) interface{} {
	return viper.Get(key)
}
