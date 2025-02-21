package services

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jimtrung/movie-reservation-system/config"
	"github.com/jimtrung/movie-reservation-system/internal/models"
)

func GetShowtimesByDate(date time.Time) ([]models.MovieResponse, error) {
    rows, err := config.DB.Query(
        context.Background(),
        `SELECT movies.title, showtimes.showtimes
        FROM movies
        LEFT JOIN showtimes
        ON movies.movie_id = showtimes.movie_id
        WHERE $1 = ANY(showtimes.showtimes::DATE[]);`,
        date,
    )
    if err != nil {
        return []models.MovieResponse{}, err
    }

    showtimesRes, err := pgx.CollectRows(
        rows,
        pgx.RowToStructByName[models.MovieResponse],
    )

    return showtimesRes, err
}
