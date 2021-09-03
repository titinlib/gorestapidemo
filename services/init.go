package services

import (
	"restapidemo/database"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = database.GetDB()
}
