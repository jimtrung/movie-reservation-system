package models

import "time"

type Role int

const (
    Admin Role = 10
    Moderator  = 1
    Client     = 0
)

type User struct {
    UserID    int        `json:"user_id"`
    Username  string     `json:"username"`
    Password  string     `json:"password"`
    Role      Role       `json:"role"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

type SignUpRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type UserResponse struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username"`
    Role     Role   `json:"role"`
}
