package postgres

import (
	"backend/src/internal/DTO"
	"backend/src/internal/domain"
	"database/sql"
	"fmt"
)

type CpiRepository struct {
	DB *sql.DB
}

func NewCpiRepository(Postgres *PostgresDB) domain.CPI {
	return &CpiRepository{DB: Postgres.DB}
}

func (r *CpiRepository) Create(cpi *DTO.CPI) error {
	err := r.DB.QueryRow(DTO.CreateCPI, cpi.Date, cpi.Value, cpi.IdTable).Scan(&cpi.ID)
	if err != nil {
		return fmt.Errorf("failed to insert CPI: %w", err)
	}
	return nil
}

func (r *CpiRepository) Creates(multipleCpi *[]DTO.CPI) []error {
	errs := []error{}
	for _, cpi := range *multipleCpi {
		err := r.Create(&cpi)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (r *CpiRepository) DeleteByIdTable(idTable int64) (int64, error) {
	result, err := r.DB.Exec(DTO.DeleteByIdTableCPI, idTable)
	if err != nil {
		return 0, fmt.Errorf("ошибка при удалении записей: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("ошибка при получении количества удаленных строк: %w", err)
	}

	return rowsAffected, nil
}
