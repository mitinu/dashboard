package application

import (
	"backend/src/internal/config"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type PasswordService struct {
	pepper []byte
	config config.Argon2Config
}

func NewPasswordService(cfg *config.Config) *PasswordService {
	return &PasswordService{
		pepper: []byte(cfg.Pepper),
		config: cfg.Argon2Config,
	}
}
func (s *PasswordService) Hash(password string) (string, error) {
	salt := make([]byte, s.config.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	pepperedPassword := append([]byte(password), s.pepper...)

	hash := argon2.IDKey(pepperedPassword, salt, s.config.Iterations, s.config.Memory, s.config.Parallelism, s.config.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, s.config.Memory, s.config.Iterations, s.config.Parallelism, b64Salt, b64Hash)

	return encoded, nil
}
func (s *PasswordService) Verify(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil || version != argon2.Version {
		return false, errors.New("incompatible version")
	}

	var memory, iterations uint32
	var parallelism uint8
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism); err != nil {
		return false, errors.New("invalid parameters")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	pepperedPassword := append([]byte(password), s.pepper...)
	comparisonHash := argon2.IDKey(pepperedPassword, salt, iterations, memory, parallelism, uint32(len(decodedHash)))

	if subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1 {
		return true, nil
	}
	return false, nil
}
