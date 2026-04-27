package postgres

import (
	"backend/src/internal/DTO"
	"backend/src/internal/domain"
	"backend/src/internal/repository/SQL"
	"database/sql"
	"fmt"
	"log"
)

type AverageSalaryRepository struct {
	DB *sql.DB
}

func NewAverageSalaryRepository(Postgres *PostgresDB) domain.AverageSalary {
	return &AverageSalaryRepository{DB: Postgres.DB}
}

// Create - создание одной записи
func (r *AverageSalaryRepository) Create(averageSalary *DTO.AverageSalary) error {
	var id int64
	err := r.DB.QueryRow(SQL.CreateAverageSalary, averageSalary.Date, averageSalary.Value, averageSalary.IdTable).Scan(&id)
	if err != nil {
		log.Printf("Ошибка при вставке записи: %v", err)
		return fmt.Errorf("не удалось создать запись: %w", err)
	}

	averageSalary.ID = id
	return nil
}

// Creates - массовое создание записей, возвращает слайс ошибок
func (r *AverageSalaryRepository) Creates(averagesSalary *[]DTO.AverageSalary) []error {
	errs := []error{}
	for _, averageSalary := range *averagesSalary {
		err := r.Create(&averageSalary)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs

}

// DeleteByIdTable - удаление всех записей по id_table, возвращает количество удаленных записей
func (r *AverageSalaryRepository) DeleteByIdTable(idTable int64) (int64, error) {
	result, err := r.DB.Exec(SQL.DeleteByIdTableAverageSalary, idTable)
	if err != nil {
		log.Printf("Ошибка при удалении записей с id_table=%d: %v", idTable, err)
		return 0, fmt.Errorf("не удалось удалить записи: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("не удалось получить количество удаленных строк: %w", err)
	}

	log.Printf("Удалено %d записей с id_table=%d", rowsAffected, idTable)
	return rowsAffected, nil
}
