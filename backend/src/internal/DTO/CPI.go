package DTO

import (
	"time"
)

type CPI struct {
	ID      int64
	Date    time.Time
	Value   float64
	IdTable int64
}

const (
	CreateCPI = `
		INSERT INTO public."CPI" (date, value, id_table)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	DeleteByIdTableCPI = `
		DELETE FROM public."CPI" WHERE id_table = $1
	`
)
