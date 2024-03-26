package prove

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/giulianorosella/ddlogic/pkg/prover9/inputs"
	"github.com/giulianorosella/ddlogic/pkg/utils"
)

func Prove(confInputFilePath string, inputTempFile *os.File, cmdPath string, tempOutPath string) (*os.File, error) {

	defer inputTempFile.Close()

	// execute prover9
	cmdArgs := []string{"-f", confInputFilePath, inputTempFile.Name()}
	cmd := exec.Command(cmdPath, cmdArgs...)

	tempOutFile, err := os.CreateTemp(tempOutPath, "outfile*.out")
	if err != nil {
		return nil, err
	}

	cmd.Stdout = tempOutFile

	if err := cmd.Run(); err != nil {
		return nil, err
	}
	_, err = utils.ResetFileOffSet(tempOutFile)
	if err != nil {
		return nil, err
	}

	return tempOutFile, nil

}

func LogOut(modelFilePath string, confInputFilePath string, cmdPath string, tempOutPath string, tempInputPath string, lineIndex int, formula string) (string, error) {

	log.Println("Starting call of input")
	// generate temp input file
	tempInputFile, err := inputs.CreateTempInputFile(modelFilePath, tempInputPath, lineIndex, formula)
	if err != nil {
		log.Fatalln("Error calling CreaTempInputFile: ", err)
		return "", err
	}

	tempOutputFile, err := Prove(confInputFilePath, tempInputFile, cmdPath, tempOutPath)
	if err != nil {
		log.Fatalln("Error calling Prove: ", err)
		return "", err
	}
	defer tempOutputFile.Close()

	// Log the content of the Prover9 output file
	tempOutBytes, err := io.ReadAll(tempOutputFile)
	if err != nil {
		fmt.Println("Error reading Prover9 output file:", err)
		return "", err
	}

	fileString := string(tempOutBytes)

	return fileString, nil

}
