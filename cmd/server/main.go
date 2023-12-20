package main

import (
	"log"

	"douq.merouaneamqor.com/internal/api"
	"douq.merouaneamqor.com/internal/db"
	"douq.merouaneamqor.com/internal/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize the database connection


	// AutoMigrate the models
	if err := db.DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// Register the routes
	api.RegisterRoutes(r)

	// Start the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
