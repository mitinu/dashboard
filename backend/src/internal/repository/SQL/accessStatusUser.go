package SQL

const (
	GetIdAccessStatusUserByTitle = `
		SELECT id
		FROM public.access_status_user
		WHERE title = $1
	`
)
