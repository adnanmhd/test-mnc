package util

import (
	"os"
)

type APPConfig struct {
	AppName  string
	Port     string
	Database DB
}

type DB struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DBName       string
	DBSchemaName string
}

func InitConfig() APPConfig {
	appConfig := APPConfig{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
		Database: DB{
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			Username:     os.Getenv("DB_USERNAME"),
			Password:     os.Getenv("DB_PASSWORD"),
			DBName:       os.Getenv("DB_NAME"),
			DBSchemaName: os.Getenv("DB_SCHEMA_NAME"),
		},
	}

	return appConfig
}
