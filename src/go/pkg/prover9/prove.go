package prover9

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/giulianorosella/ddlogic/pkg/utils"
)

type CmdArgs struct {
	flags             []string
	confInputFilePath string
	tempOutPath       string
}

func Prove(confInputFilePath string, inputTempFile *os.File, cmdPath string, cmdArgs CmdArgs) (*os.File, error) {
	// execute prover9
	args := append(cmdArgs.flags, cmdArgs.confInputFilePath, cmdArgs.tempOutPath)
	cmd := exec.Command(cmdPath, args...)

	tempOutFile, err := os.CreateTemp(cmdArgs.tempOutPath, "outfile*.out")
	if err != nil {
		utils.CloseAndRemoveTempFile(tempOutFile)
		return nil, err
	}

	cmd.Stdout = tempOutFile

	// We will considers wait errors as unproved theorems
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

func LogOut(modelFilePath string, confInputFilePath string, p9CmdPath string, m4CmdPath string, tempOutPath string, tempInputPath string, lineIndex int, formula string) (bool, bool, error) {

	log.Println("Starting call of input")
	// generate temp input file
	tempInputFile, err := CreateTempInputFile(modelFilePath, tempInputPath, lineIndex, formula)
	if err != nil {
		log.Fatalln("Error calling CreaTempInputFile: ", err)
		return false, false, err
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
	p9TempOutputFile, err := Prove(confInputFilePath, tempInputFile, p9CmdPath, p9Args)
	if err != nil {
		log.Fatalln("Error calling Prover9: ", err)
		return false, false, err
	}
	//call m4
	m4TempOutputFile, err := Prove(confInputFilePath, tempInputFile, m4CmdPath, m4Args)
	if err != nil {
		log.Fatalln("Error calling Mace4: ", err)
		return false, false, err
	}

	// Log the content of the Prover9 output file
	p9TempOutBytes, err := io.ReadAll(p9TempOutputFile)
	if err != nil {
		fmt.Println("Error reading Prover9 output file:", err)
		return false, false, err
	}

	// Log the content of the mace4 output file
	m4TempOutBytes, err := io.ReadAll(m4TempOutputFile)
	if err != nil {
		fmt.Println("Error reading Prover9 output file:", err)
		return false, false, err
	}

	p9FileString := string(p9TempOutBytes)
	m4FileString := string(m4TempOutBytes)

	defer utils.CloseAndRemoveTempFile(p9TempOutputFile)
	defer utils.CloseAndRemoveTempFile(m4TempOutputFile)
	defer utils.CloseAndRemoveTempFile(tempInputFile)

	isP9 := Parsep9(p9FileString)
	isM4 := ParseM4(m4FileString)

	return isP9, isM4, nil

}
