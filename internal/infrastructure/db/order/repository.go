package order

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewRepository(
	host, port, sslMode, user, password, name string,
	maxConns int32,
	logger *zap.Logger,
) (*Repository, error) {
	config, err := pgxpool.ParseConfig(buildPostgresDns(host, port, sslMode, user, password, name))
	if err != nil {
		return nil, err
	}
	config.MaxConns = maxConns

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	return &Repository{db, logger}, nil
}

func (r *Repository) Close() {}

func buildPostgresDns(host, port, sslMode, user, password, name string) string {
	return fmt.Sprintf(
		"host=%s port=%s sslmode=%s user=%s password=%s dbname=%s",
		host, port, sslMode, user, password, name,
	)
}
