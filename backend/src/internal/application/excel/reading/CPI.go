package reading

import (
	"backend/src/internal/DTO"
	"backend/src/pkg/logger"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ReadingCPI(path string) ([]DTO.CPI, error) {
	// Открываем Excel файл
	f, err := reading(path)
	defer f.Close()
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия файла: %w", err)
	}
	rows, err := readSheetData(f)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия листа: %w", err)
	}

	// Парсим заголовки годов (строка 2, индексы B, C, D)
	// rows[1] - это вторая строка (индекс 1), где находятся 2023, 2024, 2025
	years := make([]int, 0)
	for col := 1; col < len(rows[1]); col++ { // Идем по всем существующим колонкам
		yearStr := strings.TrimSpace(rows[1][col])
		if yearStr == "" {
			break // Прерываем цикл при первой пустой ячейке
		}

		year, err := strconv.Atoi(yearStr)
		if err != nil {
			logger.Error.Printf("Ошибка парсинга года в колонке %d: %s\n", col, yearStr)
			continue // Пропускаем ошибочные значения, продолжаем дальше
		}

		years = append(years, year)
	}

	if len(years) == 0 {
		return nil, fmt.Errorf("Годы не найдены в файле")
	}

	months := []string{
		"январь", "февраль", "март", "апрель", "май", "июнь",
		"июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь",
	}
	monthNumbers := map[string]time.Month{
		"январь": time.January, "февраль": time.February, "март": time.March,
		"апрель": time.April, "май": time.May, "июнь": time.June,
		"июль": time.July, "август": time.August, "сентябрь": time.September,
		"октябрь": time.October, "ноябрь": time.November, "декабрь": time.December,
	}
	var cpiData []DTO.CPI

	for i := 2; i < len(rows) && i-2 < len(months); i++ {
		row := rows[i]
		if len(row) < 4 {
			continue
		}
	}
	for i := 2; i < len(rows) && i-2 < len(months); i++ {
		row := rows[i]
		if len(row) < 4 {
			continue
		}

		monthName := strings.TrimSpace(strings.ToLower(row[0]))
		month, exists := monthNumbers[monthName]
		if !exists {
			continue
		}

		for colIdx, year := range years {
			if colIdx+1 >= len(row) {
				continue
			}

			valueStr := strings.TrimSpace(row[colIdx+1])
			if valueStr == "" {
				continue
			}

			valueFloat, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				continue
			}

			date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
			cpiData = append(cpiData, DTO.CPI{
				Date:  date,
				Value: valueFloat,
			})
		}
	}

	return cpiData, nil
}
