package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/giulianorosella/ddlogic/pkg/config"
	persistance "github.com/giulianorosella/ddlogic/pkg/db"
	"github.com/giulianorosella/ddlogic/pkg/generator"
	"github.com/giulianorosella/ddlogic/pkg/models"
	"github.com/giulianorosella/ddlogic/pkg/prover9"
)

func main() {

	cfgPath := "/usr/local/bin/config/config.json"

	log.Print("Loading configurations \n")

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Error when loading configurations: %s", err)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	dbServer := os.Getenv("DB_SERVER")
	dbName := os.Getenv("DB_NAME")
	dbPort, portErr := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	isAzure := strings.ToLower(os.Getenv("IS_AZURE")) == "true" || os.Getenv("IS_AZURE") == "1"

	if dbPassword == "" || dbServer == "" || dbName == "" || portErr != nil || dbUser == "" {
		log.Fatalln("Could not retrive db secrets, disconnecting")
		return
	}

	log.Printf("Config loaded: %+v\n", cfg)

	log.Printf("Loading formulas fr z3\n")

	formulas, err := generator.ProveClassic(cfg.ConNum, cfg.VarNum)

	if err != nil {
		log.Fatalf("Something went wrong when proving for z3: %s", err)
	}

	log.Print("Preparing strings for Prover9")
	p9Strings := generator.P9FormulasCreationInit(cfg.ConNum, cfg.VarNum)

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

	db, err := persistance.InitDb(dbServer, dbUser, dbPassword, dbName, dbPort, isAzure)
	if err != nil {
		return
	}
	var inputsExists bool
	var isTableCreated bool

	if isAzure {
		inputsExists, err = persistance.InputsExistAzure(db, cfg.ConNum, cfg.VarNum)
		if err != nil {
			log.Println("InputExistsAzure returned error: ", err)
			return
		}
		isTableCreated, err = persistance.CreateTableAzure(db, cfg.ConNum, cfg.VarNum)

	} else {
		inputsExists, err = persistance.InputsExist(db, cfg.ConNum, cfg.VarNum)
		if err != nil {
			log.Println("InputExists returned error: ", err)
			return
		}
		isTableCreated, err = persistance.CreateTable(db, cfg.ConNum, cfg.VarNum)
	}

	if inputsExists {
		log.Fatal("Inputs exists already in db, aborting")
		return
	}

	if !isTableCreated || err != nil {
		log.Fatal("Could not create table: ", err)
	}

	log.Printf("Going to print %d formulas", len(formulas))

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

			formulaId, err := persistance.CreateFormula(db, formula, cfg.ConNum, cfg.VarNum, isAzure)
			if err != nil {
				log.Println("Could not update in db formula: ", formula.Expression, err.Error())
				return
			}
			log.Println("Added to db", formulaId)
		}(formula)
	}

	// wait for all goroutines to finish
	wg.Wait()
	log.Println(res)
	log.Println("Finished successfully")
}
