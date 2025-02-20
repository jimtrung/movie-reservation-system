package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jimtrung/movie-reservation-system/internal/models"
	"github.com/jimtrung/movie-reservation-system/internal/services"
)


func CreateSignedToken(username string) (string, error) {
    token := jwt.NewWithClaims(
        jwt.SigningMethodHS256,
        jwt.MapClaims{
            "username": username,
            "exp": time.Now().Add(time.Hour * 24).Unix(),
        },
    )

    signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
    if err != nil {
        return "", fmt.Errorf("Failed to sign the token")
    }
    return signedToken, nil
}

func RequireAuth(c *gin.Context) {
    tokenString, err := c.Cookie("Authorization")
    if err != nil {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Wrong signing method")
        }
        return []byte(os.Getenv("SECRET_KEY")), nil
    })
    if err != nil {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    exp, ok := claims["exp"].(float64)
    if !ok || float64(time.Now().Unix()) > exp {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    username, ok := claims["username"].(string)
    if !ok {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    user, err := services.GetUserByUsername(username)
    if err != nil {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    c.Set("user", user)
    c.Next()
}

func IsAdmin(c *gin.Context) {
    userAny, exist := c.Get("user")
    if !exist {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    user, ok := userAny.(models.UserResponse)
    if !ok || user.Role != models.Admin {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    c.Next()
}
