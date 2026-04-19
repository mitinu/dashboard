package postgres

import (
	"backend/src/internal/config"
	"backend/src/pkg/logger"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(cfg *config.Config) (*PostgresDB, error) {
	dsn := cfg.DSN()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Настраиваем пул соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	// Проверяем подключение с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info.Println("Successfully connected to PostgreSQL")

	return &PostgresDB{DB: db}, nil
}

// Close закрывает соединение с БД
func (p *PostgresDB) Close() error {
	if p.DB != nil {
		logger.Info.Println("Closing database connection...")
		return p.DB.Close()
	}
	return nil
}

// Ping проверяет соединение с БД
func (p *PostgresDB) Ping() error {
	if p.DB == nil {
		return fmt.Errorf("database connection is nil")
	}
	return p.DB.Ping()
}

// GetStats возвращает статистику пула соединений
func (p *PostgresDB) GetStats() sql.DBStats {
	if p.DB == nil {
		return sql.DBStats{}
	}
	return p.DB.Stats()
}
