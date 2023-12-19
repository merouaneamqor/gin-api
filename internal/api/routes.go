// internal/api/routes.go
package api

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
    userGroup := router.Group("/api")
    RegisterUserRoutes(userGroup)

    router.GET("/ping", pingHandler)
	router.GET("/users", getUsersHandler)

}

func pingHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}
