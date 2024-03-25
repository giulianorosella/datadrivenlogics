# this file contains all the initial imput parameters that we may want to change for our investigation
from typing import Any, List

from z3 import Bool, Implies, And, Or, Not, solve  # type: ignore # package required for checking in python whether a formula is a classical tautology or not

from string import (
    ascii_uppercase,
)  # package required for our variables, our variables can potentially be all the letters, but we will consider only some of them

range_prop_vars = list(
    ascii_uppercase
)  # the list of the range of pur propositional variable, i.e. all the uppercase letters in the alphabet

var_numbers = 5  # here the only thing we can change, ie- the number of variables we want to have

atoms_for_prover: List[str] = ["0"] + range_prop_vars[0:var_numbers]  # type: ignore # the list of our atomic formulas that can be the inputs for prover/mace4, i.e. propositional variables plus the bottom element which is the int 0

bool_vars = []

for i in range(var_numbers):
    lettera = range_prop_vars[i]
    variabile = f"{lettera}"
    locals()[variabile] = Bool(str(lettera))
    bool_vars.append(locals()[variabile])
# Genera le variabili booleane per atoms z3

atoms_for_z3: List[Any] = [False] + bool_vars
# the list of our atomic formulas that can be the inputs for z3, i.e. propositional variables plus the bottom element, which is the boolean constant "False"

con_numbers = 5  # the number of connectives we want to have
