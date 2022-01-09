package models

import (
	"fmt"
	"os"
	db "src/config"
	"testing"

	"github.com/joho/godotenv"
)

// Expects OWD to be there as our db user currently doesnt have delete/update/insert privileges
var envPath string = "../.env"

func TestRead(t *testing.T) {
	// Setup
	godotenv.Load(envPath)
	db := db.Init()

	var host Host
	err := host.Read(db, "https://owd.noahvarghese.me")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if host.Name != "Oakville Windows & Doors" {
		t.Error("Wrong name " + host.Name)
	}
	if !host.Has_images {
		t.Error("No images found when they are set")
	}
}
