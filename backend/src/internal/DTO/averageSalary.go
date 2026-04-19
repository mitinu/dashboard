package DTO

import (
	"time"
)

type AverageSalary struct {
	ID      int64
	Date    time.Time
	Value   int
	IdTable int64
}

const (
	CreateAverageSalary = `
        INSERT INTO public.average_salary (date, value, id_table)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	DeleteByIdTableAverageSalary = `
		DELETE FROM public.average_salary WHERE id_table = $1
	`
)
