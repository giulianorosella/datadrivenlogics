package excel

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/giulianorosella/ddlogic/pkg/utils"
)

func ParseFormulasFromCSV(csvPath string, formulaPrefix string, formulaSuffix string, delimiter string) ([]string, error) {
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

	formulas := make([]string, 0)

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
				formulas = append(formulas, cell)
			}
		}
	}
	return formulas, nil
}
