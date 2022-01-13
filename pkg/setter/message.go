package setter

import (
	"errors"

	"gorm.io/gorm"
)

var db *gorm.DB

type Message struct {
	gorm.Model
	ID       uint
	Contents string
	Sent     bool
	Host_id  uint
}

func NewMessage(d *gorm.DB, c string, h uint) (*Message, error) {
	m := &Message{Contents: c, Host_id: h, Sent: false}
	db = d

	if db == nil {
		return nil, errors.New("database connection not instantiated")
	}

	err := m.Create()

	return m, err
}

func (m *Message) Create() error {
	if db == nil {
		return errors.New("database connection not instantiated")
	}

	tx := db.Create(m)
	return tx.Error
}

func (m *Message) SetSent() error {
	if db == nil {
		return errors.New("database connection not instantiated")
	}

	m.Sent = true

	tx := db.Update("sent", m)

	return tx.Error
}
