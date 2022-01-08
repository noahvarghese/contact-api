package models

import (
	db "contact-api/config"

	"gorm.io/gorm"
)

func Init() *gorm.DB {
	database := db.Init()
	return database
}
