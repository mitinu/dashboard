package domain

import (
	"backend/src/internal/DTO"
)

type CPI interface {
	Create(cpi *DTO.CPI) error
	Creates(multipleCpi *[]DTO.CPI) []error
	DeleteByIdTable(idTable int64) (int64, error)
}
