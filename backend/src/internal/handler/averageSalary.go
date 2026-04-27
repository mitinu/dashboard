package handler

import (
	"backend/src/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAverageSalaryHandler(repo domain.AverageSalary) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := c.Query("start")
		end := c.Query("end")

		if start == "" || end == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Параметры start и end обязательны"})
			return
		}

		data, err := repo.GetByDateRange(start, end)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
			return
		}

		if len(data) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Данные не найдены"})
			return
		}

		// Возвращаем данные в формате "YYYY-MM-DD": value
		result := make(map[string]int)

		for _, item := range data {
			// Format("2006-01-02") — это стандарт ISO для дат без времени
			dateStr := item.Date.Format("2006-01-02")
			result[dateStr] = item.Value
		}

		c.JSON(http.StatusOK, result)
	}
}
