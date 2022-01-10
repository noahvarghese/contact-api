package getter

import (
	"gorm.io/gorm"
)

type Template struct {
	gorm.Model
	Id int
	Template string
	Host_id int
}

func (t *Template) Read(db *gorm.DB, h *Host) error {
	tx := db.Unscoped().First(&t, "host_id = ?", h.Id)
	return tx.Error
}