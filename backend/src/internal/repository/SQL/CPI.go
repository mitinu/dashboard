package SQL

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
