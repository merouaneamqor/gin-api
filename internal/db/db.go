package db

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

// InitDB initializes and returns a database connection
func InitDB() *gorm.DB {
    // Load .env file
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
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Ping the database to verify the connection
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Failed to get database connection: %v", err)
    }

    err = sqlDB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    return db
}

func init() {
    // Initialize the global DB variable
    DB = InitDB()
}
