package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/movie-reservation-system/internal/models"
	"github.com/jimtrung/movie-reservation-system/internal/services"
)

func AddMovie(c *gin.Context) {
    var req models.AddMovieRequest

    if err := c.Bind(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Wrong JSON format",
        })
        return
    }

    if err := services.AddMovieToDatabase(req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "User added to database successfully",
    })
}

func UpdateMovie(c *gin.Context) {


}

func DeleteMovie(c *gin.Context) {
    movieIDString := c.Param("movie_id")

    movieID, err := strconv.Atoi(movieIDString)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to convert string into integer",
        })
        return
    }

    if err := services.DeleteMovieFromDatabase(movieID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to delete movie from database",
        })
        return
    }

    c.JSON(http.StatusNoContent, gin.H{})
}
