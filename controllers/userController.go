// userController.go

package controllers

import (
	"net/http"
	"strconv"

	"douq.merouaneamqor.com/internal/db"
	"douq.merouaneamqor.com/internal/model"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []model.User

	// Default values for pagination
	defaultLimit := 10
	defaultPage := 1

	// Read query parameters
	page, err := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
	if err != nil || page < 1 {
		page = defaultPage
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(defaultLimit)))
	if err != nil || limit < 1 {
		limit = defaultLimit
	}

	offset := (page - 1) * limit

	// Use GORM's pagination feature to retrieve a subset of users
	result := db.DB.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Send the paginated users as JSON
	c.JSON(http.StatusOK, gin.H{"users": users, "page": page, "limit": limit})
}
