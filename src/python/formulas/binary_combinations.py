from z3 import Implies, And, Or  # type: ignore

from typing import Any, List


def binary_comb_z3(list1: List[Any], list2: List[Any]) -> List[Any]:
    output_z3: List[Any] = []
    if list1 == list2:
        for i in range(len(list1)):
            for j in range(len(list2)):
                output_z3.append(Implies(list1[i], list2[j]))
                output_z3.append(And(list1[i], list2[j]))
                output_z3.append(Or(list1[i], list2[j]))
    else:
        for i in range(len(list1)):
            for j in range(len(list2)):
                output_z3.append(Implies(list1[i], list2[j]))
                output_z3.append(And(list1[i], list2[j]))
                output_z3.append(Or(list1[i], list2[j]))
                output_z3.append(Implies(list2[j], list1[i]))
                output_z3.append(And(list2[j], list1[i]))
                output_z3.append(Or(list2[j], list1[i]))
    return output_z3


# function designed to generate all the binary boolean combination
# of items from two list in the language of z3


def binary_comb_prov9(list1: List[Any], list2: List[Any]) -> List[str]:
    output_prov9: List[Any] = []
    if list1 == list2:
        for i in range(len(list1)):
            for j in range(len(list2)):
                output_prov9.append(
                    "(" + str(list1[i]) + " * " + str(list2[j]) + ")"
                )
                output_prov9.append(
                    "(" + str(list1[i]) + " ^ " + str(list2[j]) + ")"
                )
                output_prov9.append(
                    "(" + str(list1[i]) + " v " + str(list2[j]) + ")"
                )
    else:
        for i in range(len(list1)):
            for j in range(len(list2)):
                output_prov9.append(
                    "(" + str(list1[i]) + " * " + str(list2[j]) + ")"
                )
                output_prov9.append(
                    "(" + str(list1[i]) + " ^ " + str(list2[j]) + ")"
                )
                output_prov9.append(
                    "(" + str(list1[i]) + " v " + str(list2[j]) + ")"
                )
                output_prov9.append(
                    "(" + str(list2[j]) + " * " + str(list1[i]) + ")"
                )
                output_prov9.append(
                    "(" + str(list2[j]) + " ^ " + str(list1[i]) + ")"
                )
                output_prov9.append(
                    "(" + str(list2[j]) + " v " + str(list1[i]) + ")"
                )
    return output_prov9


# function designed to generate all the binary boolean
# combination of items from two list in the language of prover9
