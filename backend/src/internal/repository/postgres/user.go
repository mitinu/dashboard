package postgres

import (
	"backend/src/internal/DTO"
	"backend/src/internal/config"
	"backend/src/internal/domain"
	"backend/src/internal/repository/SQL"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	DB  *sql.DB
	cfg *config.Config
}

func NewUserRepository(Postgres *PostgresDB, cfg *config.Config) domain.User {
	return &UserRepository{
		DB:  Postgres.DB,
		cfg: cfg,
	}
}

func (r *UserRepository) Create(user *DTO.User) error {
	result, err := r.DB.Exec(SQL.CreateUser, user.AccessStatusId, user.Login, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("не удалось создать запись: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("не удалось создать запись: строка не была добавлена")
	}
	return nil
}

func (r *UserRepository) GetByLogin(login string) (*DTO.User, error) {
	var user DTO.User
	err := r.DB.QueryRow(SQL.GetUserByLogin, login).Scan(
		&user.ID,
		&user.Login,
		&user.PasswordHash,
		&user.AccessStatusId,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) DeleteSuperadmin() error {
	result, err := r.DB.Exec(SQL.DeleteUserSuperadmin, r.cfg.TitleSadminInDB)
	if err != nil {
		return fmt.Errorf("не удалось удалить запись: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("не удалось удалить запись")
	}
	return nil
}
