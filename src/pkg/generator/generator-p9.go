package generator

import "github.com/giulianorosella/ddlogic/pkg/utils"

func getBinCombP9(list1, list2 []string) []string {
	res := []string{}
	if utils.AreStringListsEqual(list1, list2) {
		for _, val1 := range list1 {
			for _, val2 := range list2 {
				res = append(res, utils.CreateExpression(val1, val2, "*"))
				res = append(res, utils.CreateExpression(val1, val2, "^"))
				res = append(res, utils.CreateExpression(val1, val2, "v"))
			}
		}
	} else {
		for _, val1 := range list1 {
			for _, val2 := range list2 {
				res = append(res, utils.CreateExpression(val1, val2, "*"))
				res = append(res, utils.CreateExpression(val1, val2, "^"))
				res = append(res, utils.CreateExpression(val1, val2, "v"))
				res = append(res, utils.CreateExpression(val2, val1, "*"))
				res = append(res, utils.CreateExpression(val2, val1, "^"))
				res = append(res, utils.CreateExpression(val2, val1, "v"))
			}
		}
	}

	return res
}

func generateFormulasForP9(connectivesNum, varsNum int) []string {
	if connectivesNum == 0 {
		return append([]string{"0"}, utils.GetAtoms(varsNum)...)
	}
	var index [][]int
	var res []string
	for i := range connectivesNum {
		for j := range connectivesNum {
			if i >= j && i+j == connectivesNum-1 {
				index = append(index, []int{i, j})
			}
		}
		res = make([]string, 0)
		for _, pair := range index {
			res = append(res, getBinCombP9(generateFormulasForP9(pair[0], varsNum), generateFormulasForP9(pair[1], varsNum))...)
		}
	}
	return res
}

func P9FormulasCreationInit(connectivesNum, varsNum int) []string {
	res := []string{}
	formulas := generateFormulasForP9(connectivesNum, varsNum)
	for _, f := range formulas {
		res = append(res, f+" = 1.")
	}
	return res
}
