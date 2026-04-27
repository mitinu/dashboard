package handler

import (
	"backend/src/internal/domain"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUnemployedByDate(repo domain.LaborMarket) gin.HandlerFunc {
	return func(c *gin.Context) {
		startDate := c.Query("start")
		endDate := c.Query("end")

		if startDate == "" || endDate == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Параметры 'start' и 'end' обязательны"})
			return
		}

		data, err := repo.GetLaborMarketByDateRange(startDate, endDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
			return
		}

		if len(data) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Данные не найдены"})
			return
		}

		// Создаем карту (map), где ключом будет дата, а значением — количество безработных
		result := make(map[string]int)

		for _, item := range data {
			dateStr := item.Date.Format("2006-01-02")
			result[dateStr] = item.Unemployed
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetUnemployedPercentageByDate(repo domain.LaborMarket) gin.HandlerFunc {
	return func(c *gin.Context) {
		startDate := c.Query("start")
		endDate := c.Query("end")

		if startDate == "" || endDate == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Параметры 'start' и 'end' обязательны"})
			return
		}

		data, err := repo.GetLaborMarketByDateRange(startDate, endDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
			return
		}

		if len(data) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Данные не найдены"})
			return
		}

		// Используем float64 для хранения процентов
		result := make(map[string]float64)

		for _, item := range data {
			dateStr := item.Date.Format("2006-01-02")

			totalWorkforce := item.Unemployed + item.Employed

			if totalWorkforce > 0 {
				// Рассчитываем процент
				percentage := (float64(item.Unemployed) / float64(totalWorkforce)) * 100

				// Округляем до 2 знаков после запятой для красоты
				result[dateStr] = math.Round(percentage*100) / 100
			} else {
				// Если данных о людях нет вовсе, ставим 0
				result[dateStr] = 0
			}
		}

		c.JSON(http.StatusOK, result)
	}
}
