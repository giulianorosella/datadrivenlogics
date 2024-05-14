package generator

import (
	"github.com/aclements/go-z3/z3"
	"github.com/giulianorosella/ddlogic/pkg/models"
)

func generateBoolVars(varsNum int, ctx *z3.Context) []z3.Bool {

	varNames := "ABCDEFGHILMNOPQRSTUVZ"
	formulas := []z3.Bool{}
	for _, v := range varNames[0:varsNum] {
		letter := v
		formula := ctx.BoolConst(string(letter))
		formulas = append(formulas, formula)
	}
	return append([]z3.Bool{ctx.FromBool(false)}, formulas...)
}

func ProveClassic(c, v int) ([]models.Formula, error) {
	cfg1 := z3.NewContextConfig()
	ctx := z3.NewContext(cfg1)
	z3formulas := generateFormulasForZ3(c, v, ctx)
	var res []models.Formula
	for _, v := range z3formulas {
		f := v.AsAST().String()
		isInt := models.Unset
		isCl, err := checkTheorem(v, ctx)
		if err != nil {
			return []models.Formula{}, err
		}
		res = append(res, models.Formula{
			Expression:       f,
			IsClassicTh:      isCl,
			IsIntuitionistTh: isInt,
		})
	}
	return res, nil
}

func checkTheorem(f z3.Bool, ctx *z3.Context) (models.IsTheorem, error) {
	s := z3.NewSolver(ctx)
	s.Assert(f.Not())
	isSat, err := s.Check()
	if err != nil {
		return models.Unset, err
	}
	if isSat {
		return models.False, nil
	}
	return models.True, nil
}

func generateFormulasForZ3(connectivesNum, varsNum int, ctx *z3.Context) []z3.Bool {
	if connectivesNum == 0 {
		return generateBoolVars(varsNum, ctx)
	}
	var index [][]int
	var res []z3.Bool
	for i := range connectivesNum {
		for j := range connectivesNum {
			if i >= j && i+j == connectivesNum-1 {
				index = append(index, []int{i, j})
			}
		}
		for _, pair := range index {
			res = append(res, getBinCombZ3(generateFormulasForZ3(pair[0], varsNum, ctx), generateFormulasForZ3(pair[1], varsNum, ctx))...)
		}
	}
	return res
}

func areZ3ListEqual(list1, list2 []z3.Bool) bool {
	if len(list1) != len(list2) {
		return false
	}
	for i, v := range list1 {
		if !(v.AsAST().Equal(list2[i].AsAST())) {
			return false
		}
	}
	return true
}

func getBinCombZ3(list1, list2 []z3.Bool) []z3.Bool {
	var res []z3.Bool
	if areZ3ListEqual(list1, list2) {
		for _, v1 := range list1 {
			for _, v2 := range list2 {
				res = append(res, v1.Implies(v2))
				res = append(res, v1.And(v2))
				res = append(res, v1.Or(v2))
			}
		}
	} else {
		for _, v1 := range list1 {
			for _, v2 := range list2 {
				res = append(res, v1.Implies(v2))
				res = append(res, v1.And(v2))
				res = append(res, v1.Or(v2))
				res = append(res, v2.Implies(v1))
				res = append(res, v2.And(v1))
				res = append(res, v2.Or(v1))
			}
		}
	}
	return res
}
