package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	CSVFilePath         string `json:"csvFilePath"`
	CsvComma            string `json:"csvComma"`
	FormulasChunkLength int    `json:"formulasChunkLength"`
	FormulasPrefix      string `json:"formulasPrefix"`
	FormulasSuffix      string `json:"formulasSuffix"`
}

func LoadConfig(configPath string) (Config, error) {
	var confs Config
	confsFile, err := os.Open(configPath)

	if err != nil {
		return confs, err
	}

	defer confsFile.Close()

	jsonParser := json.NewDecoder(confsFile)

	if err := jsonParser.Decode(&confs); err != nil {
		return confs, err
	}

	return confs, nil
}
