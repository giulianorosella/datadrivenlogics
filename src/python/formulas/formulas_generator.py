from formulas.binary_combinations import (
    binary_comb_prov9,
    binary_comb_z3,
)  # type:ignore

from vars.vars import (
    atoms_for_prover,
    atoms_for_z3,
)  # type:ignore

from typing import Any, List


def formulas_z3(arg: int) -> List[Any]:
    if arg == 0:
        return atoms_for_z3
    else:
        index: List[Any] = []
        for j in range(0, arg):
            for n in range(0, arg):
                if j >= n and j + n == arg - 1:
                    index.append([j, n])
            form: List[Any] = []
            for h in range(len(index)):
                form = form + binary_comb_z3(
                    formulas_z3(index[h][0]), formulas_z3(index[h][1])
                )
    return form


# function that takes as an input an integer n and generates
# all the formulas - in the language for z3 - with n
# connectives starting from our initial propositional variables atoms_for_z3


def formulas(arg: int) -> List[Any]:
    if arg == 0:
        return atoms_for_prover
    else:
        index: List[Any] = []
        for j in range(0, arg):
            for n in range(0, arg):
                if j >= n and j + n == arg - 1:
                    index.append([j, n])
            form: List[Any] = []
            for h in range(len(index)):
                form = form + binary_comb_prov9(
                    formulas(index[h][0]), formulas(index[h][1])
                )
    return form


def formulas_prov9(arg: int) -> List[str]:
    form_prov9: List[Any] = []
    form: List[Any] = formulas(arg)
    for i in range(len(form)):
        form_prov9.append(form[i] + str("=1."))
    return form_prov9


# function that takes as an input an integer n and generates all the formulas
# - in the language for prover9 - with n
# connectives starting from our initial propositional
# variables atoms_for_prover
