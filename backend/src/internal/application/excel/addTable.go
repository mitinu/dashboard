package excel

import (
	"backend/src/internal/DTO"
	"backend/src/internal/application/excel/reading"
	"backend/src/pkg/logger"
	"path/filepath"
)

type Item struct {
	IdTable int64
}

//func addTable(s *ExcelService, excelInfo *DTO.ExcelInfo, path string, name string, hash []byte, excelTypeRu string) {
//	excelInfo, err := saveExcelInfo(s, name, hash, excelTypeRu)
//	if err != nil {
//		logger.Error.Println(err)
//		return
//	}
//
//	var data any
//	switch excelTypeRu {
//	case "рынок труда":
//		data, err = reading.ReadingLaborMarket(filepath.Join(path, name))
//	case "ИПЦ":
//		data, err = reading.ReadingCPI(filepath.Join(path, name))
//	case "заработная плата":
//		data, err = reading.ReadingAverageSalary(filepath.Join(path, name))
//	default:
//		logger.Error.Printf("unknown field: %s", excelTypeRu)
//		return
//	}
//	if err != nil {
//		logger.Error.Printf("Error reading %s data: %v", excelTypeRu, err)
//		return
//	}
//	items := data.([]Item)
//	for i := range items {
//		items[i].IdTable = excelInfo.ID
//	}
//	value, err := s.Call(excelTypeRu)
//	if err != nil {
//		// handle error
//	}
//	table := value.(domain.TableDomain)
//
//	_, err = table.DeleteByIdTable(excelInfo.ID)
//	if err != nil {
//		logger.Error.Println(err)
//	}
//
//	errs := table.Creates(&items)
//	for _, err := range errs {
//		logger.Error.Println(err)
//	}
//}

func addTableCPI(s *ExcelService, excelInfo *DTO.ExcelInfo, path string, name string, hash []byte) {
	excelInfo, err := saveExcelInfo(s, name, hash, "ИПЦ")
	if err != nil {
		logger.Error.Println(err)
		return
	}

	cpiData, err := reading.ReadingCPI(filepath.Join(path, name))
	if err != nil {
		logger.Error.Println("Error reading CPI data: ", err)
		return
	}
	for i := range cpiData {
		cpiData[i].IdTable = excelInfo.ID
	}

	_, err = s.CPI.DeleteByIdTable(excelInfo.ID)
	if err != nil {
		logger.Error.Println(err)
	}

	errs := s.CPI.Creates(&cpiData)
	for _, err := range errs {
		logger.Error.Println(err)
	}
}
func addTableLaborMarket(s *ExcelService, excelInfo *DTO.ExcelInfo, path string, name string, hash []byte) {
	excelInfo, err := saveExcelInfo(s, name, hash, "рынок труда")
	if err != nil {
		logger.Error.Println(err)
		return
	}
	laborMarketData, err := reading.ReadingLaborMarket(filepath.Join(path, name))
	if err != nil {
		logger.Error.Println("Error reading CPI data: ", err)
		return
	}
	for i := range laborMarketData {
		laborMarketData[i].IdTable = excelInfo.ID
	}

	_, err = s.LaborMarket.DeleteByIdTable(excelInfo.ID)
	if err != nil {
		logger.Error.Println(err)
	}

	errs := s.LaborMarket.Creates(&laborMarketData)
	for _, err := range errs {
		logger.Error.Println(err)
	}
}

func addTableAverageSalary(s *ExcelService, excelInfo *DTO.ExcelInfo, path string, name string, hash []byte) {
	excelInfo, err := saveExcelInfo(s, name, hash, "заработная плата")
	if err != nil {
		logger.Error.Println(err)
		return
	}
	averageSalaryData, err := reading.ReadingAverageSalary(filepath.Join(path, name))
	if err != nil {
		logger.Error.Println("Error reading CPI data: ", err)
		return
	}
	for i := range averageSalaryData {
		averageSalaryData[i].IdTable = excelInfo.ID
	}
	_, err = s.AverageSalary.DeleteByIdTable(excelInfo.ID)
	if err != nil {
		logger.Error.Println(err)
	}

	errs := s.AverageSalary.Creates(&averageSalaryData)
	for _, err := range errs {
		logger.Error.Println(err)
	}
}
