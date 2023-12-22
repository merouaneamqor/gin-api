package api

import (
	"douq.merouaneamqor.com/controllers"

	"github.com/gin-gonic/gin"

)

func SetupUserRoutes(router *gin.RouterGroup) {
	usersGroup := router.Group("/users")
	usersGroup.GET("/", controllers.GetUsers)
}