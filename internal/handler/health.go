package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Handler {
	return &Handler{pool: pool}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := h.pool.Ping(ctx); err != nil {
		log.Printf("health: ping database: %v", err)
		http.Error(w, "database unavailable", http.StatusServiceUnavailable)
		return
	}

	if _, err := w.Write([]byte("ok")); err != nil {
		log.Printf("health: write response: %v", err)
	}
}
