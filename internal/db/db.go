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

// initialize_database_config initializes the database configuration by loading environment variables
func initialize_database_config() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set default values for connection pool parameters if they are not already set
	set_default_env("DB_MaxOpenConns", "5")
	set_default_env("DB_MaxIdleConns", "2")
}

// set_default_env sets the specified environment variable to the default value if it is not set
func set_default_env(key, defaultValue string) {
	if os.Getenv(key) == "" {
		os.Setenv(key, defaultValue)
	}
}

// build_dsn builds the connection string for PostgreSQL based on environment variables
func build_dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
}

// initialize_database initializes and returns a database connection
func initialize_database() *gorm.DB {
	initialize_database_config()

	// Build the connection string for PostgreSQL
	dsn := build_dsn()

	// Establish a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Ping the database to verify the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get the database connection: %v", err)
	}

	// Configure connection pool parameters
	sqlDB.SetMaxIdleConns(get_int_env("DB_MaxIdleConns"))
	sqlDB.SetMaxOpenConns(get_int_env("DB_MaxOpenConns"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping the database again to ensure the changes to the connection pool parameters take effect
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return db
}

// get_int_env gets the integer value of the specified environment variable
func get_int_env(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Printf("Error converting %s to an integer: %v", key, err)
	}
	return value
}

// init is automatically called when the package is imported
func init() {
	DB = initialize_database()
}
