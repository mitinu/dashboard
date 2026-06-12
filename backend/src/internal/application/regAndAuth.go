package application

import (
	"backend/src/internal/DTO"
	"backend/src/internal/config"
	"backend/src/internal/domain"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ReqAndAutoService struct {
	User domain.User
	cfg  *config.Config
}

func NewReqAndAutoService(user domain.User, cfg *config.Config) *ReqAndAutoService {
	return &ReqAndAutoService{
		User: user,
		cfg:  cfg,
	}
}

func (s ReqAndAutoService) Register(userReq DTO.UserReq) error {
	passwordService := NewPasswordService(s.cfg)
	hashedPassword, err := passwordService.Hash(userReq.Password)
	if err != nil {
		return fmt.Errorf("failed to secure password: %v", err)
	}

	user := &DTO.User{
		Login:          userReq.Login,
		PasswordHash:   hashedPassword,
		AccessStatusId: userReq.AccessStatusId,
	}

	if err := s.User.Create(user); err != nil {
		return fmt.Errorf("Could not create user: %v", err)
	}
	return nil
}

func (s ReqAndAutoService) CreateRefreshToken() (string, error) {
	b := make([]byte, s.cfg.SizeRefreshToken)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate secure random bytes: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

type UserClaims struct {
	ID             int64
	AccessStatusId int64
	jwt.RegisteredClaims
}

func (s ReqAndAutoService) CreateJwtToken(userID int64, AccessStatusId int64) (string, error) {
	claims := UserClaims{
		ID:             userID,
		AccessStatusId: AccessStatusId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.DurationJWT)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	signedToken, err := token.SignedString(s.cfg.SecretKeyJWT)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}
