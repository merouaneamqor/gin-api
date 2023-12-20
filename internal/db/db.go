package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDBConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set default values for connection pool parameters
	setDefaultEnv("DB_MaxOpenConns", "5")
	setDefaultEnv("DB_MaxIdleConns", "2")
}

// setDefaultEnv sets the specified environment variable to the default value if it is not set
func setDefaultEnv(key, defaultValue string) {
	if os.Getenv(key) == "" {
		os.Setenv(key, defaultValue)
	}
}

// buildDSN builds the connection string for PostgreSQL based on environment variables
func buildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
}

// InitDB initializes and returns a database connection
func InitDB() *gorm.DB {
	initDBConfig()

	// Build the connection string for PostgreSQL
	dsn := buildDSN()

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

	// Configure connection pool parameters
	sqlDB.SetMaxIdleConns(getIntEnv("DB_MaxIdleConns"))
	sqlDB.SetMaxOpenConns(getIntEnv("DB_MaxOpenConns"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping the database again to ensure the changes to the connection pool parameters take effect
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return db
}

// getIntEnv gets the integer value of the specified environment variable
func getIntEnv(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Printf("Error converting %s to int: %v", key, err)
	}
	return value
}

func init() {
	DB = InitDB()
}
