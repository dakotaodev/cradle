package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dakotaodev/cradle/internal/api"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println(".env not found; using environment variables")
	}
	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal("DATABASE_URL must be provided.")
	}


	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	defer pool.Close()

	router := api.NewRouter()
	if err := router.Run(); err != nil {
		log.Fatalf("unable to start the router: %v", err)
	}

}
