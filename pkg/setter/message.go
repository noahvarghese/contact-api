package setter

import (
	"contact-api/pkg/getter"
	"encoding/json"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Original string
	Contents string
	Host_id  int
}

func (m *Message) Create(db *gorm.DB, fieldsInMessage **getter.Field) error {
	fields := make(map[string]string)

	json.Unmarshal([]byte(m.Contents), &fields)

	tx := db.Create(m)
	return tx.Error
}
