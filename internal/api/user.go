// internal/api/user.go
package api

import (
	"fmt"
	"net/http"

	"douq.merouaneamqor.com/internal/db"
	"douq.merouaneamqor.com/internal/model"
	"douq.merouaneamqor.com/internal/util"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.GET("/users/fake", createFakeUsers)
}

func createFakeUsers(c *gin.Context) {
	var users []model.User
	number_users:=60000
	for i := 0; i < number_users; i++ {

		users = append(users, util.GenerateFakeUser())
	}

	// Batch insert the fake users into the database using GORM
	result := db.DB.CreateInBatches(users, 100) // Using a batch size of 100
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d fake users created successfully",number_users)})

}
