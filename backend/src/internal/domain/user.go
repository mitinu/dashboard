package domain

import "backend/src/internal/DTO"

type User interface {
	Create(user *DTO.User) error
	GetByLogin(login string) (*DTO.User, error)
	DeleteSuperadmin() error
}
