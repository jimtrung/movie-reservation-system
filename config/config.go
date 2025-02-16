package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn
var PORT string

func SetupDatabaseConnection() error {
    if err := godotenv.Load(); err != nil {
        return fmt.Errorf("Failed to load .env file")
    }

    conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
    if err != nil {
        return fmt.Errorf("Failed to connect to database")
    }

    DB = conn
    fmt.Println("Connected to database successfully")

    return nil
}

func CloseDatabaseConnection() {
    DB.Close(context.Background())
}

func SetupPort() {
    PORT = os.Getenv("PORT")
}
