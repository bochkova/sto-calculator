package app

import (
	"context"
	"fmt"

	"sto-calculator/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func initDB(ctx context.Context, cfg config.DatabaseConfig) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database,
	)
	db, err := sqlx.Open("pgx", dataSource)

	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}
