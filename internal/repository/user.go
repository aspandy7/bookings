package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aspandy7/booking/internal/domain"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) (int, error) {
	const query = `
	INSERT INTO users (name, phone, password_hash)
	VALUES ($1, $2, $3)
	RETURNING id;`

	var id int
	var pgErr *pgconn.PgError

	err := r.pool.QueryRow(ctx, query, user.Name, user.Phone, user.PasswordHash).Scan(&id)

	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return 0, domain.ErrPhoneTaken
	}

	if err != nil {
		return 0, fmt.Errorf("insert user: %w", err)
	}

	return id, nil
}
