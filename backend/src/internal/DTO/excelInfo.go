package DTO

import (
	"time"
)

type ExcelInfo struct {
	ID          int64
	Name        string
	Hash        []byte
	IdType      int64
	DateUpdated time.Time
}
