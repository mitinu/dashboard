package DTO

import "time"

type LaborMarket struct {
	ID         int64
	Date       time.Time
	Unemployed int
	Employed   int
	IdTable    int64
}
