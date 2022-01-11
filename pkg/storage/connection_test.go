package storage

import (
	"testing"
)

func TestLoadConfigFromEnv(t *testing.T) {
	config := loadConfigFromEnv()
	// t.Log(config, "\n")

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
	db := Connection()

	if db == nil {
		t.Error("Invalid connection string")
	}
}
