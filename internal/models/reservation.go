package models

import "time"

type MovieResponse struct {
    Title     string      `json:"title"     db:"title"`
    Showtimes []time.Time `json:"showtimes" db:"showtimes"`
}
