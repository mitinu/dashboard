package postgres

import (
	"backend/src/internal/domain"
	"backend/src/internal/repository/SQL"
	"database/sql"
)

type AccessStatusUserRepository struct {
	DB *sql.DB
}

func NewAccessStatusUserRepository(Postgres *PostgresDB) domain.AccessStatusUser {
	return &AccessStatusUserRepository{DB: Postgres.DB}
}

func (r *AccessStatusUserRepository) GetIdByTitle(title string) (int64, error) {
	var id int64
	err := r.DB.QueryRow(SQL.GetIdByTitle, title).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
