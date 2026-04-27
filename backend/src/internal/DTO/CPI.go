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
