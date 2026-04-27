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
