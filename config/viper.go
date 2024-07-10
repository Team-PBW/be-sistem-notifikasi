package config

import (
	"log"

	"github.com/spf13/viper"
)

func ViperGetEnv(key string) string {
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	// change with your own env file directory
	viper.AddConfigPath("D:/DEVELOPMENT/golang/sistem reminder si-be/be-sistem-reminder")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Can't find the env file")
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion for key '%s'", key)
	}

	return value
}
