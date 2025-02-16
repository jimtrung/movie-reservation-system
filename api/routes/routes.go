package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimtrung/movie-reservation-system/api/handlers"
)

func SetupRoutes(r *gin.Engine) {
    // --> User <-- //
    r.POST("/user/signup", handlers.Signup)
    r.POST("/user/login", handlers.Login)
}
