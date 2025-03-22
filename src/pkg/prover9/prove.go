package prover9

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/giulianorosella/ddlogic/pkg/models"
	"github.com/giulianorosella/ddlogic/pkg/utils"
)

type CmdArgs struct {
	flags             []string
	confInputFilePath string
	tempOutPath       string
}

func prove(inputTempFile *os.File, cmdPath string, cmdArgs CmdArgs) (*os.File, error) {
	args := append(cmdArgs.flags, cmdArgs.confInputFilePath, inputTempFile.Name())
	cmd := exec.Command(cmdPath, args...)

	tempOutFile, err := os.CreateTemp(cmdArgs.tempOutPath, "outfile*.out")
	if err != nil {
		utils.CloseAndRemoveTempFile(tempOutFile)
		return nil, err
	}

	cmd.Stdout = tempOutFile

	// we will considers wait errors as unproved theorems
	if err := cmd.Run(); err != nil {
		errorIsWait := false

		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok && status.ExitStatus() == 4 {
				errorIsWait = true
			}
		}

		if !errorIsWait {
			utils.CloseAndRemoveTempFile(tempOutFile)
			return nil, err
		}
	}

	_, err = utils.ResetFileOffSet(tempOutFile)
	if err != nil {
		utils.CloseAndRemoveTempFile(tempOutFile)
		return nil, err
	}

	return tempOutFile, nil

}

func GetProofs(modelFilePath string, confInputFilePath string, p9CmdPath string, m4CmdPath string, tempOutPath string, tempInputPath string, lineIndex int, f models.Formula) (models.Formula, error) {

	formula := f.Expression
	log.Println("Starting call of input")
	// generate temp input file
	tempInputFile, err := createTempInputFile(modelFilePath, tempInputPath, lineIndex, formula)
	if err != nil {
		log.Fatalln("Error calling CreaTempInputFile: ", err)
		return f, err
	}

	// conf for calling prover9 and mace4
	p9Args := CmdArgs{
		flags:             []string{"-f"},
		confInputFilePath: confInputFilePath,
		tempOutPath:       tempOutPath,
	}

	m4Args := CmdArgs{
		flags:             []string{"-c", "-f"},
		confInputFilePath: confInputFilePath,
		tempOutPath:       tempOutPath,
	}

	//call p9
	p9TempOutputFile, err := prove(tempInputFile, p9CmdPath, p9Args)
	if err != nil {
		log.Fatalln("Error calling Prover9: ", err)
		return f, err
	}

	// Log the content of the Prover9 output file
	p9TempOutBytes, err := io.ReadAll(p9TempOutputFile)
	if err != nil {
		fmt.Println("Error reading Prover9 output file:", err)
		return f, err
	}

	p9FileString := string(p9TempOutBytes)
	isP9 := parsep9(p9FileString)

	var isM4 bool
	if isP9 {
		isM4 = false
	} else {
		//call m4
		m4TempOutputFile, err := prove(tempInputFile, m4CmdPath, m4Args)
		if err != nil {
			log.Fatalln("Error calling Mace4: ", err)
			return f, err
		}

		// Log the content of the mace4 output file
		m4TempOutBytes, err := io.ReadAll(m4TempOutputFile)
		if err != nil {
			fmt.Println("Error reading Prover9 output file:", err)
			return f, err
		}

		m4FileString := string(m4TempOutBytes)

		isM4 = parseM4(m4FileString)
		defer utils.CloseAndRemoveTempFile(m4TempOutputFile)

	}

	defer utils.CloseAndRemoveTempFile(p9TempOutputFile)
	defer utils.CloseAndRemoveTempFile(tempInputFile)

	return updateFormula(isP9, isM4, f), nil

}

func updateFormula(isP9 bool, isM4 bool, formula models.Formula) models.Formula {
	if isP9 {
		formula.IsIntuitionistTh = models.True
	} else {
		if !isM4 {
			formula.IsIntuitionistTh = models.Unset
		} else {
			formula.IsIntuitionistTh = models.False
		}
	}
	return formula
}
