package excel

import (
	"backend/src/pkg/logger"
	"bytes"
	"path/filepath"
	"strings"
)

func (s *ExcelService) CheckHashes(path string) {
	names := getNameFiles(path)
	for i := range names {
		name := names[i]
		checkHash(s, path, name)
	}

}
func checkHash(s *ExcelService, path string, name string) {
	lowerName := strings.ToLower(name)
	substrings := []string{"ипц", "рынок труда", "заработная плата"}
	if hasAnySubstring(lowerName, substrings) {
		filePath := filepath.Join(path, name)
		hash, err := hashing(filePath)
		if err != nil {
			logger.Error.Println("не удалось хэшировать файл", err)
			return
		}
		excelInfo, err := s.ExcelInfo.GetByName(name)
		logger.Debug.Println(excelInfo)
		if err != nil || !bytes.Equal(hash, excelInfo.Hash) {
			switch {
			case strings.Contains(lowerName, "ипц"):
				addTableCPI(s, excelInfo, path, name, hash)
			case strings.Contains(lowerName, "рынок труда"):
				addTableLaborMarket(s, excelInfo, path, name, hash)
			case strings.Contains(lowerName, "заработная плата"):
				addTableAverageSalary(s, excelInfo, path, name, hash)
			default:
				logger.Warn.Println("сервер не определил файл", name)
			}
		}
		logger.Info.Printf("файл %s с ID: %d загружен/проверн успешно", name, excelInfo.ID)
	} else {
		logger.Warn.Println("сервер не определил файл", name)
	}
}
