package main

import (
	"backend/src/internal/application"
	"backend/src/internal/config"
	"backend/src/internal/repository/postgres"
	"backend/src/internal/router"
	"backend/src/pkg/logger"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Подключаемся к базе данных
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	application.CreateSuperadmin(cfg, db)

	c, err := application.PlannerInit(db, cfg)
	if err != nil {
		logger.Error.Fatalf("Ошибка при добавлении задачи в cron: %v", err)
	}
	application.PlannerStart(c, cfg)
	defer c.Stop()

	r := router.SetupRouter(cfg)

	srv := &http.Server{
		Addr:           ":" + cfg.AppPort,
		Handler:        r,
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		IdleTimeout:    cfg.Server.IdleTimeout,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	// Запуск сервера в горутине для graceful shutdown
	go func() {
		logger.Info.Printf("Server starting on port %s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to run server:", err)
		}
	}()

	// Ожидание сигнала для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// Закрываем соединение с БД
	if err := db.Close(); err != nil {
		logger.Error.Printf("Error closing database connection: %v", err)
	}

	log.Println("Server exited properly")

	//	TODO HealthCheck
}
