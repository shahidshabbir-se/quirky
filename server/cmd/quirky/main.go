package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/shahidshabbir-se/quirky/internal/api"
	db "github.com/shahidshabbir-se/quirky/internal/db/sqlc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	conn, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close()

	queries := db.New(conn)

	server := api.NewServer(queries)
	err = server.Start(":8080")
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
