package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    _ "myapp/docs" // Import auto-generated docs
)

// @title My Gin API
// @version 1.0
// @description This is a sample Gin API documentation.

// @host localhost:8085
// @BasePath /

func main() {
    r := gin.Default()

    // Swagger endpoint
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // @Summary Ping route
    // @Description Check if the API is alive
    // @Success 200 {string} string "pong"
    // @Router /ping [get]
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    // @Summary DB test route
    // @Description Check if the DB connection is working
    // @Success 200 {string} string "db connected"
    // @Router /dbtest [get]
    r.GET("/dbtest", func(c *gin.Context) {
        // Ici, tu ajouteras la logique pour interagir avec ta base de données MariaDB
        c.JSON(http.StatusOK, gin.H{
            "message": "db connected",
        })
    })

    r.Run(":8080") // Démarre l'API sur le port 8080
}
