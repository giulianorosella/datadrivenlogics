package z3

import (
	"github.com/aclements/go-z3/z3"
)

func TestZ3() ([]z3.Bool, error) {

	cfg := z3.NewContextConfig()
	ctx := z3.NewContext(cfg)
	var_num := 3
	var_names := "ABCDEFGHILMNOPQRSTUVZ"
	formulas := []z3.Bool{}
	for i := 0; i < var_num; i++ {
		letter := var_names[i]
		formula := ctx.BoolConst(string(letter))
		formulas = append(formulas, formula)
	}
	return formulas, nil
}
