package main

import (
	"fmt"
	"kasir-api/internal"
	"kasir-api/internal/adapter/outbound/persistence/entity"
	"kasir-api/internal/config"
	"log"
	"net/http"

	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Failed to load config: %v, using defaults", err)
		// Assuming defaults are handled inside LoadConfig or we proceed with empty values checking
		if cfg == nil {
			cfg = &config.Config{Port: "8080", Storage: "memory"}
		}
	}

	var db *gorm.DB
	if cfg.Storage == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)
		db, err = gorm.Open(driver.Open(dsn), &gorm.Config{
			PrepareStmt: false, // Required for Supabase Transaction Mode (PgBouncer)
		})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Auto Migrate
		// Explicitly auto migrate the entities.
		// Note: Ideally migration should be separate, but for this challenge we put it here.
		if err := db.AutoMigrate(&entity.Produk{}, &entity.Category{}); err != nil {
			log.Printf("Failed to auto migrate: %v", err)
		}
	}

	handler := internal.NewApp(cfg, db)

	fmt.Printf("Server running di localhost:%s\n", cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, handler)
	if err != nil {
		fmt.Printf("Failed running server: %v\n", err)
	}
}
