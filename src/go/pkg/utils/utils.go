package utils

import "fmt"

func StringToRune(s string) (rune, error) {
	r := []rune(s)
	if len(r) != 1 {
		return 0, fmt.Errorf("Invalid delimiter: %s", s)
	}
	return r[0], nil
}
