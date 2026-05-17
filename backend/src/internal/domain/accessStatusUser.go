package domain

type AccessStatusUser interface {
	GetIdByTitle(title string) (int64, error)
}
