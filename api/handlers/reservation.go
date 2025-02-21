package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/movie-reservation-system/internal/services"
)

func GetShowtimes(c *gin.Context) {
    dateString := c.Param("date")

    date, err := time.Parse("2006-01-02", dateString)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Wrong time format",
        })
        return
    }

    showtimesRes, err := services.GetShowtimesByDate(date)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, showtimesRes)
}
