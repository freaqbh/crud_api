package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in .env file")
	}

	var errConn error
	DB, errConn = pgx.Connect(context.Background(), dbURL)
	if errConn != nil {
		log.Fatal("Unable to connect to database: ", errConn)
	}

	fmt.Println("Connected to database")
}
