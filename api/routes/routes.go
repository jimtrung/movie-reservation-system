package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimtrung/movie-reservation-system/api/handlers"
	"github.com/jimtrung/movie-reservation-system/internal/middleware"
)

func SetupRoutes(r *gin.Engine) {
    // --> User <-- //
    // -> JWT Auth <-
    user := r.Group("/user")
    user.POST("/signup", handlers.Signup)
    user.POST("/login", handlers.Login)
    user.GET("/validate", middleware.RequireAuth, handlers.Validate)

    // -> OAuth <-
    auth := r.Group("/auth")
    auth.GET("/:provider", middleware.BeginAuthProviderCallback)
    auth.GET("/:provider/callback", middleware.GetAuthCallBackFunction)

    // --> Movie <-- //
    movie := r.Group("/movie")
    movie.POST("/", middleware.RequireAuth, middleware.IsAdmin, handlers.AddMovie)
    movie.PUT("/:movie_id", middleware.RequireAuth, middleware.IsAdmin, handlers.UpdateMovie)
    movie.DELETE("/:movie_id", middleware.RequireAuth, middleware.IsAdmin, handlers.DeleteMovie)

    SetupFilePath(r)
}
