package utils

import (
	"fmt"
	"log"
	"os"
)

func StringToRune(s string) (rune, error) {
	r := []rune(s)
	if len(r) != 1 {
		return 0, fmt.Errorf("invalid delimiter: %s", s)
	}
	return r[0], nil
}

func ResetFileOffSet(file *os.File) (*os.File, error) {
	_, err := file.Seek(0, 0)
	if err != nil {
		log.Fatalln("error when resetting file: ", err)
		return nil, err
	}
	return file, nil
}

func CloseAndRemoveTempFile(file *os.File) error {
	if err := file.Close(); err != nil {
		return err
	}
	if err := os.Remove(file.Name()); err != nil {
		return err
	}
	return nil
}

func AreStringListsEqual(list1, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

func CreateExpression(str1, str2, operator string) string {
	return "(" + str1 + " " + operator + " " + str2 + ")"
}

func GetAtoms(varNumbers int) []string {
	var letters []string
	for i := 'A'; i <= 'Z'; i++ {
		letters = append(letters, string(i))
	}
	atoms := letters[0:varNumbers]
	return atoms
}
