package config

import (
	"backend/src/pkg/logger"
	"fmt"
	"log"
	"os"
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

type Config struct {
	GinMode string
	AppPort string
	Server  ServerConfig
	// Reading
	PathExcel         string
	IntervalDaysReads string
	CronTime          string
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSslMode  string
}

func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		logger.Warn.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		GinMode: getEnv("GIN_MODE", "debug"),
		AppPort: getEnv("APP_PORT", "8080"),

		PathExcel:         getEnv("PATH_EXCEL", "resource/excel"),
		IntervalDaysReads: getEnv("INTERVAL_DAYS_READS", "1"),
		CronTime:          getEnv("CRON_TIME", "00:00"),
		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "postgres"),
		DBSslMode:  getEnv("DB_SSL_MODE", "disable"),
	}

	// Настройка таймаутов сервера
	cfg.Server = ServerConfig{
		ReadTimeout:     getDurationEnv("SERVER_READ_TIMEOUT", 15*time.Second),
		WriteTimeout:    getDurationEnv("SERVER_WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:     getDurationEnv("SERVER_IDLE_TIMEOUT", 60*time.Second),
		ShutdownTimeout: getDurationEnv("SERVER_SHUTDOWN_TIMEOUT", 10*time.Second),
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
