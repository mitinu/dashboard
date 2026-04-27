package domain

import "backend/src/internal/DTO"

type AverageSalary interface {
	Create(averageSalary *DTO.AverageSalary) error
	Creates(averagesSalary *[]DTO.AverageSalary) []error
	DeleteByIdTable(idTable int64) (int64, error)
	GetByDateRange(startDate, endDate string) ([]DTO.AverageSalary, error)
}
