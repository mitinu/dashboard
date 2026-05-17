package SQL

const (
	CreateUser = `
		INSERT INTO public.user (access_status_id, login, password_hash)
		VALUES ($1, $2, $3)
        RETURNING id
	`
	DeleteSuperadmin = `
		DELETE FROM public.user 
		WHERE access_status_id = (SELECT id FROM public.access_status_user WHERE title = $1)
	`
)
