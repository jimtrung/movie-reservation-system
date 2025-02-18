package services

import (
	"context"
	"fmt"

	"github.com/jimtrung/movie-reservation-system/config"
	"github.com/jimtrung/movie-reservation-system/internal/models"
)

func AddMovieToDatabase(movie models.AddMovieRequest) error {
    _, err := config.DB.Exec(
        context.Background(),
        `INSERT INTO movies (title, description, image, genre)
         VALUES ($1, $2, $3, $4)`,
        movie.Title, movie.Description, movie.Image, models.Genre(movie.Genre),
    )

    return err
}

func UpdateMovieField(movieID int, req models.UpdateRequest) error {
    updateQuery := fmt.Sprintf(
        `UPDATE movies SET %s = $1 WHERE movie_id = $2`,
        req.Field,
    )

    _, err := config.DB.Exec(
        context.Background(),
        updateQuery,
        req.Value, movieID,
    )

    return err
}

func DeleteMovieFromDatabase(movieID int) error {
    _, err := config.DB.Exec(
        context.Background(),
        `DELETE FROM movies WHERE movie_id = $1`,
        movieID,
    )

    return err
}
