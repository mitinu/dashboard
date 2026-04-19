package DTO

type TypeExcelInfo struct {
	ID    int64
	title string
}

const (
	GetIdByTitleTypeExcelInfo = `
		SELECT id FROM public.type_excel_file WHERE title = $1
	`
)
