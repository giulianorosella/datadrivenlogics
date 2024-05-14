package main

import (
	"log"
	"os"
	"sync"

	"github.com/giulianorosella/ddlogic/pkg/config"
	"github.com/giulianorosella/ddlogic/pkg/csv"
	"github.com/giulianorosella/ddlogic/pkg/generator"
	"github.com/giulianorosella/ddlogic/pkg/models"
	"github.com/giulianorosella/ddlogic/pkg/prover9"
)

func main() {

	env := os.Getenv("ENV")

	cfgPath := "../config/config.json"
	if env == "dev" {
		cfgPath = "../config/config_dev.json"
	}

	log.Print("Loading configurations \n")

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Error when loading configurations: %s", err)
	}

	log.Printf("Config loaded: %+v\n", cfg)

	log.Printf("Loading formulas fr z3\n")

	formulas, err := generator.ProveClassic(cfg.VarNum, cfg.ConNum)

	if err != nil {
		log.Fatalf("Something went wrong when proving for z3: %s", err)
	}

	log.Print("Preparing strings for Prover9")
	p9Strings := generator.P9FormulasCreationInit(cfg.VarNum, cfg.ConNum)

	for i, p9String := range p9Strings {
		formulas[i].Expression = p9String
	}

	log.Printf("Formulas loaded, printing first ten\n")

	for i := 0; i < len(formulas) && i < 10; i++ {
		log.Printf("Index %d: %s\n", i, formulas[i].Expression)
	}

	modelFilePath, confInputFilePath, p9ExePath, m4Exepath, tempOutPath, tempInputPath, lineIndex := cfg.P9InputTemplate, cfg.P9DefaultInput, cfg.P9Exe, cfg.Mace4Exe, cfg.P9OutpuDir, cfg.P9InputDir, cfg.P9InputTemplateFormulaIndex

	if err != nil {
		log.Fatalln("error when trying to log prover9: ", err)
	}

	var wg sync.WaitGroup
	maxWorkers := cfg.FormulasChunkLength
	semaphore := make(chan struct{}, maxWorkers)
	var res []models.Formula
	for i, formula := range formulas {
		wg.Add(1)
		log.Println("start formulas in array: ", i)

		go func(formula models.Formula) {

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			defer wg.Done()

			if formula.IsClassicTh == models.True {

				formula, err = prover9.GetProofs(modelFilePath, confInputFilePath, p9ExePath, m4Exepath, tempOutPath, tempInputPath, lineIndex, formula)
				if err != nil {
					log.Println("error when trying to log prover9: ", err)
					return
				}
			} else {
				formula.IsIntuitionistTh = models.False
			}
			log.Println("just updated in res ", i)
			res = append(res, formula)

		}(formula)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	log.Println(res)
	csv.CreateCSV(cfg.CsvFileName, res)
}
