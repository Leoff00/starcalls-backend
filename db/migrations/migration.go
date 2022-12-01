package migrations

import (
	"log"

	"github.com/starcalls-backend/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(&models.Questions{})

	if err != nil {
		log.Default().Fatal("Failed to run migration, cause -> ", err.Error())
	}
}
