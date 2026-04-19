package reading

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func reading(fileName string) (*excelize.File, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		// Исправлено: добавлена запятая между аргументами и правильный формат ошибки
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	return f, nil
}

func readSheetData(f *excelize.File) ([][]string, error) {
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("нет листов")
	}
	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения строк: %w", err)
	}

	if len(rows) < 3 {
		return nil, fmt.Errorf("недостаточно данных в файле: получено %d строк, требуется минимум 3", len(rows))
	}

	return rows, nil
}
