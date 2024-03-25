from formulas_generator import formulas_z3  # type: ignore
from typing import List, Any
from z3 import solve, Not, And, Implies, Bool, Solver, sat, unsat  # type: ignore


def th_checker(arg1: List[Any]) -> List[Any]:
    s = Solver()
    solutions: List[Any] = []
    for i in range(len(arg1)):
        if s.check(Not(arg1[i])) == unsat:
            solutions.append(1)
        else:
            solutions.append(0)
    return solutions
