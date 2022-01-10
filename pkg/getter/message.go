package getter

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Id int
	Contents string
	Sent bool
	Host_id int
}

func (m *Message) Read(db *gorm.DB) error {
	tx := db.Unscoped().First(&m, "id = ?", m.Id)
	return tx.Error
}