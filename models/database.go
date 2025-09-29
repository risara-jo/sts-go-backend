package models

import (
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(databaseURL string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return err
	}

	// Test the connection and configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	// Set connection pool settings for Supabase Pooler
	// Pooler connections have different limits than direct connections
	if strings.Contains(databaseURL, "pooler.supabase.com") {
		// Settings optimized for Supabase pooler
		sqlDB.SetMaxOpenConns(10)                 // Lower for pooler
		sqlDB.SetMaxIdleConns(5)                  // Lower idle connections
		sqlDB.SetConnMaxLifetime(time.Minute * 2) // Shorter lifetime for pooler
	} else {
		// Settings for direct connection
		sqlDB.SetMaxOpenConns(25)                 // Higher for direct connection
		sqlDB.SetMaxIdleConns(10)                 // More idle connections
		sqlDB.SetConnMaxLifetime(time.Minute * 5) // Longer lifetime
	}

	return sqlDB.Ping()
}
