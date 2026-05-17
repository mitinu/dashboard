package application

import (
	"backend/src/internal/DTO"
	"backend/src/internal/config"
	"backend/src/internal/domain"
	"fmt"
)

type ReqAndAutoService struct {
	User domain.User
}

func NewReqAndAutoService(user domain.User) *ReqAndAutoService {
	return &ReqAndAutoService{
		User: user,
	}
}

func (s ReqAndAutoService) Register(userReq DTO.UserReq, cfg *config.Config) error {
	passwordService := NewPasswordService(cfg)
	hashedPassword, err := passwordService.Hash(userReq.Password)
	if err != nil {
		return fmt.Errorf("failed to secure password: %v", err)
	}

	user := &DTO.User{
		Login:          userReq.Login,
		PasswordHash:   hashedPassword,
		AccessStatusId: userReq.AccessStatusId,
	}

	if err := s.User.CreateUser(user); err != nil {
		return fmt.Errorf("Could not create user: %v", err)
	}
	return nil
}
