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
    r.POST("/movie/add", middleware.IsAdmin, handlers.AddMovie)
    r.PUT("/movie/update/:movie_id", middleware.IsAdmin, handlers.UpdateMovie)
    r.DELETE("/movie/delete/:movie_id", middleware.IsAdmin, handlers.DeleteMovie)

    // --> Static File <-- //
    r.GET("/signup", func(ctx *gin.Context) {
        ctx.File("./static/signup.html")
    })

    r.GET("/login", func(ctx *gin.Context) {
        ctx.File("./static/login.html")
    })

    r.GET("/forgotpassword", func(ctx *gin.Context) {
        ctx.File("./static/forgotpassword.html")
    })

    r.GET("/profile", func(ctx *gin.Context) {
        ctx.File("./static/profile.html")
    })
}
