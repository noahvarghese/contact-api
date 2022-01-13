package getter

import (
	"time"

	"gorm.io/gorm"
)

type Host struct {
	gorm.Model
	ID        uint
	Name      string
	Email     string
	Subject   string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (h *Host) Read(db *gorm.DB) error {
	tx := db.First(&h, "url = ?", h.Url)
	return tx.Error
}
