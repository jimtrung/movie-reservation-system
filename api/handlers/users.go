package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/movie-reservation-system/internal/middleware"
	"github.com/jimtrung/movie-reservation-system/internal/models"
	"github.com/jimtrung/movie-reservation-system/internal/services"
)

func Signup(c *gin.Context) {
    var userRequest models.SignUpRequest

    if err := c.Bind(&userRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Wrong JSON format",
        })
        return
    }

    hashedPassword, err := services.HashPassword(userRequest.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to hash the password",
        })
        return
    }

    userRequest.Password = hashedPassword
    if err := services.AddUserToDatabse(userRequest); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User signed up successfully",
    })
}

func Login(c *gin.Context) {
    var userRequest models.LoginRequest

    if err := c.Bind(&userRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Wrong JSON format",
        })
        return
    }

    if err := services.IsValidUser(userRequest.Username, userRequest.Password);
    err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Create a jwt token and add to cookie
    signedToken, err := middleware.CreateSignedToken(userRequest.Username)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": err.Error(),
        })
    }

    // Add token to cookie
	c.SetSameSite(http.SameSiteLaxMode)
    c.SetCookie("Authorization", signedToken, 60*60*24, "/", "", false, true)

    // Print out success message
    c.JSON(http.StatusOK, gin.H{
        "message": "Login successfully",
    })
}

func Validate(c *gin.Context) {
    userInfo, ok := c.Get("user")
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to get user from cookie",
        })
        return
    }

    c.JSON(http.StatusOK, userInfo)
}
