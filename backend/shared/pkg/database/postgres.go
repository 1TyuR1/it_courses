// shared/pkg/database/postgres.go
package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string

	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
}

func NewPostgresPool(ctx context.Context, cfg Config, logger *slog.Logger) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSLMode,
	)

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns
	poolConfig.MaxConnLifetime = cfg.MaxConnLifetime
	poolConfig.MaxConnIdleTime = cfg.MaxConnIdleTime
	poolConfig.HealthCheckPeriod = cfg.HealthCheckPeriod

	logger.Info("Connecting to PostgreSQL",
		slog.String("host", cfg.Host),
		slog.Int("port", cfg.Port),
		slog.String("database", cfg.Database),
	)

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	// Исправлено: используем slog.Int() вместо slog.Int32()
	logger.Info("Successfully connected to PostgreSQL",
		slog.Int("max_conns", int(cfg.MaxConns)), // конвертируем int32 → int
		slog.Int("min_conns", int(cfg.MinConns)), // конвертируем int32 → int
	)

	return pool, nil
}

func DefaultConfig() Config {
	return Config{
		SSLMode:           "disable",
		MaxConns:          25,
		MinConns:          5,
		MaxConnLifetime:   time.Hour,
		MaxConnIdleTime:   time.Minute * 30,
		HealthCheckPeriod: time.Minute,
	}
}
