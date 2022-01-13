package migrations

import (
	"github.com/LuanFreitasRibeiro/webapi-mvc-golang.git/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.User{})
}
