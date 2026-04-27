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
