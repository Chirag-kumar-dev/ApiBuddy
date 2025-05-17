package controllers

import (
    "ApiBuddy/config"
    "ApiBuddy/models"
    "ApiBuddy/utils"
    "net/http"
    "github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
    var newUser models.User

    if err := c.ShouldBindJSON(&newUser); err != nil {
        utils.SendError(c, http.StatusBadRequest, "Invalid request")
        return
    }

    // Check if username already exists
    var existingUser models.User
    if err := config.DB.Where("username = ?", newUser.Username).First(&existingUser).Error; err == nil {
        utils.SendError(c, http.StatusBadRequest, "Username already exists")
        return
    }

    // Save new user
    if err := config.DB.Create(&newUser).Error; err != nil {
        utils.SendError(c, http.StatusInternalServerError, "Failed to create user")
        return
    }

    utils.SendSuccess(c, "Signup successful", nil)
}

func Login(c *gin.Context) {
    var loginUser models.User

    if err := c.ShouldBindJSON(&loginUser); err != nil {
        utils.SendError(c, http.StatusBadRequest, "Invalid request")
        return
    }

    var user models.User
    if err := config.DB.Where("username = ?", loginUser.Username).First(&user).Error; err != nil {
        utils.SendError(c, http.StatusUnauthorized, "Invalid credentials")
        return
    }

    if user.Password != loginUser.Password {
        utils.SendError(c, http.StatusUnauthorized, "Invalid credentials")
        return
    }

    utils.SendSuccess(c, "Login successful", nil)
}
