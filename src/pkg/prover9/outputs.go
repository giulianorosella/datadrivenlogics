package prover9

import (
	"log"
	"strings"
)

func parsep9(output string) bool {
	res := false

	successString := "THEOREM PROVED"
	failString := "SEARCH FAILED"

	if strings.Contains(output, successString) {
		log.Println("Theorem proved")
		res = true
	}
	if strings.Contains(output, failString) {
		log.Println("Theorem not proved")

	}
	return res
}

func parseM4(output string) bool {
	res := false

	successString := "Exiting with 1 model"
	failString := "Exiting with failure"

	if strings.Contains(output, successString) {
		log.Println("Mace4 found 1 model")
		res = true
	}
	if strings.Contains(output, failString) {
		log.Println("Mace4 went out of time")

	}
	return res
}
