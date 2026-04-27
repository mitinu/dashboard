package SQL

const (
	CreateLaborMarket = `
        INSERT INTO public.labor_market (date, number_unemployed, number_employed, id_table)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	DeleteByIdTableLaborMarket = `
		DELETE FROM public.labor_market WHERE id_table = $1
	`

	GetLaborMarketByDateRange = `
       SELECT id, date, number_unemployed, number_employed, id_table 
       FROM public.labor_market 
       WHERE date BETWEEN $1 AND $2
       ORDER BY date ASC
    `
)
