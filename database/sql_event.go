package database

import (
	// "database/sql"
	// "os"
	// "strconv"
	// "log"
	"fmt"

	// "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"golang.org/x/e-calender/config"
	"gorm.io/gorm"
	"errors"
)

// func viperEnvVar(key string) string {
// 	viper.SetConfigType("env")
// 	viper.SetConfigName("app")
// 	viper.AddConfigPath("D:/DEVELOPMENT/golang/src/user_service")

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Println("Can't find the env file")
// 	}

// 	value, ok := viper.Get(key).(string)
// 	if !ok {
// 		log.Fatalf("Invalid type assertion for key '%s'", key)
// 	}

// 	return value
// }

var (
	DB_USERNAME = config.ViperGetEnv("DB_USERNAME")
	DB_PASS     = config.ViperGetEnv("DB_PASS")
	DB_HOST     = config.ViperGetEnv("DB_HOST")
	// DB_PORT = config.ViperGetEnv("DB_PORT")
	DB_SCHEMA = config.ViperGetEnv("DB_SCHEMA")
	// DB_PORT = 5432 // for postgresql
	DB_PORT = 8111 // for mysql
)

func NewGorm() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USERNAME, DB_PASS, DB_HOST, DB_PORT, DB_SCHEMA)

	// dbPort, err := strconv.ParseUint(DB_PORT, 10, 32)
	// if err != nil {
	// 	// Handle the error, e.g., log it and return an error
	// 	return nil, err
	// }

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USERNAME, DB_PASS, DB_SCHEMA, dbPort)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	return nil, errors.New("database not found")
	// }
	// return db, nil

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USERNAME, DB_PASS, DB_HOST, DB_PORT, DB_SCHEMA)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, errors.New("database not found")
    }
    return db, nil
}
