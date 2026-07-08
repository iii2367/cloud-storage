package database

import (
	"fmt"
	"context"
	"cloud-storage/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Database) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s",
        cfg.User,
        cfg.Password,
        cfg.Host,
        cfg.Port,
        cfg.DBName,
    )

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
        return nil, err
    }

	if err := pool.Ping(context.Background()); err != nil {
        pool.Close()
        return nil, err
    }

	return pool, nil
}
