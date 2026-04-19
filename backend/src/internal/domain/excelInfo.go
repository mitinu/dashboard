package domain

import "backend/src/internal/DTO"

type ExcelInfo interface {
	GetByName(name string) (*DTO.ExcelInfo, error)
	CreateOrUpdate(file *DTO.ExcelInfo) error
}
