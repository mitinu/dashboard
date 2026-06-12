package SQL

const (
	CreateUser = `
		INSERT INTO public.user (access_status_id, login, password_hash)
		VALUES ($1, $2, $3)
        RETURNING id
	`
	GetUserByLogin = `
		SELECT id, username, email, password_hash, created_at 
		FROM users 
		WHERE username = $1 OR email = $1 
		LIMIT 1
	`
	DeleteUserSuperadmin = `
		DELETE FROM public.user 
		WHERE access_status_id = (SELECT id FROM public.access_status_user WHERE title = $1)
	`
)
