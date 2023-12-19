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

func init() {
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
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Ping the database to verify the connection (GORM does this implicitly)
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get database connection: %v", err)
    }

    err = sqlDB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
}
func InitDB() error {
    dsn := "host=" + os.Getenv("DB_HOST") + 
           " user=" + os.Getenv("DB_USER") +
           " password=" + os.Getenv("DB_PASS") +
           " dbname=" + os.Getenv("DB_NAME") +
           " port=" + os.Getenv("DB_PORT") +
           " sslmode=disable"

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    sqlDB, err := DB.DB()
    if err != nil {
        return err
    }

    if err := sqlDB.Ping(); err != nil {
        return err
    }

    return nil
}