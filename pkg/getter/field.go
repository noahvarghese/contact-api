package getter

import (
	"time"

	"gorm.io/gorm"
)

type Field struct {
	gorm.Model
	ID        int
	Name      string
	Required  bool
	Host_id   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (f *Field) Read(db *gorm.DB, host_id int) error {
	tx := db.First(&f, "host_id = ?", host_id)
	return tx.Error
}

func (f *Field) GetAll(db *gorm.DB, host_id uint) ([]Field, error) {
	var fields []Field

	tx := db.Where("host_id = ?", host_id).Find(&fields)

	return fields, tx.Error
}
