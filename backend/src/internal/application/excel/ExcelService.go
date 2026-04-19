package excel

import "backend/src/internal/domain"

type ExcelService struct {
	ExcelInfo     domain.ExcelInfo
	TypeExcelInfo domain.TypeExcelInfo
	CPI           domain.CPI
	LaborMarket   domain.LaborMarket
	AverageSalary domain.AverageSalary
}

// рудимент от попытки унифицировать addTable
//func (s *ExcelService) Call(key string) (any, error) {
//	switch key {
//	case "рынок труда":
//		return s.LaborMarket, nil
//	case "ИПЦ":
//		return s.CPI, nil
//	case "заработная плата":
//		return s.AverageSalary, nil
//	default:
//		return nil, fmt.Errorf("unknown field: %s", key)
//	}
//}
