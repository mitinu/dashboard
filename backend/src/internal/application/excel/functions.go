package excel

import (
	"backend/src/internal/DTO"
	"backend/src/pkg/logger"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getNameFiles(path string) []string {
	var names []string

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error.Println("Error accessing path:", path, err)
			return nil
		}

		if info != nil && !info.IsDir() {
			name := info.Name()
			if strings.HasSuffix(name, ".xlsx") || strings.HasSuffix(name, ".xls") {
				names = append(names, name)
			}
		}

		return nil
	})

	return names
}

func saveExcelInfo(s *ExcelService, name string, hash []byte, excelType string) (*DTO.ExcelInfo, error) {
	idType, err := s.TypeExcelInfo.GetIdByTitle(excelType)
	excelInfo := &DTO.ExcelInfo{
		Name:   name,
		Hash:   hash,
		IdType: idType,
	}
	err = s.ExcelInfo.CreateOrUpdate(excelInfo)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сохранении файла %s: %v", name, err)
	}
	return excelInfo, nil
}

func hasAnySubstring(s string, substrings []string) bool {
	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}
