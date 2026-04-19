package reading

import (
	"backend/src/internal/DTO"
	"backend/src/pkg/logger"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ReadingSalary читает данные о среднемесячной заработной плате из Excel файла
func ReadingAverageSalary(path string) ([]DTO.AverageSalary, error) {
	// Открываем Excel файл
	f, err := reading(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer f.Close()

	rows, err := readSheetData(f)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения данных листа: %w", err)
	}

	var salaryData []DTO.AverageSalary

	// Пропускаем заголовок (первая строка)
	// Начинаем со второй строки (индекс 1)
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 2 {
			continue // Пропускаем строки с недостаточным количеством колонок
		}

		// Получаем период (например "1 кв. 2023 г.")
		periodStr := strings.TrimSpace(row[0])
		if periodStr == "" {
			continue
		}

		// Получаем значение зарплаты
		salaryStr := strings.TrimSpace(row[1])
		if salaryStr == "" {
			continue
		}

		// Парсим зарплату в int
		salary, err := strconv.Atoi(salaryStr)
		if err != nil {
			logger.Error.Printf("Ошибка парсинга зарплаты в строке %d: %s\n", i+1, salaryStr)
			continue
		}
		// Парсим период и получаем дату (последний месяц квартала)
		date, err := parseQuarterPeriod(periodStr)
		if err != nil {
			logger.Error.Printf("Ошибка парсинга периода в строке %d: %s - %v\n", i+1, periodStr, err)
			continue
		}

		salaryData = append(salaryData, DTO.AverageSalary{
			Date:  date,
			Value: salary,
		})
	}

	if len(salaryData) == 0 {
		return nil, fmt.Errorf("не найдено данных о заработной плате в файле")
	}

	return salaryData, nil
}

// parseQuarterPeriod парсит строку периода вида "1 кв. 2023 г." и возвращает дату последнего месяца квартала
func parseQuarterPeriod(periodStr string) (time.Time, error) {
	// Убираем лишние пробелы и приводим к нижнему регистру
	periodStr = strings.TrimSpace(strings.ToLower(periodStr))

	// Разбиваем строку на части
	parts := strings.Fields(periodStr)
	if len(parts) < 4 {
		return time.Time{}, fmt.Errorf("неверный формат периода: %s", periodStr)
	}

	// Парсим номер квартала (убираем "кв.")
	quarterStr := strings.TrimSuffix(parts[0], "кв.")
	quarter, err := strconv.Atoi(quarterStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("ошибка парсинга номера квартала: %s", quarterStr)
	}

	if quarter < 1 || quarter > 4 {
		return time.Time{}, fmt.Errorf("неверный номер квартала: %d", quarter)
	}

	// Парсим год (убираем "г.")
	yearStr := strings.TrimSuffix(parts[2], "г.")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("ошибка парсинга года: %s", yearStr)
	}

	// Определяем последний месяц квартала
	var month time.Month
	switch quarter {
	case 1:
		month = time.March
	case 2:
		month = time.June
	case 3:
		month = time.September
	case 4:
		month = time.December
	}

	// Возвращаем первый день последнего месяца квартала
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC), nil
}
