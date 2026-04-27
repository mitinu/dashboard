package main

import (
	"backend/src/internal/application/excel"
	"backend/src/internal/config"
	"backend/src/internal/repository/postgres"
	"backend/src/internal/router"
	"backend/src/pkg/logger"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/robfig/cron/v3"
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

	excelInfoRepo := postgres.NewExcelInfoRepository(db)
	typeExcelInfoRepo := postgres.NewTypeExcelInfoRepository(db)
	cpiRepo := postgres.NewCpiRepository(db)
	laborMarketRepo := postgres.NewLaborMarketRepository(db)
	averageSalaryRepo := postgres.NewAverageSalaryRepository(db)

	// Создание сервиса
	excelService := &excel.ExcelService{
		ExcelInfo:     excelInfoRepo,
		TypeExcelInfo: typeExcelInfoRepo,
		CPI:           cpiRepo,
		LaborMarket:   laborMarketRepo,
		AverageSalary: averageSalaryRepo,
	}

	c := cron.New(cron.WithSeconds())
	parts := strings.Split(cfg.CronTime, ":")
	hour, min := parts[0], parts[1]
	cronSpec := fmt.Sprintf("0 %s %s */%s * *", min, hour, cfg.IntervalDaysReads)

	_, err = c.AddFunc(cronSpec, func() { excelService.CheckHashes(cfg.PathExcel) })
	if err != nil {
		logger.Error.Fatalf("Ошибка при добавлении задачи в cron: %v", err)
	}
	logger.Info.Printf("чтение excel каждые %s дней в %s", cfg.IntervalDaysReads, cfg.CronTime)

	c.Start() // Запуск планировщика
	defer c.Stop()

	logger.Debug.Println("testStart")
	excelService.CheckHashes(cfg.PathExcel)
	logger.Debug.Println("testEnd")

	r := router.SetupRouter(router.Repositories{
		LaborMarket:   postgres.NewLaborMarketRepository(db),
		AverageSalary: postgres.NewAverageSalaryRepository(db),
	})

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
