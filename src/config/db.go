package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseConfig struct {
	DB_NAME string
	DB_URL  string
	DB_USER string
	DB_PWD  string
	DB_PORT string
}

func LoadConfigFromEnv() databaseConfig {
	config := databaseConfig{
		DB_NAME: os.Getenv("DB_NAME"),
		DB_PWD:  os.Getenv("DB_PWD"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_URL:  os.Getenv("DB_URL"),
		DB_USER: os.Getenv("DB_USER"),
	}

	return config
}

func GenerateDSN(config databaseConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		config.DB_USER,
		config.DB_PWD,
		config.DB_URL,
		config.DB_PORT,
		config.DB_NAME,
	)
}

func Init() (*gorm.DB, error) {
	config := LoadConfigFromEnv()
	dsn := GenerateDSN(config)

	db, err := gorm.Open(mysql.Open(dsn))
	return db, err
}
