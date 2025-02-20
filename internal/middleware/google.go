package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/jimtrung/movie-reservation-system/internal/models"
	"github.com/jimtrung/movie-reservation-system/internal/services"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "jimtrungzxjehihihehehahahoho"
	maxAge = 3600 * 24
	isProd = true
)

func NewAuth() {
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(
			googleClientId,
			googleClientSecret,
			"http://127.0.0.1:8080/auth/google/callback",
		),
	)
}

func BeginAuthProviderCallback(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(
        context.Background(),
		"provider",
		provider,
	))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func GetAuthCallBackFunction(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(
        context.Background(),
		"provider",
		provider,
	))

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var userInfo = models.SignUpRequest{
		Username: user.Email,
		Password: "email",
		Email:    user.Email,
	}

	if err := services.AddUserToDatabse(userInfo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

    signedToken, err := CreateSignedToken(userInfo.Username)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Failed to sign token",
        })
        return
    }

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", signedToken, 3600*24, "/", "", false, true)

	c.Redirect(http.StatusFound, "http://127.0.0.1:8080/validate")
}
