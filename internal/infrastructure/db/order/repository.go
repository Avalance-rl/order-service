package order

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(connString string, maxConns int32, logger *slog.Logger) (*Repository, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	config.MaxConns = maxConns

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	return &Repository{db}, nil
}

func (r *Repository) Close() {}
