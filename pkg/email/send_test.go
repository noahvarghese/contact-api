package email

import (
	"contact-api/pkg/getter"
	"errors"
	"testing"

	"github.com/joho/godotenv"
)

const EnvPath = "../../.env"
const Template = "Hello {{ .Name }}"
const Correct = "Hello Noah"

func TestBind(t *testing.T) {
	godotenv.Load(EnvPath)

	tpl := &getter.Template{
		Template: Template,
	}
	body := map[string]string{"Name": "Noah"}

	tplString, err := Bind(tpl, body)

	if err != nil {
		t.Error(err)
	}

	if tplString != Correct {
		t.Error(errors.New("incorrect templated string: " + tplString + "\nExpected: " + Correct))
	}
}

func TestSend() {

}
