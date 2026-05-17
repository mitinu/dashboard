package DTO

type User struct {
	ID             int64
	Login          string
	PasswordHash   string
	AccessStatusId int64
}
