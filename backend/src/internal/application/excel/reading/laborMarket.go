package reading

import (
	"backend/src/internal/DTO"
	"backend/src/pkg/logger"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ReadingLaborMarket читает данные о зарегистрированном рынке труда из Excel файла
func ReadingLaborMarket(path string) ([]DTO.LaborMarket, error) {
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

	var laborMarketData []DTO.LaborMarket

	// Пропускаем заголовки (первые 2 строки)
	// Строка 0: "Зарегистрированный рынок труда"
	// Строка 1: "Численность безработных, чел." | "Численность трудоустроенных, чел."
	// Данные начинаются со строки 2 (индекс 2)

	for i := 2; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 3 {
			continue // Пропускаем строки с недостаточным количеством колонок
		}

		// Получаем дату
		dateStr := strings.TrimSpace(row[0])
		if dateStr == "" {
			continue
		}

		// Парсим дату
		date, err := parseDate(dateStr)
		if err != nil {
			logger.Error.Printf("Ошибка парсинга даты в строке %d: %s - %v\n", i+1, dateStr, err)
			continue
		}

		// Получаем численность безработных
		unemployedStr := strings.TrimSpace(row[1])
		var unemployed int
		if unemployedStr != "" && unemployedStr != "-" {
			unemployed, err = parseInt(unemployedStr)
			if err != nil {
				logger.Error.Printf("Ошибка парсинга численности безработных в строке %d: %s - %v\n", i+1, unemployedStr, err)
				continue
			}
		}

		// Получаем численность трудоустроенных
		employedStr := strings.TrimSpace(row[2])
		var employed int
		if employedStr != "" && employedStr != "-" {
			employed, err = parseInt(employedStr)
			if err != nil {
				logger.Error.Printf("Ошибка парсинга численности трудоустроенных в строке %d: %s - %v\n", i+1, employedStr, err)
				continue
			}
		}

		laborMarketData = append(laborMarketData, DTO.LaborMarket{
			Date:       date,
			Unemployed: unemployed,
			Employed:   employed,
		})
	}

	if len(laborMarketData) == 0 {
		return nil, fmt.Errorf("не найдено данных о рынке труда в файле")
	}

	return laborMarketData, nil
}

// parseDate парсит дату из строки
func parseDate(dateStr string) (time.Time, error) {
	// Пробуем разные форматы даты
	formats := []string{
		"Jan-06",               // Dec-20 (месяц-год)
		"Jan-2006",             // Dec-2020 (месяц-год с полным годом)
		"2006-01-02 15:04:05",  // 2020-12-01 00:00:00
		"2006-01-02",           // 2020-12-01
		"02.01.2006",           // 01.12.2020
		"2006-01-02T15:04:05Z", // ISO формат
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			// Для форматов "Jan-06" и "Jan-2006" устанавливаем первый день месяца
			if format == "Jan-06" || format == "Jan-2006" {
				// Создаем новую дату с первым числом месяца
				return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()), nil
			}
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("не удалось распарсить дату: %s", dateStr)
}

// parseInt парсит целое число из строки, удаляя запятые и пробелы
func parseInt(s string) (int, error) {
	// Удаляем все запятые и пробелы
	cleaned := strings.ReplaceAll(s, ",", "")
	cleaned = strings.ReplaceAll(cleaned, " ", "")
	cleaned = strings.TrimSpace(cleaned)

	// Парсим как целое число
	return strconv.Atoi(cleaned)
}
