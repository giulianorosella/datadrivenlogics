package main

import (
	"log"

	"github.com/giulianorosella/ddlogic/pkg/config"
	"github.com/giulianorosella/ddlogic/pkg/excel"
	"github.com/giulianorosella/ddlogic/pkg/prover9/prove"
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

	sheetPath, formulaPrefix, formulaSuffix := cfg.SheetPath, cfg.FormulasPrefix, cfg.FormulasSuffix

	formulas, err := excel.ParseFormulasFromXLSX(sheetPath, formulaPrefix, formulaSuffix)

	if err != nil {
		log.Fatalf("Error when loading formulas: %s", err)

	}

	log.Printf("Formulas loaded, printing first ten\n")

	for i := 0; i < len(formulas) && i < 10; i++ {
		log.Printf("Index %d: %s\n", i, formulas[i].Expression)
	}

	modelFilePath, confInputFilePath, cmdPath, tempOutPath, tempInputPath, lineIndex := cfg.P9InputTemplate, cfg.P9DefaultInput, cfg.P9Exe, cfg.P9OutpuDir, cfg.P9InputDir, cfg.P9InputTemplateFormulaIndex

	res, err := prove.LogOut(modelFilePath, confInputFilePath, cmdPath, tempOutPath, tempInputPath, lineIndex, formulas[0].Expression)
	if err != nil {
		log.Fatalln("error when trying to log prover9: ", err)
	}

	log.Println("Result fo calling prover9: ", res)
}
