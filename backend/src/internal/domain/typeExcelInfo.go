package domain

type TypeExcelInfo interface {
	GetIdByTitle(title string) (int64, error)
}
