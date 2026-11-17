package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aspandy7/booking/internal/config"
	"github.com/aspandy7/booking/internal/handler"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("err to create pool: %w", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	fmt.Println("connected to database")

	h := handler.New(pool)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)

	srv := &http.Server{Addr: ":" + cfg.ServerPort, Handler: mux, ReadHeaderTimeout: 5 * time.Second}

	return srv.ListenAndServe()
}
