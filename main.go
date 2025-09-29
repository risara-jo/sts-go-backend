package main

import (
	"log"
	"net/http"

	"smart-transit-backend/config"
	"smart-transit-backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database if URL is provided
	if cfg.DatabaseURL != "" {
		if err := models.InitDB(cfg.DatabaseURL); err != nil {
			log.Printf("Warning: Failed to connect to database: %v", err)
		} else {
			log.Println("✅ Database connection established successfully")
		}
	}

	// Initialize Gin router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "Smart Transit System Backend",
			"version": "1.0.0",
		})
	})

	// Simple health check for containers
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Readiness probe with database check
	router.GET("/ready", func(c *gin.Context) {
		if models.DB != nil {
			if sqlDB, err := models.DB.DB(); err == nil {
				if err := sqlDB.Ping(); err == nil {
					c.JSON(http.StatusOK, gin.H{
						"status":   "ready",
						"database": "connected",
					})
					return
				}
			}
		}
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":   "not ready",
			"database": "disconnected",
		})
	})

	// Database health check
	router.GET("/health/db", func(c *gin.Context) {
		if models.DB == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "unhealthy",
				"service": "database",
				"error":   "database connection not initialized",
			})
			return
		}

		sqlDB, err := models.DB.DB()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "unhealthy",
				"service": "database",
				"error":   "failed to get database instance",
			})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "unhealthy",
				"service": "database",
				"error":   "database ping failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "database",
		})
	})

	// Start server
	log.Printf("🚀 Smart Transit System Backend starting on port %s", cfg.Port)
	log.Printf("📊 Health check: http://localhost:%s/health", cfg.Port)
	log.Printf("🔍 Database health: http://localhost:%s/health/db", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
