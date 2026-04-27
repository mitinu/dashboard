package SQL

const (
	CreateAverageSalary = `
        INSERT INTO public.average_salary (date, value, id_table)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	DeleteByIdTableAverageSalary = `
		DELETE FROM public.average_salary WHERE id_table = $1
	`

	GetAverageSalaryByDateRange = `
       SELECT id, date, value, id_table 
       FROM public.average_salary 
       WHERE date BETWEEN $1 AND $2
       ORDER BY date ASC
    `
)
