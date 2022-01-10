package getter

import (
	"gorm.io/gorm"
)

type Host struct {
	gorm.Model
	Id int 
	Name string
	Has_images bool
	Url string
}

func (h *Host) Read(db *gorm.DB) error {
	tx := db.Unscoped().First(&h, "url = ?", h.Url)
	return tx.Error
}