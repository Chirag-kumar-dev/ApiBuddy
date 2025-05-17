package main

import (
    "ApiBuddy/config"
    "ApiBuddy/models"
    "ApiBuddy/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Connect to the database
    config.ConnectDatabase()

    // Auto-migrate the User model
    config.DB.AutoMigrate(&models.User{})

    // Set up the Gin router
    router := gin.Default()

    // Register authentication routes
    routes.AuthRoutes(router)

    // Start the server
    router.Run(":8080")
}
