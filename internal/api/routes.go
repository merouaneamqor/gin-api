// internal/api/routes.go
package api

import (
    "github.com/gin-gonic/gin"

)

func RegisterRoutes(router *gin.Engine) {
    userGroup := router.Group("/api")
    RegisterUserRoutes(userGroup)
    SetupUserRoutes(userGroup)


}

func pingHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}
