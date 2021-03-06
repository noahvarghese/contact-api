package email

import (
	"bytes"
	"contact-api/pkg/getter"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func Bind(t *getter.Template, b map[string]interface{}) (string, error) {
	tmpl, err := template.New("template").Parse(t.Template)

	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, b)

	mail := buffer.String()

	return mail, nil
}

func Send(m string, h *getter.Host) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	to := []string{h.Email}

	password := os.Getenv("SMTP_PWD")
	port := os.Getenv("SMTP_PORT")
	url := os.Getenv("SMTP_URL")
	user := os.Getenv("SMTP_USER")

	subject := fmt.Sprintf("Subject: %s\r\n", h.Subject)
	message := []byte(subject + mime + "\r\n" + m)

	auth := smtp.PlainAuth("", user, password, url)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", url, port), auth, user, to, message)

	return err
}
