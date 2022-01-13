package getter

import (
	"contact-api/pkg/storage"
	"testing"

	"github.com/joho/godotenv"
)

// Expects OWD to be there as our db user currently doesnt have delete/update/insert privileges
var envPath string = "../../.env"

func TestRead(t *testing.T) {
	godotenv.Load(envPath)

	db, err := storage.Connection()

	if err != nil {
		t.Error(err)
	}

	host := &Host{
		Url: "owd.noahvarghese.me",
	}

	err = host.Read(db)

	if err != nil {
		t.Error(err)
	}

	if host.Name == "" {
		t.Error("Test Host Read Failed: No data returned")
	}
}
