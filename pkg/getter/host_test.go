package getter

import (
	"contact-api/pkg/storage"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// Expects OWD to be there as our db user currently doesnt have delete/update/insert privileges
var envPath string = "../../.env"

func TestRead(t *testing.T) {
	// Setup
	godotenv.Load(envPath)
	db := storage.Connection()

	host := &Host{
		Url: "https://owd.noahvarghese.me",
	}

	err := host.Read(db)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if host.Name == "" {
		t.Error("Test Host Read Failed: No data returned")
	}
}
