package SQL

const (
	СreteSession = `
		INSERT INTO public.sesion (user_id, token, expires_at)
		VALUES ($1, $2, CURRENT_DATE + INTERVAL '30 days')
	`
)
