package models

import (
	db "src/config"

	"gorm.io/gorm"
)

func Init() *gorm.DB {
	database := db.Init()
	return database
}
