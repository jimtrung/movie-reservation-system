package models

import "time"

type Showtimes struct {
    ShowtimeID int          `json:"showtime_id"`
    MovieID    int          `json:"movie_id"`
    Showtimes  *[]time.Time `json:"showtimes"`
    CreatedAt  *time.Time   `json:"created_at"`
    UpdatedAt  *time.Time   `json:"updated_at"`
}
