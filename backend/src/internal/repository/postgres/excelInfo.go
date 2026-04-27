package postgres

import (
	"backend/src/internal/DTO"
	"backend/src/internal/domain"
	"backend/src/internal/repository/SQL"
	"database/sql"
	"fmt"
)

type ExcelInfoRepository struct {
	DB *sql.DB
}

func NewExcelInfoRepository(Postgres *PostgresDB) domain.ExcelInfo {
	return &ExcelInfoRepository{DB: Postgres.DB}
}

func (r *ExcelInfoRepository) GetByName(name string) (*DTO.ExcelInfo, error) {
	// Здесь пишем РЕАЛЬНЫЙ SQL запрос
	var excelInfo DTO.ExcelInfo

	err := r.DB.QueryRow(SQL.GetByNameExcelInfo, name).Scan(
		&excelInfo.ID,
		&excelInfo.IdType,
		&excelInfo.Name,
		&excelInfo.DateUpdated,
		&excelInfo.Hash,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("file not found: %s", name)
	}
	if err != nil {
		return nil, err
	}

	return &excelInfo, nil
}

func (r *ExcelInfoRepository) CreateOrUpdate(file *DTO.ExcelInfo) error {
	err := r.DB.QueryRow(SQL.CreateOrUpdateExcelInfo, file.Name, file.Hash, file.IdType).Scan(&file.ID)
	if err != nil {
		return err
	}

	return nil
}
