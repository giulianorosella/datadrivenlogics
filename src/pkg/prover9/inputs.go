package prover9

import (
	"bufio"
	"os"

	"github.com/giulianorosella/ddlogic/pkg/utils"
)

func createTempInputFile(modelFilePath string, outputPath string, lineIndex int, formula string) (*os.File, error) {
	//open model file
	modelFile, err := os.Open(modelFilePath)
	if err != nil {
		return nil, err
	}

	defer modelFile.Close()

	//create scanner
	scanner := bufio.NewScanner(modelFile)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// modify and write to temp

	lines[lineIndex-1] = formula

	tempFile, err := os.CreateTemp(outputPath, "tempfile*.in")

	if err != nil {
		return nil, err
	}

	writer := bufio.NewWriter(tempFile)

	for _, line := range lines {
		_, err := writer.WriteString((line + "\n"))
		if err != nil {
			utils.CloseAndRemoveTempFile(tempFile)
			return nil, err
		}
	}

	writer.Flush()

	// reset the pointer to the beginning of the file
	_, err = utils.ResetFileOffSet(tempFile)
	if err != nil {
		utils.CloseAndRemoveTempFile(tempFile)
		return nil, err
	}

	return tempFile, nil

}
