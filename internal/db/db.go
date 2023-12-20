package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
    // Load .env file
    println("connect to db =================1")
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Build the connection string for PostgreSQL
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    // Establish a connection to the database
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Ping the database to verify the connection (GORM does this implicitly)
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get database connection: %v", err)
    }
    sqlDB.SetMaxIdleConns(2)
    sqlDB.SetMaxOpenConns(5)
    sqlDB.SetConnMaxLifetime(time.Hour)
    err = sqlDB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
}
