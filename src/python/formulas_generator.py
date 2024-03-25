from middle_step import binary_comb_prov9, binary_comb_z3

from vars import atoms_for_prover, atoms_for_z3, con_numbers, var_numbers

from typing import Any, List


# ignore the following commentes function

# def formulas_z3(arg: int) -> List[Any]:
#   if arg == 0:
#      return atoms_for_z3
# else:
#    for i in range(1, arg + 1):
#       index: List[Any] = []
#      for j in range(0, arg):
#         for l in range(0, arg):
#            if j >= l and j + l == i - 1:
#               index.append([j, l])
#  form: List[Any] = []
# for h in range(len(index)):
#    form = form + binary_comb_z3(
#       formulas_z3(index[h][0]), formulas_z3(index[h][1])
#  )
# return form

# ignore the following commented function

# def formulas_prov9(arg: int) -> List[Any]:
#   if arg == 0:
#      return atoms_for_prover
# else:
#   for i in range(1, arg + 1):
#        index: List[Any] = []
#      for j in range(0, arg):
#         for l in range(0, arg):
#            if j >= l and j + l == i - 1:
#               index.append([j, l])
#  form: List[Any] = []
# for h in range(len(index)):
#    form = form + binary_comb_prov9(
#       formulas_prov9(index[h][0]), formulas_prov9(index[h][1])
#  )
# return form


def formulas_z3(arg: int) -> List[Any]:
    if arg == 0:
        return atoms_for_z3
    else:
        index: List[Any] = []
        for j in range(0, arg):
            for l in range(0, arg):
                if j >= l and j + l == arg - 1:
                    index.append([j, l])
            form: List[Any] = []
            for h in range(len(index)):
                form = form + binary_comb_z3(
                    formulas_z3(index[h][0]), formulas_z3(index[h][1])
                )
    return form


# function that takes as an input an integer n and generates all the formulas - in the language for z3 - with n
# connectives starting from our initial propositional variables atoms_for_z3


def formulas_prov9(arg: int) -> List[Any]:
    if arg == 0:
        return atoms_for_z3
    else:
        index: List[Any] = []
        for j in range(0, arg):
            for l in range(0, arg):
                if j >= l and j + l == arg - 1:
                    index.append([j, l])
            form: List[Any] = []
            for h in range(len(index)):
                form = form + binary_comb_prov9(
                    formulas_prov9(index[h][0]), formulas_prov9(index[h][1])
                )
    return form


# function that takes as an input an integer n and generates all the formulas - in the language for prover9 - with n
# connectives starting from our initial propositional variables atoms_for_prover


print(len(formulas_prov9(con_numbers)))

print(len(formulas_z3(con_numbers)))
