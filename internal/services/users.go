package services

import (
	"context"
	"fmt"

	"github.com/jimtrung/movie-reservation-system/config"
	"github.com/jimtrung/movie-reservation-system/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    if err != nil {
        return "", fmt.Errorf("Failed to generate hash from password")
    }

    return string(hashedBytes), nil
}

func AddUserToDatabse(user models.SignUpRequest) error {
    _, err := config.DB.Exec(
        context.Background(),
        `INSERT INTO users(username, email, password) VALUES ($1, $2, $3)`,
        user.Username, user.Email, user.Password,
    )
    if err != nil {
        return fmt.Errorf("Failed to insert user data into database")
    }

    return nil
}

func IsValidUser(username, password string) error {
    row := config.DB.QueryRow(
        context.Background(),
        `SELECT password FROM users WHERE username = $1`,
        username,
    )

    var hashedPassword string
    if err := row.Scan(&hashedPassword); err != nil {
        return fmt.Errorf("Failed to get user info from database")
    }

    if err := bcrypt.CompareHashAndPassword(
        []byte(hashedPassword),
        []byte(password),
    ); err != nil {
        return fmt.Errorf("Wrong password")
    }

    return nil
}


func GetUserByUsername(username string) (models.UserResponse, error) {
    row := config.DB.QueryRow(
        context.Background(),
        `SELECT user_id, username, role FROM users
         WHERE username = $1`,
        username,
    )

    var userResponse models.UserResponse
    var roleString models.Role
    if err := row.Scan(
        &userResponse.UserID,
        &userResponse.Username,
        &roleString,
    ); err != nil {
        return models.UserResponse{}, fmt.Errorf("Failed to get user")
    }
    userResponse.Role = models.Role(roleString)

    return userResponse, nil
}
