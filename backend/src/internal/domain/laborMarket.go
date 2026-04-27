package domain

import "backend/src/internal/DTO"

type LaborMarket interface {
	Create(laborMarket *DTO.LaborMarket) error
	Creates(laborMarkets *[]DTO.LaborMarket) []error
	DeleteByIdTable(idTable int64) (int64, error)
	GetLaborMarketByDateRange(startDate, endDate string) ([]DTO.LaborMarket, error)
}
