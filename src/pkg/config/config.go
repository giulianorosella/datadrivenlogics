package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	FormulasChunkLength               int    `json:"formulasChunkLength"`
	VarNum                            int    `json:"varNum"`
	ConNum                            int    `json:"conNum"`
	CsvFileName                       string `json:"csvFileName"`
	P9Exe                             string `json:"p9Exe"`
	Mace4Exe                          string `json:"mace4Exe"`
	P9DefaultInput                    string `json:"p9DefaultInput"`
	P9InputDir                        string `json:"p9InputDir"`
	P9OutpuDir                        string `json:"p9OutputDir"`
	P9InputTemplate                   string `json:"p9InputTemplate"`
	P9InputTemplateFormulaIndex       int    `json:"p9InputTemplateFormulaNumber"`
	P9InputTemplateFormulaPlaceholder string `json:"p9InputTemplateFormulaPlaceHolder"`
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
