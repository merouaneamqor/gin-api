// api/handlers.go
package api

import (
    "github.com/gin-gonic/gin"
    "douq.merouaneamqor.com/internal/db"
    "douq.merouaneamqor.com/internal/model"
)

func getUsersHandler(c *gin.Context) {
    var users []model.User

    // Use GORM's Find method to retrieve all users
    result := db.DB.Find(&users)
    if result.Error != nil {
        // Handle the error appropriately
        c.JSON(500, gin.H{"error": result.Error.Error()})
        return
    }

    // Send the retrieved users as JSON
    c.JSON(200, gin.H{"users": users})
}
