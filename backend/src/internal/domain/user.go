package domain

import "backend/src/internal/DTO"

type User interface {
	CreateUser(user *DTO.User) error
	DeleteSuperadmin() error
}
