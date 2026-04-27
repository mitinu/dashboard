package router

import (
	"backend/src/internal/domain"
	"backend/src/internal/handler"

	"github.com/gin-gonic/gin"
)

type Repositories struct {
	LaborMarket   domain.LaborMarket
	AverageSalary domain.AverageSalary
}

func SetupRouter(repos Repositories) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/unemployed", handler.GetUnemployedByDate(repos.LaborMarket))
		api.GET("/UnemployedPercentage", handler.GetUnemployedPercentageByDate(repos.LaborMarket))
		api.GET("/average-salary", handler.GetAverageSalaryHandler(repos.AverageSalary))
	}

	return router
}
