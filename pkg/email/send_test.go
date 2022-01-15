package email

import (
	"contact-api/pkg/getter"
	"errors"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const EnvPath = "../../.env"
const Template = "This is a {{ .Test }}"
const Correct = "This is a Test"

func TestBind(t *testing.T) {
	godotenv.Load(EnvPath)

	tpl := &getter.Template{
		Template: Template,
	}
	body := map[string]interface{}{"Test": "Test"}

	tplString, err := Bind(tpl, body)

	if err != nil {
		t.Error(err)
	}

	if tplString != Correct {
		t.Error(errors.New("incorrect templated string: " + tplString + "\nExpected: " + Correct))
	}
}

func TestSend(t *testing.T) {
	godotenv.Load(EnvPath)

	tpl := &getter.Template{
		Template: Template,
	}
	body := map[string]interface{}{"Test": "Test"}

	tplString, _ := Bind(tpl, body)

	Send(tplString, &getter.Host{Email: os.Getenv("SMTP_USER"), Subject: "TEST"})
}
