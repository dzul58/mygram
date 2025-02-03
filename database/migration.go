package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	// Baca semua file .up.sql dari folder migrations
	migrationFiles, err := filepath.Glob("database/migrations/*.up.sql")
	if err != nil {
		log.Fatal("Failed to read migration files:", err)
	}

	// Jalankan setiap file migrasi
	for _, file := range migrationFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("Failed to read migration file:", err)
		}

		// Eksekusi SQL
		if err := db.Exec(string(content)).Error; err != nil {
			log.Fatal("Failed to execute migration:", err)
		}

		fmt.Printf("Migration %s executed successfully\n", file)
	}
}
