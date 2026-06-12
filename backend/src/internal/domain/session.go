package domain

type Session interface {
	Create(userId int64, token string) error
}
