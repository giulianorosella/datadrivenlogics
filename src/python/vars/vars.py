from typing import Any, List

from z3 import Bool  # type: ignore

from string import (
    ascii_uppercase,
)

range_prop_vars = list(ascii_uppercase)

# here the only thing we can change, ie- the number of
# variables we want to have
var_numbers = 1


# the list of our atomic formulas that
# can be the inputs for prover/mace4, i.e. propositional variables plus
# the bottom element which is the int 0
atoms_for_prover: List[str] = ["0"] + range_prop_vars[0:var_numbers]


def generate_bool_vars() -> List[Any]:
    bool_vars = []
    for i in range(var_numbers):
        lettera = range_prop_vars[i]
        variabile = f"{lettera}"
        locals()[variabile] = Bool(str(lettera))
        bool_vars.append(locals()[variabile])
    return bool_vars
    # Genertes boolean variables for  z3


bool_vars: List[Any] = generate_bool_vars()

atoms_for_z3: List[Any] = [False] + bool_vars
# the list of our atomic formulas that can be the inputs for z3,
# i.e. propositional variables plus the bottom element,
#  which is the boolean constant "False"

con_numbers = 1  # the number of connectives we want to have


path_ex = "prover9/excel/"


rows_numbers = 1048575
