package SQL

const (
	GetIdByTitle = `
		SELECT id
		FROM public.access_status_user
		WHERE title = $1
	`
)
