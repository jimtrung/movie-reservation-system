package models

import "time"

type Genre string

const (
    Action    Genre = "Action"
    Comedy    Genre = "Comedy"
    Drama     Genre = "Drama"
    Horror    Genre = "Horror"
    SciFi     Genre = "Sci-Fi"
    Romance   Genre = "Romance"
    Fantasy   Genre = "Fantasy"
    Thriller  Genre = "Thriller"
    Adventure Genre = "Adventure"
)

type Movie struct {
    MovieID     int        `json:"movie_id"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Image       string     `json:"image"`
    Genre       Genre      `json:"genre"`
    CreatedAt   *time.Time `json:"createx_at"`
    UpdatedAt   *time.Time `json:"updated_at"`
}

type AddMovieRequest struct {
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Image       string     `json:"image"`
    Genre       Genre      `json:"genre"`
}

type Field string

const (
    Title Field = "title"
    Description = "description"
    Image       = "image"
    GenreField  = "genre"
)

type UpdateRequest struct {
    Field string `json:"field"`
    Value string `json:"value"`
}
