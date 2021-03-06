package storage

import (
	"testing"

	"github.com/joho/godotenv"
)

const EnvPath = "../../.env"

func TestLoadConfigFromEnv(t *testing.T) {
	godotenv.Load(EnvPath)
	config := loadConfigFromEnv()

	if config.DB_NAME == "" {
		t.Error("DATABASE NAME NOT SET")
	}

	if config.DB_PWD == "" {
		t.Error("DATABASE PASSWORD NOT SET")
	}

	if config.DB_PORT == "" {
		t.Error("DATABASE PORT NOT SET")
	}

	if config.DB_URL == "" {
		t.Error("DATABASE URL NOT SET")
	}

	if config.DB_USER == "" {
		t.Error("DATABASE USER NOT SET")
	}
}

func TestInit(t *testing.T) {
	db, err := Connection()

	if db == nil {
		t.Error("Invalid connection string")
	}

	if err != nil {
		t.Error(err)
	}
}
