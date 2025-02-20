package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimtrung/movie-reservation-system/api/handlers"
	"github.com/jimtrung/movie-reservation-system/internal/middleware"
)

func SetupRoutes(r *gin.Engine) {
    // --> User <-- //
    // -> JWT Auth <-
    r.POST("/user/signup", handlers.Signup)
    r.POST("/user/login", handlers.Login)
    r.GET("/validate", middleware.RequireAuth, handlers.Validate)

    // -> OAuth <-
    r.GET("/auth/:provider", middleware.BeginAuthProviderCallback)
    r.GET("/auth/:provider/callback", middleware.GetAuthCallBackFunction)

    // --> Movie <-- //
    r.POST("/movie/add", middleware.RequireAuth, middleware.IsAdmin, handlers.AddMovie)
    r.PUT("/movie/update/:movie_id", middleware.RequireAuth, middleware.IsAdmin, handlers.UpdateMovie)
    r.DELETE("/movie/delete/:movie_id", middleware.RequireAuth, middleware.IsAdmin, handlers.DeleteMovie)

    SetupFilePath(r)
}
