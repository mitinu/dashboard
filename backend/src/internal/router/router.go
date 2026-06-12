package router

import (
	"backend/src/internal/application/middleware"
	"backend/src/internal/config"
	"backend/src/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	averageSalaryHandler := handler.AverageSalaryHandler{}

	laborMarketHandler := handler.LaborMarketHandler{}

	api := router.Group("/api")
	api.Use(middleware.UserMiddleware(cfg))
	{
		api.GET("/unemployed", laborMarketHandler.GetUnemployedByDate())
		api.GET("/UnemployedPercentage", laborMarketHandler.GetUnemployedPercentageByDate())
		api.GET("/average-salary", averageSalaryHandler.Get())
	}

	return router
}
