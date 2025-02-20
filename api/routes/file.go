package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupFilePath(r *gin.Engine) {
    r.LoadHTMLGlob("./static/*")

    r.GET("/login", func(ctx *gin.Context) {
        ctx.HTML(http.StatusOK, "login", gin.H{})
    })

    r.GET("/signup", func(ctx *gin.Context) {
        ctx.HTML(http.StatusOK, "signup", gin.H{})
    })

    r.GET("/profile", func(ctx *gin.Context) {
        ctx.HTML(http.StatusOK, "profile", gin.H{})
    })

    r.GET("/forgotpassword", func(ctx *gin.Context) {
        ctx.HTML(http.StatusOK, "forgotpassword", gin.H{})
    })
}
