package main

import (
	"log"

	"github.com/giulianorosella/ddlogic/pkg/config"
	"github.com/giulianorosella/ddlogic/pkg/excel"
)

func main() {
	cfgPath := "../config/config.json"

	log.Print("Loading configurations \n")

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Error when loading configurations: %s", err)
	}

	log.Printf("Config loaded: %+v\n", cfg)

	log.Printf("Loading formulas\n")

	csvPath, formulaPrefix, formulaSuffix, delimiter := cfg.CSVFilePath, cfg.FormulasPrefix, cfg.FormulasSuffix, cfg.CsvComma

	formulas, err := excel.ParseFormulasFromCSV(csvPath, formulaPrefix, formulaSuffix, delimiter)

	if err != nil {
		log.Fatalf("Error when loading formulas: %s", err)

	}

	log.Printf("Formulas loaded, printing first ten\n")

	for i := 0; i < len(formulas) && i < 10; i++ {
		log.Printf("Index %d: %s\n", i, formulas[i])
	}

}
