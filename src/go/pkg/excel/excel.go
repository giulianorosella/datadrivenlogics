package excel

import (
	"errors"
	"strings"

	"github.com/tealeg/xlsx"
)

func ParseFormulasFromXLSX(xlsxPath string, formulaPrefix string, formulaSuffix string) ([]string, error) {
	// open file
	xlFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		return nil, err
	}

	formulas := make([]string, 0)

	// iterate on cells
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				// get cell
				cellValue := cell.String()

				// check cell
				if strings.HasPrefix(cellValue, formulaPrefix) && strings.HasSuffix(cellValue, formulaSuffix) {
					formulas = append(formulas, cellValue)
				}
			}
		}
	}

	if len(formulas) == 0 {
		return nil, errors.New("no formulas found")
	}

	return formulas, nil
}
