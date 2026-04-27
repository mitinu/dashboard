package SQL

const (
	GetIdByTitleTypeExcelInfo = `
		SELECT id FROM public.type_excel_file WHERE title = $1
	`
)
