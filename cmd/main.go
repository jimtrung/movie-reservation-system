package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/movie-reservation-system/api/routes"
	"github.com/jimtrung/movie-reservation-system/config"
)

func main() {
    // --> Config <--
    if err := config.SetupDatabaseConnection(); err != nil {
        log.Fatal(err)
    }
    defer config.CloseDatabaseConnection()

    config.SetupPort()

    // --> Server <--
    r := gin.New()

    // Set up routes
    routes.SetupRoutes(r)

    r.SetTrustedProxies([]string{"127.0.0.1"})

    r.Run("127.0.0.1:" + config.PORT)
}
