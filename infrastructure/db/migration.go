package db

import (
	"log"

	"room-service-msc/domain"

	"gorm.io/gorm"
)

// RunMigrations runs all GORM migrations
func RunMigrations(db *gorm.DB) {
	log.Println("🚀 Running database migrations...")

	err := db.AutoMigrate(&domain.Hotel{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migration complete")
}
