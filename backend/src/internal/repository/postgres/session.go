package postgres

import (
	"backend/src/internal/config"
	"backend/src/internal/domain"
	"backend/src/internal/repository/SQL"
	"database/sql"
	"fmt"
)

type SessionRepository struct {
	DB  *sql.DB
	cfg *config.Config
}

func NewSessionRepository(Postgres *PostgresDB, cfg *config.Config) domain.Session {
	return &SessionRepository{
		DB:  Postgres.DB,
		cfg: cfg,
	}
}

func (r *SessionRepository) Create(userId int64, token string) error {
	err := r.DB.QueryRow(SQL.СreteSession, userId, token)
	if err != nil {
		return fmt.Errorf("не удалось создать запись: %w", err)
	}
	return nil
}
