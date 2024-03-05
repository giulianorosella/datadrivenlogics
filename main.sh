#!/bin/bash

#usefull vars
launch_py="src/bash/launch_py.sh"
verify_prover9="src/bash/verify_prover.sh"
launch_prover9="src/bash/launch_prover.sh"

# Print the current working directory
echo "Executing main.sh in $(pwd)"

# Call python modules
#"$launch_py"

# Call prover9 modules only if test passes
verify_and_launch() {
    source $verify_prover9
    check_prover9
    echo

    test_prover
    echo

    if [ $? -eq 0 ]; then
        echo "Prover9 verification successful. Launching Prover9."
        # Execute the launch.sh script
        $launch_prover9
    else
        echo "Error: Prover9 verification failed. Aborting."
        exit 1
    fi
}
verify_and_launch

# Call prover9 modules directly
#"$launch_prover9"
