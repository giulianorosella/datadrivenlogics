#!/bin/bash

#usefull vars
launch_py="src/bash/launch_py.sh"
launch_prover9="src/bash/launch_prover.sh"
build_go="src/bash/build_go.sh"
run_go="src/bash/launch_go.sh"


# Print the current working directory
echo "Executing main.sh in $(pwd)"
echo

# Call python modules
#"$launch_py"

echo

# Call prover9 modules directly if pass tests
#"$launch_prover9"

echo

# Build go executables
"$build_go"

# run go filename entered
"$run_go"