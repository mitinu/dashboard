package application

import (
	"backend/src/internal/application/excel"
	"backend/src/internal/config"
	"backend/src/internal/repository/postgres"
	"backend/src/pkg/logger"
	"fmt"
	"strings"

	"github.com/robfig/cron/v3"
)

var excelService *excel.ExcelService

func PlannerInit(db *postgres.PostgresDB, cfg *config.Config) (*cron.Cron, error) {

	excelInfoRepo := postgres.NewExcelInfoRepository(db)
	typeExcelInfoRepo := postgres.NewTypeExcelInfoRepository(db)
	cpiRepo := postgres.NewCpiRepository(db)
	laborMarketRepo := postgres.NewLaborMarketRepository(db)
	averageSalaryRepo := postgres.NewAverageSalaryRepository(db)

	// Создание сервиса
	excelService = &excel.ExcelService{
		ExcelInfo:     excelInfoRepo,
		TypeExcelInfo: typeExcelInfoRepo,
		CPI:           cpiRepo,
		LaborMarket:   laborMarketRepo,
		AverageSalary: averageSalaryRepo,
	}

	c := cron.New(cron.WithSeconds())
	parts := strings.Split(cfg.CronTime, ":")
	hour, min := parts[0], parts[1]
	cronSpec := fmt.Sprintf("0 %s %s */%d * *", min, hour, cfg.IntervalDaysReads)

	_, err := c.AddFunc(cronSpec, func() { excelService.CheckHashes(cfg.PathExcel) })
	if err != nil {
		return nil, err
	}
	logger.Info.Printf("чтение excel каждые %d дней в %s", cfg.IntervalDaysReads, cfg.CronTime)
	return c, nil
}

func PlannerStart(c *cron.Cron, cfg *config.Config) {
	c.Start() // Запуск планировщика
	excelService.CheckHashes(cfg.PathExcel)
}
