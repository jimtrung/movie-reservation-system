package services

import (
	"context"

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

func DeleteMovieFromDatabase(movieID int) error {
    _, err := config.DB.Exec(
        context.Background(),
        `DELETE FROM movies WHERE movie_id = $1`,
        movieID,
    )

    return err
}
