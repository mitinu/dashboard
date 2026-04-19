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

const (
	GetByNameExcelInfo = `
		SELECT id, id_type, name, date_updated, hash 
	  	FROM public.excel_file WHERE name = $1
	`
	CreateOrUpdateExcelInfo = `
		INSERT INTO public.excel_file (name, hash, id_type, date_updated) 
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
		ON CONFLICT (name) DO UPDATE SET 
			hash = EXCLUDED.hash,
			id_type = EXCLUDED.id_type,
			date_updated = CURRENT_TIMESTAMP
		RETURNING id
	`
)
