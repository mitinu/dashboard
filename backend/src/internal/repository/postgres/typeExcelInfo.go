package postgres

import (
	"backend/src/internal/DTO"
	"backend/src/internal/domain"
	"database/sql"
	"errors"
	"fmt"
)

type TypeExcelInfoRepository struct {
	DB *sql.DB
}

func NewTypeExcelInfoRepository(Postgres *PostgresDB) domain.TypeExcelInfo {
	return &TypeExcelInfoRepository{DB: Postgres.DB}
}

func (r *TypeExcelInfoRepository) GetIdByTitle(title string) (int64, error) {
	var id int64
	err := r.DB.QueryRow(DTO.GetIdByTitleTypeExcelInfo, title).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("тип Excel файла с названием '%s' не найден: %w", title, err)
		}
		return 0, fmt.Errorf("ошибка при получении ID типа Excel файла по названию: %w", err)
	}

	return id, nil
}
