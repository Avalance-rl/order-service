package order

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Pgx struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewRepository(
	host, port, sslMode, user, password, name string,
	maxConns int32,
	logger *zap.Logger,
) (*Pgx, error) {
	config, err := pgxpool.ParseConfig(buildPostgresDns(host, port, sslMode, user, password, name))
	if err != nil {
		return nil, err
	}
	config.MaxConns = maxConns

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return &Pgx{db, logger}, nil
}

func (r *Pgx) Close(ctx context.Context) error {
	done := make(chan struct{}, 1)
	go func() {
		r.pool.Close()
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}

func buildPostgresDns(host, port, sslMode, user, password, name string) string {
	return fmt.Sprintf(
		"host=%s port=%s sslmode=%s user=%s password=%s dbname=%s",
		host, port, sslMode, user, password, name,
	)
}
