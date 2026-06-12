package config

import (
	"backend/src/pkg/logger"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}
type Argon2Config struct {
	Memory      uint32 //MB
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

type Config struct {
	GinMode string
	AppPort string
	Server  ServerConfig
	// Reading
	PathExcel         string
	IntervalDaysReads int
	CronTime          string
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSslMode  string
	// Hash
	Argon2Config     Argon2Config
	Pepper           string
	SizeRefreshToken int
	DurationJWT      time.Duration
	SecretKeyJWT     []byte
	// Superadmin
	Login           string
	Password        string
	TitleSadminInDB string
}

func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		logger.Warn.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		GinMode: getEnv("GIN_MODE", "debug"),
		AppPort: getEnv("APP_PORT", "8080"),

		PathExcel:         getEnv("PATH_EXCEL", "resource/excel"),
		IntervalDaysReads: getIntEnv("INTERVAL_DAYS_READS", 1),
		CronTime:          getEnv("CRON_TIME", "00:00"),
		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "postgres"),
		DBSslMode:  getEnv("DB_SSL_MODE", "disable"),
		// Hash
		Pepper:           getEnv("PEPPER", "pizze"),
		SizeRefreshToken: getIntEnv("SIZE_REFRESH_TOKEN", 32),
		DurationJWT:      getDurationEnv("DURATION_JWT", 24*time.Hour),
		SecretKeyJWT:     []byte(getEnv("SECRET_KEY_JWT", "pizze")),
		// Superadmin
		Login:           getEnv("LOGIN", "root"),
		Password:        getEnv("PASSWORD", "root"),
		TitleSadminInDB: getEnv("TITLE_SADMIN_IN_DB", "суперадмин"),
	}
	// Настройка таймаутов сервера
	cfg.Server = ServerConfig{
		ReadTimeout:     getDurationEnv("SERVER_READ_TIMEOUT", 15*time.Second),
		WriteTimeout:    getDurationEnv("SERVER_WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:     getDurationEnv("SERVER_IDLE_TIMEOUT", 60*time.Second),
		ShutdownTimeout: getDurationEnv("SERVER_SHUTDOWN_TIMEOUT", 10*time.Second),
	}
	cfg.Argon2Config = Argon2Config{
		Memory:      uint32(getIntEnv("MEMORY", 64) * 1024),
		Iterations:  uint32(getIntEnv("ITERATIONS", 3)),
		Parallelism: uint8(getIntEnv("PARALLELISM", 4)),
		SaltLength:  uint32(getIntEnv("SALT_LENGTH", 16)),
		KeyLength:   uint32(getIntEnv("KEY_LENGTH", 32)),
	}

	// Устанавливаем режим Gin после загрузки конфигурации
	switch cfg.GinMode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "debug":
		fallthrough
	default:
		gin.SetMode(gin.DebugMode)
	}

	return cfg, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSslMode,
	)
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getDurationEnv возвращает duration из переменной окружения
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
		log.Printf("WARN: Invalid duration format for %s, using default: %v", key, defaultValue)
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return intValue
		}
		log.Printf("WARN: Invalid int format for %s, using default: %d", key, defaultValue)
	}
	return defaultValue
}
