package csv

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/giulianorosella/ddlogic/pkg/models"
	"github.com/giulianorosella/ddlogic/pkg/utils"
)

func ParseFormulasFromCSV(csvPath string, formulaPrefix string, formulaSuffix string, delimiter string) ([]models.Formula, error) {
	//open file
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// set up reader
	reader := csv.NewReader(bufio.NewReader(file))

	d, err := utils.StringToRune(delimiter)
	if err != nil {
		return nil, err
	}
	reader.Comma = d

	formulas := []models.Formula{}

	// parse excell till EOF
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if len(record) > 0 {
			cell := record[0]
			cell = strings.TrimSpace(cell)
			if strings.HasPrefix(cell, formulaPrefix) && strings.HasSuffix(cell, formulaSuffix) {
				formulas = append(formulas, models.Formula{Expression: cell, IsClassicTh: models.Unset, IsIntuitionistTh: models.Unset})
			}
		}
	}
	return formulas, nil
}
