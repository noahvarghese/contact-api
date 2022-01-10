package storage

import (
	"testing"

	"github.com/joho/godotenv"
)

var envPath string = "../.env"

func TestLoadConfigFromEnv(t *testing.T) {
	godotenv.Load(envPath)
	config := loadConfigFromEnv()

	t.Log(config)

	if config.DB_NAME == "" {
		t.Error("DB_NAME NOT SET")
	}

	if config.DB_NAME == "" {
		t.Error("DB_PWD NOT SET")
	}

	if config.DB_NAME == "" {
		t.Error("DB_PORT NOT SET")
	}

	if config.DB_NAME == "" {
		t.Error("DB_URL NOT SET")
	}

	if config.DB_NAME == "" {
		t.Error("DB_USER NOT SET")
	}
}

func TestInit(t *testing.T) {
	godotenv.Load(envPath)

	db := Connection()

	if db == nil {
		t.Error("Invalid connection string")
	}
}
