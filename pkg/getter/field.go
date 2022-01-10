package getter

import (
	"gorm.io/gorm"
)

type Field struct {
	gorm.Model
	Id      int
	Name    string
	Type    string
	Host_id int
}

func (f *Field) Read(db *gorm.DB, host_id int) error {
	tx := db.Unscoped().First(&f, "host_id = ?", host_id)
	return tx.Error
}
