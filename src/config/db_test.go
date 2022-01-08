package config

import (
	"os"
	"testing"
)

func TestLoadConfigFromEnv(t *testing.T) {

	os.Setenv("DB_NAME", "NAME")
	os.Setenv("DB_PWD", "PWD")
	os.Setenv("DB_PORT", "PORT")
	os.Setenv("DB_URL", "URL")
	os.Setenv("DB_USER", "USER")

	config := LoadConfigFromEnv()

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
	db, err := Init()

	if err != nil {
		t.Error("Error creating database")
	}

	if db == nil {
		t.Error("Invalid config")
	}
}
