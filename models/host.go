package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	Id         uint
	Name       string
	Has_images bool
	Url        string
}

func (host *Host) Read(db *gorm.DB, url string) error {
	tx := db.Unscoped().First(&host, "url = ?", url)
	return tx.Error
}
