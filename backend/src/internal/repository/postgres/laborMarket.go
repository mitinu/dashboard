package postgres

import (
	"backend/src/internal/DTO"
	"backend/src/internal/domain"
	"backend/src/internal/repository/SQL"
	"database/sql"
	"fmt"
	"log"
)

type LaborMarketRepository struct {
	DB *sql.DB
}

func NewLaborMarketRepository(Postgres *PostgresDB) domain.LaborMarket {
	return &LaborMarketRepository{DB: Postgres.DB}
}

// Create - создание одной записи
func (r *LaborMarketRepository) Create(laborMarket *DTO.LaborMarket) error {
	var id int64
	err := r.DB.QueryRow(SQL.CreateLaborMarket, laborMarket.Date, laborMarket.Unemployed, laborMarket.Employed, laborMarket.IdTable).Scan(&id)
	if err != nil {
		log.Printf("Ошибка при вставке записи: %v", err)
		return fmt.Errorf("не удалось создать запись: %w", err)
	}

	laborMarket.ID = id
	return nil
}

// Creates - массовое создание записей, возвращает слайс ошибок
func (r *LaborMarketRepository) Creates(laborMarkets *[]DTO.LaborMarket) []error {
	errs := []error{}
	for _, laborMarket := range *laborMarkets {
		err := r.Create(&laborMarket)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// DeleteByIdTable - удаление всех записей по id_table, возвращает количество удаленных записей
func (r *LaborMarketRepository) DeleteByIdTable(idTable int64) (int64, error) {
	result, err := r.DB.Exec(SQL.DeleteByIdTableLaborMarket, idTable)
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

func (r *LaborMarketRepository) GetLaborMarketByDateRange(startDate, endDate string) ([]DTO.LaborMarket, error) {
	rows, err := r.DB.Query(SQL.GetLaborMarketByDateRange, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("ошибка при поиске в диапазоне дат: %w", err)
	}
	defer rows.Close()

	var results []DTO.LaborMarket
	for rows.Next() {
		var lm DTO.LaborMarket
		err := rows.Scan(&lm.ID, &lm.Date, &lm.Unemployed, &lm.Employed, &lm.IdTable)
		if err != nil {
			return nil, err
		}
		results = append(results, lm)
	}

	return results, nil
}
