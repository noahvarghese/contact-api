package getter

import (
	"time"

	"gorm.io/gorm"
)

type Template struct {
	gorm.Model
	ID        uint
	Template  string
	Host_id   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (t *Template) Read(db *gorm.DB, id uint) error {
	tx := db.First(&t, "host_id = ?", id)
	return tx.Error
}
