package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/giulianorosella/ddlogic/pkg/models"
)

func CreateCSV(fileName string, formulas []models.Formula) error {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Something went wrong when creating csv")
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Expression", "IsClassicTh", "IsIntuitionistTh"}
	if err := writer.Write(header); err != nil {
		log.Fatal("Something went wrong when creating csv")
		return err
	}

	for _, f := range formulas {
		record := []string{f.Expression, strconv.Itoa(int(f.IsClassicTh)), strconv.Itoa(int(f.IsIntuitionistTh))}
		if err := writer.Write(record); err != nil {
			log.Fatal("Something went wrong when creating csv")
			return err

		}
	}
	return nil
}
